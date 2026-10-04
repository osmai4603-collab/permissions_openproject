# Database tables of the authorization system (Permissions) in OpenProject

This document is a **data-model reference** for the tables that back
OpenProject's RBAC system. For the *application-level* concept (how
permissions are declared and how to check them), see the official developer
documentation [Development concept: Permissions](./development/concepts/permissions/)
and the Arabic application-level analysis [permissions-analysis-ar.md](./permissions-analysis-ar.md).
This file only covers what lives in PostgreSQL.

All DDL below is quoted from `db/structure.sql`. The schema is managed by
`db/migrate` and dumped into `db/structure.sql`; `roles`, `role_permissions`,
`member_roles`, `members` and `custom_actions_roles` originate from the
aggregated baseline `db/migrate/1000016_aggregated_migrations.rb` via
`db/migrate/tables/*.rb`.

## Contents

1. [The graph](#1-the-graph)
2. [Core tables](#2-core-tables)
   - [`roles`](#21-roles)
   - [`role_permissions`](#22-role_permissions)
   - [`members`](#23-members)
   - [`member_roles`](#24-member_roles)
   - [`group_users`](#25-group_users)
   - [`group_details`](#26-group_details)
3. [Secondary permission tables](#3-secondary-permission-tables)
   - [`enabled_modules`](#31-enabled_modules)
   - [`workflows`](#32-workflows)
   - [`custom_fields_roles`](#33-custom_fields_roles)
   - [`custom_actions_roles`](#34-custom_actions_roles)
4. [How a permission check reaches the database](#4-how-a-permission-check-reaches-the-database)
5. [Built-in roles](#5-built-in-roles)
6. [Group role inheritance (`inherited_from`)](#6-group-role-inheritance-inherited_from)
7. [Entity-scoped roles](#7-entity-scoped-roles)
8. [SSO/Enterprise tables that feed the group graph](#8-ssoenterprise-tables-that-feed-the-group-graph)
9. [Look-alike tables that are *not* authorization tables](#9-look-alike-tables-that-are-not-authorization-tables)
10. [Migration history](#10-migration-history)
11. [Integrity gaps and quirks](#11-integrity-gaps-and-quirks)
12. [Useful SQL recipes](#12-useful-sql-recipes)

---

## 1. The graph

```text
users                                   (Principal STI: User, Group, AnonymousUser,
  │                                       SystemUser, PlaceholderUser)
  │
  ├──< group_users >── users              group membership  (Group ↔ User)
  ├──< group_details (1:1) >── users      group metadata    (organizational_unit, parent)
  │
  ▼
members                                 one row = "principal P is a member of scope S"
  │  user_id, project_id, entity_type, entity_id
  ▼
member_roles                            (member, role, inherited_from)
  │
  ▼
roles                                   RBAC role  (STI: GlobalRole, ProjectRole,
  │                                       WorkPackageRole, ProjectQueryRole)
  ▼
role_permissions                        (role_id, permission) — one row per granted permission
```

Gating / side tables consulted at check time:

| Table                    | Role in the check                                                            |
|--------------------------|------------------------------------------------------------------------------|
| `enabled_modules`        | A project-scoped permission is only granted if its module is enabled         |
| `workflows`              | Allowed status transitions, per role and type variant                        |
| `projects.public`        | Decides whether the builtin Non-member / Anonymous role applies              |
| `users.admin`, `users.status` | Admin bypass (only for `grant_to_admin?` permissions) and locked/deleted rejection        |

Everything hangs together as:

> `permission ∈ role_permissions ⋈ roles ⋈ member_roles ⋈ members ⋈ users`
> **AND** `role` must be reachable in the right scope
> **AND** `project.enabled_modules` must enable the permission's module
> **AND** the user is not `locked` / `deleted`

---

## 2. Core tables

### 2.1 `roles`

```sql
CREATE TABLE public.roles (
    id bigint NOT NULL,
    name character varying DEFAULT ''::character varying NOT NULL,
    "position" integer DEFAULT 1,
    builtin integer DEFAULT 0 NOT NULL,
    type character varying(30) DEFAULT 'Role'::character varying,
    created_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP
);
```

Model: [`app/models/role.rb`](https://github.com/opf/openproject/blob/dev/app/models/role.rb)
plus STI subclasses `global_role.rb`, `project_role.rb`, `work_package_role.rb`,
`project_query_role.rb`.

| Column       | Type                     | Meaning                                                                                             |
|--------------|--------------------------|-----------------------------------------------------------------------------------------------------|
| `id`         | `bigint` PK              |                                                                                                     |
| `name`       | `varchar NOT NULL`       | Display name. Unique **case-insensitively** (`index_roles_on_LOWER_name`). Max 256 chars.           |
| `position`   | `integer DEFAULT 1`      | `acts_as_list` ordering column. Used for display order and for picking the default new-project role. |
| `builtin`    | `integer NOT NULL`       | `0` = user-defined. Any other value = builtin, see [§5](#5-built-in-roles).                          |
| `type`       | `varchar(30)`            | STI discriminator: `GlobalRole`, `ProjectRole`, `WorkPackageRole`, `ProjectQueryRole`.                |
| `created_at` / `updated_at` | `timestamptz`  | Defaults to `CURRENT_TIMESTAMP`; the model touches `updated_at` via `MemberRole belongs_to :member, touch: true`. |

Indexes / constraints:

```sql
CREATE UNIQUE INDEX "index_roles_on_LOWER_name" ON public.roles USING btree (lower((name)::text));
```

Notes:

- `WorkPackageRole` and `ProjectQueryRole` are hidden from the roles UI
  (`Role::HIDDEN_ROLE_TYPES`, `scope :visible`).
- `Role#builtin?` is simply `builtin != 0`; `Role#member?` is `!builtin?`.
  `MemberRole#validate_project_member_role` rejects built-in roles, so built-ins
  can never be attached to a member. `WorkPackageRole#member?` is overridden to
  `true` so entity roles *can* be attached.
- `Role#deletable?` is `!builtin?`; `before_destroy(prepend: true)` raises
  `ActiveRecord::RecordNotDestroyed` for built-ins.
- `Role.default_scope` eagerly `includes(:role_permissions)` — permission
  lookups are almost always eager-loaded.
- Built-in roles are created lazily on demand by `ProjectRole.non_member`,
  `ProjectRole.anonymous` and `GlobalRole.standard` if a database is missing them.

### 2.2 `role_permissions`

```sql
CREATE TABLE public.role_permissions (
    id bigint NOT NULL,
    permission character varying,
    role_id bigint,
    created_at timestamp(6) with time zone NOT NULL,
    updated_at timestamp(6) with time zone NOT NULL
);
```

Model: [`app/models/role_permission.rb`](https://github.com/opf/openproject/blob/dev/app/models/role_permission.rb)
— essentially a `belongs_to :role` plus `validates :permission, presence: true`.

```sql
CREATE INDEX index_role_permissions_on_role_id ON public.role_permissions USING btree (role_id);
```

**There is no unique index on `(role_id, permission)` and no foreign key on
`role_id`.** Duplicate grants for the same role are possible at the database
level; `Role#add_permission` is only called for permissions not already present
(`Role#permissions=` diffs before writing), and the app-level integrity is what
holds today.

#### Permission encoding

`permission` stores the **plain symbol name as a string** (e.g.
`'edit_work_packages'`). This is a change from Redmine/ChiliProject, which
packed the permission map into a single `permissions` bitmask integer column.
Consequences:

- A permission is stored once per role, so granting a new permission is an
  `INSERT`, and renaming one is a `DELETE` + `INSERT`.
- No integer decoding step at runtime — `Role#permissions` is just
  `role_permissions.map { |rp| rp.permission.to_sym }`.
- Stale rows for permissions that no longer exist simply never match, which is
  why permission-removal migrations can be no-ops in `down`.

Permissions are declared in two places:

- Core: [`config/initializers/permissions.rb`](https://github.com/opf/openproject/blob/dev/config/initializers/permissions.rb) — **39** permissions across 6 project modules (`nil`, `work_package_tracking`, `news`, `repository`, `forums`, `activity`).
- Modules: each `modules/*/lib/**/engine.rb` declares its own inside `OpenProject::AccessControl.map` — **65** more (backlogs, bim, boards, budgets, calendar, costs, documents, github_integration, meeting, overviews, reporting, resource_management, storages, team_planner, wikis).

So a fully-loaded installation has roughly **104** distinct permission strings.

Declaration shape (`lib/open_project/access_control/permission.rb`):

```ruby
map.permission :manage_members,
               { members: %i[index new create update destroy] },
               permissible_on: :project,       # :global | :project | :work_package | :project_query
               require: :member,               # :member | :loggedin | nil
               public: false,                  # granted to anyone with any role in scope
               visible: true,
               project_module: :members,
               dependencies: :view_all_principals,
               contract_actions: { members: %i[read create] },
               grant_to_admin: true
```

`dependencies:` is enforced in
[`app/contracts/roles/base_contract.rb`](https://github.com/opf/openproject/blob/dev/app/contracts/roles/base_contract.rb)
(`check_permission_prerequisites`) — you cannot grant a permission without its
dependencies.

### 2.3 `members`

```sql
CREATE TABLE public.members (
    id bigint NOT NULL,
    user_id bigint NOT NULL,
    project_id bigint,
    created_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
    entity_type character varying,
    entity_id bigint
);
```

Model: [`app/models/member.rb`](https://github.com/opf/openproject/blob/dev/app/models/member.rb).

A member is *"(principal) is a member of (scope)"*. The scope is polymorphic:

| `project_id` | `entity_type` / `entity_id` | Meaning                                                        |
|--------------|------------------------------|----------------------------------------------------------------|
| `NOT NULL`   | both `NULL`                  | Project membership. `project_id IS NULL AND entity_* IS NULL` = **global membership** |
| `NOT NULL`   | `'WorkPackage'` / wp id      | Work-package-level membership (sharing / WP roles)              |
| `NOT NULL`   | `'ProjectQuery'` / query id  | Saved project query membership                                  |
| `NULL`       | both `NULL`                  | Global membership — carries `GlobalRole`s                        |

```sql
CREATE INDEX index_members_on_user_id     ON public.members USING btree (user_id);
CREATE INDEX index_members_on_project_id  ON public.members USING btree (project_id);
CREATE INDEX index_members_on_entity      ON public.members USING btree (entity_type, entity_id);

CREATE UNIQUE INDEX index_members_on_user_id_and_project_without_entity
  ON public.members USING btree (user_id, project_id)
  WHERE ((entity_type IS NULL) AND (entity_id IS NULL));

CREATE UNIQUE INDEX index_members_on_user_id_and_project_with_entity
  ON public.members USING btree (user_id, project_id, entity_type, entity_id)
  WHERE ((entity_type IS NOT NULL) AND (entity_id IS NOT NULL));
```

Two *partial* unique indexes instead of one composite, because a project
membership and an entity membership must not collide.

`user_id` points at `users.id` but is a legacy name — it is the `Principal`
(`belongs_to :principal, foreign_key: "user_id"`). It can hold a `Group`'s id,
which is how group-derived permissions work (see §6).

Validation:

```ruby
validates :user_id, uniqueness: { scope: %i[project_id entity_type entity_id] }
validates :entity_type, inclusion: { in: ALLOWED_ENTITIES, allow_blank: true }
# ALLOWED_ENTITIES = %w[WorkPackage ProjectQuery]
```

There is **no foreign key on `user_id` or `project_id`.**

### 2.4 `member_roles`

```sql
CREATE TABLE public.member_roles (
    id bigint NOT NULL,
    member_id bigint NOT NULL,
    role_id bigint NOT NULL,
    inherited_from bigint
);
```

Model: [`app/models/member_role.rb`](https://github.com/opf/openproject/blob/dev/app/models/member_role.rb).

The join table between memberships and roles. Three columns of interest:

- `member_id` — FK-free reference to `members.id`.
- `role_id` — FK-free reference to `roles.id`. Must point at a role where
  `member?` is true (`validate_project_member_role`), i.e. not a built-in.
- `inherited_from` — `NULL` for directly-assigned roles; otherwise the
  `member_roles.id` of the *group* membership row this role was copied from.
  See [§6](#6-group-role-inheritance-inherited_from).

```sql
CREATE INDEX index_member_roles_on_member_id     ON public.member_roles USING btree (member_id);
CREATE INDEX index_member_roles_on_role_id       ON public.member_roles USING btree (role_id);
CREATE INDEX index_member_roles_on_inherited_from ON public.member_roles USING btree (inherited_from);

CREATE UNIQUE INDEX unique_inherited_role
  ON public.member_roles USING btree (member_id, role_id, inherited_from);
```

`unique_inherited_role` only deduplicates **inherited** rows in practice —
PostgreSQL treats the `NULL` in `inherited_from` as distinct, so several
`(member_id, role_id, NULL)` rows are permitted (they mean the same role was
assigned twice on purpose, e.g. via two different code paths).

The app-level distinction is expressed through scopes:

```ruby
scope :only_inherited,     -> { where.not(inherited_from: nil) }
scope :only_non_inherited, -> { where(inherited_from: nil) }
```

### 2.5 `group_users`

```sql
CREATE TABLE public.group_users (
    id bigint NOT NULL,
    group_id bigint NOT NULL,
    user_id bigint NOT NULL
);
```

Model: [`app/models/group_user.rb`](https://github.com/opf/openproject/blob/dev/app/models/group_user.rb) —
`belongs_to :group` and `belongs_to :user`, both non-validated-presence-only.

```sql
CREATE UNIQUE INDEX group_user_ids ON public.group_users USING btree (group_id, user_id);
CREATE UNIQUE INDEX index_group_users_on_user_id_and_group_id
  ON public.group_users USING btree (user_id, group_id);
```

Two unique indexes on the same column pair in both orders — redundant but
harmless (the first is the historical name).

This table does **not** grant permissions by itself. It is the source for
`Groups::CreateInheritedRolesService`, which materialises group permissions into
`members` + `member_roles` for each individual user (see §6).

### 2.6 `group_details`

```sql
CREATE TABLE public.group_details (
    id bigint NOT NULL,
    principal_id bigint NOT NULL,
    organizational_unit boolean DEFAULT false NOT NULL,
    parent_id bigint,
    created_at timestamp(6) with time zone NOT NULL,
    updated_at timestamp(6) with time zone NOT NULL
);
```

Model: [`app/models/groups/hierarchy.rb`](https://github.com/opf/openproject/blob/dev/app/models/groups/hierarchy.rb)
(`Group#parent`, `#organizational_unit?`, `#users` scoping).

A 1:1 side table on `users` holding attributes only groups have. Introduced in
`20260309130829_add_group_details.rb` to replace the columns that used to live
on `groups`.

```sql
CREATE UNIQUE INDEX index_group_details_on_principal_id
  ON public.group_details USING btree (principal_id);
CREATE INDEX index_group_details_on_organizational_unit
  ON public.group_details USING btree (organizational_unit);
CREATE INDEX index_group_details_on_parent_id
  ON public.group_details USING btree (parent_id);

ALTER TABLE ONLY public.group_details
  ADD CONSTRAINT fk_rails_2b463af4b4 FOREIGN KEY (principal_id) REFERENCES public.users(id);
ALTER TABLE ONLY public.group_details
  ADD CONSTRAINT fk_rails_6b0aad5675 FOREIGN KEY (parent_id) REFERENCES public.users(id) ON DELETE SET NULL;
```

`parent_id` is a hierarchy of **organizational units** (departments), not of
arbitrary groups. It changes the *group membership propagation* semantics
(`Groups::AncestorMembershipPropagation`) and group-name uniqueness scoping, so
it indirectly affects which `group_users` rows exist, and therefore which
inherited roles users get.

---

## 3. Secondary permission tables

### 3.1 `enabled_modules`

```sql
CREATE TABLE public.enabled_modules (
    id bigint NOT NULL,
    project_id bigint,
    name character varying NOT NULL
);
```

```sql
CREATE INDEX enabled_modules_project_id ON public.enabled_modules USING btree (project_id);
CREATE INDEX index_enabled_modules_on_name ON public.enabled_modules USING btree (name);
```

Not an authorization table per se, but the **hard gate** for module-scoped
permissions. `Project#allowed_permissions` intersects the project's enabled
module names with `OpenProject::AccessControl.modules_permissions(names)`:

```ruby
# app/models/project.rb
def allowed_permissions
  names = enabled_modules.loaded? ? enabled_module_names : enabled_modules.pluck(:name)
  OpenProject::AccessControl.modules_permissions(names).map(&:name)
end
```

`Authorization::UserPermissibleService#allowed_in_single_project?` filters the
requested permissions through `project.allowed_permissions` **before**
intersecting with the role permissions, so a role that holds
`manage_backlogs` gets nothing if the `backlogs` module is off. Uniqueness of
`(project_id, name)` is validated in the model only
(`validates :name, uniqueness: { scope: :project_id }`).

### 3.2 `workflows`

```sql
CREATE TABLE public.workflows (
    id bigint NOT NULL,
    old_status_id bigint NOT NULL,
    new_status_id bigint NOT NULL,
    role_id bigint NOT NULL,
    assignee boolean DEFAULT false NOT NULL,
    author boolean DEFAULT false NOT NULL,
    type_variant_id bigint NOT NULL
);
```

Model: [`app/models/workflow.rb`](https://github.com/opf/openproject/blob/dev/app/models/workflow.rb).

Role-scoped state machine: given the current status, the user's roles in the
project and whether they are the author/assignee, this table decides the set of
reachable statuses.

```sql
CREATE INDEX index_workflows_on_role_id       ON public.workflows USING btree (role_id);
CREATE INDEX index_workflows_on_old_status_id ON public.workflows USING btree (old_status_id);
CREATE INDEX index_workflows_on_new_status_id ON public.workflows USING btree (new_status_id);
CREATE INDEX index_workflows_on_type_variant_id ON public.workflows USING btree (type_variant_id);
CREATE INDEX wkfs_role_type_variant_old_status
  ON public.workflows USING btree (role_id, type_variant_id, old_status_id);

ALTER TABLE ONLY public.workflows
  ADD CONSTRAINT fk_rails_0c5f149c21 FOREIGN KEY (role_id) REFERENCES public.roles(id)
  ON UPDATE CASCADE ON DELETE CASCADE;
```

This is one of the **two** real foreign keys into `roles`. Deleting a role
cascades its workflows. Workflows are also copied between roles by
`Workflow.copy` / `Role#workflows#copy_from_role` in the role admin UI.

### 3.3 `custom_fields_roles`

```sql
CREATE TABLE public.custom_fields_roles (
    id bigint NOT NULL,
    custom_field_id bigint NOT NULL,
    role_id bigint NOT NULL,
    created_at timestamp(6) with time zone NOT NULL,
    updated_at timestamp(6) with time zone NOT NULL
);
```

```sql
CREATE UNIQUE INDEX index_custom_fields_roles_on_custom_field_id
  ON public.custom_fields_roles USING btree (custom_field_id);
CREATE INDEX index_custom_fields_roles_on_role_id
  ON public.custom_fields_roles USING btree (role_id);

ALTER TABLE ONLY public.custom_fields_roles
  ADD CONSTRAINT fk_rails_950d979cb9 FOREIGN KEY (role_id) REFERENCES public.roles(id);
ALTER TABLE ONLY public.custom_fields_roles
  ADD CONSTRAINT fk_rails_96bf507f15 FOREIGN KEY (custom_field_id) REFERENCES public.custom_fields(id);
```

This is the odd one out: `custom_field_id` is **unique**, so it is a 1:1 link,
not a join table. `ProjectCustomField` declares it as such:

```ruby
has_one :custom_fields_role, foreign_key: :custom_field_id, dependent: :destroy
has_one :role, through: :custom_fields_role
```

Semantics: a **user-type project custom field whose value determines membership
role**. Filling the field with a user adds that role to their `members` row;
clearing it removes it. Implemented by
`Projects::ManageMembershipsFromCustomFieldsService` and guarded by
`CustomField#user_field_with_role_assignment?`:

```ruby
is_a?(ProjectCustomField) && user? && custom_fields_role.present?
```

Cleanup on role removal is handled in
`MemberRole#cleanup_associated_custom_values`.

### 3.4 `custom_actions_roles`

```sql
CREATE TABLE public.custom_actions_roles (
    id bigint NOT NULL,
    role_id bigint,
    custom_action_id bigint
);
```

```sql
CREATE INDEX index_custom_actions_roles_on_role_id         ON public.custom_actions_roles USING btree (role_id);
CREATE INDEX index_custom_actions_roles_on_custom_action_id ON public.custom_actions_roles USING btree (custom_action_id);
```

A plain HABTM join (no unique index, no FKs). `CustomAction` declares
`has_and_belongs_to_many :role_conditions, class_name: "Role"`; the behaviour
lives in `CustomActions::Conditions::Role`, which evaluates
`user.roles_for_project(work_package.project)` against the stored role ids.
Empty `values` means "no role restriction".

---

## 4. How a permission check reaches the database

Entry points on `User`
([`app/models/users/permission_checks.rb`](https://github.com/opf/openproject/blob/dev/app/models/users/permission_checks.rb)):

| Method | Scope |
|--------|-------|
| `allowed_globally?(permission)` | global |
| `allowed_in_project?(permission, project)` | project |
| `allowed_in_any_project?(permission)` | any project |
| `allowed_in_work_package?(permission, wp)` / `allowed_in_any_work_package?` | entity (generated from `Member::ALLOWED_ENTITIES`) |
| `allowed_in_project_query?(permission, query)` | entity |
| `allowed_based_on_permission_context?(permission, project:, entity:)` | guesses the context |

All of them delegate to `Authorization::UserPermissibleService`. The actual
data access is:

1. `Authorization.contextual_permissions` filters the `map.permission`
   declarations by `permissible_on:`. Unknown name → `UnknownPermissionError`;
   wrong context → `IllegalPermissionContextError`.
2. `User#all_permissions_for(context)` (`users/permission_checks.rb:112`) is the
   single read point. For admins it returns every `grant_to_admin?` permission
   for the context. Otherwise it calls `Authorization.roles(user, context)` and
   `pluck(:permission)` from `role_permissions`.
3. `Authorization.roles` (`app/services/authorization.rb:50`) dispatches to one of
   four Arel queries, all built from the same four tables:

| Context | Query class | Extra conditions |
|---------|-------------|------------------|
| `Project` | `Authorization::UserProjectRolesQuery` | `members.project_id = <id> AND members.entity_type IS NULL AND members.entity_id IS NULL`; **plus** `OR members.id IS NULL AND roles.builtin = 1|2` when `project.public?` |
| `WorkPackage` / `ProjectQuery` | `Authorization::UserEntityRolesQuery` | `members.entity_type = entity.class.to_s AND members.entity_id = entity.id`, plus `members.project_id` when the entity exposes `project_id` |
| `nil` (global) | `Authorization::UserGlobalRolesQuery` | `members.entity_type IS NULL AND members.entity_id IS NULL`; **`OR roles.builtin IN (1, 8)`** |
| anything else | `Authorization::UserGlobalRolesQuery` | same as global |

The joins are defined once in
[`Authorization::AbstractUserQuery`](https://github.com/opf/openproject/blob/dev/app/services/authorization/abstract_user_query.rb):

```ruby
users       LEFT JOIN members       ON users.id       = members.user_id
members     LEFT JOIN member_roles ON members.id     = member_roles.member_id
member_roles LEFT JOIN roles       ON member_roles.role_id = roles.id
```

Group-derived roles are **not** joined here. They are already materialised into
`members` / `member_roles` by the inheritance services, so the query is a plain
single-principal lookup.

The `builtin` fallback is what makes Non-member / Anonymous work without a
`members` row: for a public project, a logged-in user gets `builtin = 1` and an
anonymous visitor gets `builtin = 2`; globally, everyone additionally gets
`builtin = 8` (Standard global role).

Everything is memoised per request:

```ruby
# Users::ProjectRoleCache / reset_permission_caches
user.roles_for_project(project)         # roles in a project
user.roles_for_work_package(wp)          # project roles + WP-specific roles
user.all_permissions_for(context)        # plain array of permission symbols
```

Guard rails at the top of every check:

```ruby
def authorizable_user?
  !(user.locked? || user.deleted?) || user.is_a?(SystemUser)
end

def admin_and_all_granted_to_admin?(permissions)
  user.active_admin? && permissions.all?(&:grant_to_admin?)
end
```

---

## 5. Built-in roles

`roles.builtin` constants (`app/models/role.rb:37`):

| Constant | Value | `type` | Meaning |
|----------|-------|--------|---------|
| `NON_BUILTIN` | 0 | any | user-defined, assignable, deletable |
| `BUILTIN_NON_MEMBER` | 1 | `ProjectRole` | applies to any logged-in user in a public project they are not a member of |
| `BUILTIN_ANONYMOUS` | 2 | `ProjectRole` | same, for unauthenticated visitors |
| `BUILTIN_WORK_PACKAGE_VIEWER` | 3 | `WorkPackageRole` | view-only, work-package scope |
| `BUILTIN_WORK_PACKAGE_COMMENTER` | 4 | `WorkPackageRole` | + comment |
| `BUILTIN_WORK_PACKAGE_EDITOR` | 5 | `WorkPackageRole` | + edit |
| `BUILTIN_PROJECT_QUERY_VIEW` | 6 | `ProjectQueryRole` | view a saved project query |
| `BUILTIN_PROJECT_QUERY_EDIT` | 7 | `ProjectQueryRole` | + edit |
| `BUILTIN_STANDARD_GLOBAL` | 8 | `GlobalRole` | minimal global role everyone has |

Built-in roles are hidden from `Role.givable`, are rejected by
`MemberRole#validate_project_member_role` (with the `WorkPackageRole#member?`
override), and cannot be destroyed (`Role#deletable?`).

Default permission sets are seeded from `app/seeders/common.yml` and read by
`BasicData::BaseRoleSeeder` (which also merges per-module `modules_permissions`
overrides such as `modules/boards/app/seeders/common.yml`):

```yaml
project_roles:
  - reference: :default_role_non_member
    builtin: :non_member          # → builtin = 1
    position: 0
  - reference: :default_role_anonymous
    builtin: :anonymous           # → builtin = 2
  - reference: :default_role_member          # builtin = 0, position 3
  - reference: :default_role_reader          # builtin = 0, position 4
  - reference: :default_role_project_admin
    permissions: :all_assignable_permissions # resolved via Roles::CreateContract
global_roles:
  - reference: :default_role_project_creator_and_staff_manager  # add_project, create_user, manage_user, …
  - reference: :default_role_standard_global
    builtin: :standard_global      # → builtin = 8, permissions: []
```

`ProjectRole.in_new_project` picks the role offered to a non-admin project
creator: among `ProjectRole.assignable_to_project_creator` (givable roles
containing all of `PERMISSIONS_FOR_PROJECT_CREATOR` =
`view_project, view_project_attributes, edit_project_attributes, view_members,
manage_members`), preferring `Setting.new_project_user_role_id` then the lowest
`position`.

---

## 6. Group role inheritance (`inherited_from`)

A group is a first-class `Principal`, so it can be a member of a project
(`members.user_id = <group id>`) and hold roles. Those roles are then
**materialised onto every user of the group** as extra `members` rows and
`member_roles` rows tagged with `inherited_from`.

`Groups::CreateInheritedRolesService` does it in a single CTE chain
(`INSERT … SELECT … ON CONFLICT DO NOTHING`):

```sql
WITH found_users          AS (SELECT id AS user_id FROM users WHERE id IN (:user_ids)),
     group_memberships    AS (SELECT project_id, user_id, entity_type, entity_id
                              FROM members WHERE user_id = :group_id AND <project limit>),
     group_roles          AS (SELECT members.project_id, members.entity_type, members.entity_id,
                                     member_roles.role_id, member_roles.id AS member_role_id
                              FROM member_roles
                              JOIN members ON members.id = member_roles.member_id
                                           AND members.user_id = :group_id
                              WHERE member_roles.inherited_from IS NULL),
     existing_members     AS (…),
     new_members          AS (INSERT INTO members (project_id, user_id, updated_at, created_at, entity_type, entity_id)
                              SELECT … FROM found_users, group_memberships
                              WHERE NOT EXISTS (SELECT 1 FROM existing_members …)
                              ON CONFLICT DO NOTHING RETURNING id, user_id, project_id, entity_type, entity_id),
     add_roles            AS (INSERT INTO member_roles (member_id, role_id, inherited_from)
                              SELECT members.id, group_roles.role_id, group_roles.member_role_id
                              FROM group_roles
                              JOIN (SELECT * FROM new_members UNION SELECT * FROM existing_members) members
                                ON group_roles.project_id   IS NOT DISTINCT FROM members.project_id
                               AND group_roles.entity_type IS NOT DISTINCT FROM members.entity_type
                               AND group_roles.entity_id   IS NOT DISTINCT FROM members.entity_id
                              ON CONFLICT DO NOTHING
                              RETURNING id, member_id, role_id)
SELECT … FROM add_roles …;
```

Reading it:

- Only `member_roles` rows with `inherited_from IS NULL` are propagated — the
  group must not have received its roles by inheritance itself (no transitive
  explosion).
- `inherited_from` is set to the **group's `member_roles.id`**, i.e. the row that
  is the *origin* of the grant. `Groups::CleanupInheritedRolesService` removes a
  user's inherited row once that originating row is gone.
- `IS NOT DISTINCT FROM` is used throughout to make `NULL` project/entity
  comparable, so global and entity-scoped group roles propagate too.
- Inserting a `members` row is required even when the user had no membership yet;
  the `ON CONFLICT DO NOTHING` relies on the two partial unique indexes (§2.3).

Related services:

| Service | Purpose |
|---------|---------|
| `Groups::CreateInheritedRolesService` | add inherited rows |
| `Groups::UpdateRolesService` | sync when a group's own roles change (adds, removes, and deletes orphaned rows) |
| `Groups::CleanupInheritedRolesService` | delete rows whose `inherited_from` no longer resolves |
| `Groups::AncestorMembershipPropagation` | propagate along the `group_details.parent_id` hierarchy |
| `Groups::UpdateService` | prune inherited rows when group hierarchy changes |

Because `inherited_from` is meaningful, member deletion must be careful:
`Members::DeleteService` / `Shares::DeleteService` only destroy the
`inherited_from IS NULL` rows and check that no inherited rows remain.

---

## 7. Entity-scoped roles

`members.entity_type` is limited to `Member::ALLOWED_ENTITIES`
(`%w[WorkPackage ProjectQuery]`, enforced by validation):

- **`WorkPackage`** — work-package roles (`WorkPackageRole`) plus the
  `Sharing` mechanism. `user.roles_for_work_package(wp)` =
  `roles_for_project(wp.project) + roles of Member.of_work_package(wp)`.
- **`ProjectQuery`** — `ProjectQueryRole` (`view_project_query`,
  `edit_project_query`), the newest addition, mirroring the entity roles for
  shared work-package queries.

`Authorization.roles` unwraps `SimpleDelegator` before dispatching, because
queries decorate work packages in `WorkPackageEagerLoadingWrapper` and the
class name is used to pick the query.

---

## 8. SSO/Enterprise tables that feed the group graph

These are *authentication*-side tables, but they change `group_users`, which in
turn changes inherited roles. They are therefore part of the authorization data
flow, just not RBAC proper.

| Table | Owner | Effect on authorization |
|-------|-------|-------------------------|
| `oidc_group_links` | `OpenIDConnect::GroupLink` | Maps an OIDC group claim name to an OpenProject group (`group_id`, `auth_provider_id`, `oidc_group_name`) |
| `oidc_group_memberships` | `OpenIDConnect::GroupMembership` | Marks a `group_users` row as externally sourced (`auth_provider_id`, `group_user_id`). **The only FK into `group_users` in the schema** (`ON DELETE CASCADE`) |
| `ldap_groups_memberships` | Enterprise LDAP sync | Materialised `group_users` rows coming from an LDAP group |
| `ldap_departments_memberships` | Enterprise LDAP sync | Materialised memberships in `organizational_unit` groups |
| `ldap_groups_synchronized_groups`, `ldap_departments_synchronized_departments`, `ldap_groups_synchronized_filters` | Enterprise LDAP sync | Which LDAP groups/departments are synchronised |
| `auth_providers`, `user_auth_provider_links`, `oidc_user_tokens`, `remote_identities` | SSO | Identity → `users.id` mapping; determines *who* the principal is |
| `scim_clients`, `service_account_associations` | Enterprise SCIM | Service accounts and their linked resources — a service account is a `User` and can hold roles, so it participates in RBAC like any other principal |
| `enterprise_tokens` | Enterprise | Gates the availability of the above features |

---

## 9. Look-alike tables that are *not* authorization tables

Listed to avoid false positives when auditing.

| Table | Why it is not an authorization table |
|-------|------------------------------------------|
| `settings` | Holds global configuration (`login_required`, `default_projects_public`, `default_projects_modules`, `new_project_user_role_id`). Feeds *which* built-in roles/modules apply, but stores no per-role grant. Definitions in `config/constants/settings/definition.rb`. |
| `queries.public`, `project_queries` / `persisted_queries` `public` flag, `views`, `persisted_views.public` | Visibility *within* an already-authorized scope, plus a `public` boolean. Object-level sharing of a saved query, not RBAC. |
| `watchers`, `favorites` | Per-user subscriptions; no permission semantics. |
| `tokens`, `two_factor_authentication_devices`, `oauth_*` | Authentication / API tokens. |
| `enabled_modules` | Not grants, but a hard gate (see §3.1). |
| `done_statuses_for_project`, `statuses` | Status configuration consumed by `workflows`. |
| `grids`, `grid_widgets`, `menu_items`, `custom_styles` | UI configuration; `menu_items` has an `admin` scope but no role linkage. |
| `placeholder_user_details`, `resource_allocations`, `service_account_associations` | Feature data; only relevant to RBAC insofar as the owner is a principal that holds roles. |

---

## 10. Migration history

Permissions data is migrated with two helpers that live next to the migrations:

- `db/migrate/migration_utils/permission_adder.rb` — `add(having, add, force: false)` grants `add` to every role that has **all** of `having`, skipping `require_member?` permissions for the Non-member role and `require_loggedin?` permissions for Anonymous.
- `db/migrate/migration_utils/permission_renamer.rb` — `rename(from, to)` = `PermissionAdder.add(from, to)` + `RolePermission.delete_by(permission: from)`.

Notable migrations:

| Migration | Change |
| ----------- | -------- |
| `1000016_aggregated_migrations.rb` + `db/migrate/tables/{roles,role_permissions,member_roles,custom_actions_roles}.rb` | Baseline DDL for the five core tables |
| `20250422072119_rename_comment_permissions.rb` | 7 comment permission renames (`add_work_package_notes` → `add_work_package_comments`, `*_comments_with_restricted_visibility` → `*_internal_comments`) |
| `20250929070310_add_view_all_principals_permission_to_existing_roles.rb` | Introduces `:view_all_principals`; creates the "View all users (migration)" global role, back-fills `:manage_user` → `:view_all_principals`, and grants a global role to every user who already had `manage_members` in any project (excluding `PlaceholderUser`) |
| `20251029134217_avoid_global_role_on_placeholder_user.rb` | Cleans up global roles on placeholder users |
| `20251111145039_create_custom_field_role_association.rb` | Creates `custom_fields_roles` with a **unique** index on `custom_field_id` |
| `20251209191935_remove_template_permissions_from_roles.rb` | `RolePermission.delete_by(...)` of the three `*_from_template` permissions |
| `20251218100721_add_uniqueness_for_role_names.rb` | Renames duplicates then adds `UNIQUE (lower(name))` concurrently |
| `20260212145213_migrate_backlogs_permissions.rb` | `view_master_backlog` / `view_taskboards` → `view_sprints`; `update_sprints` → `create_sprints`; adds `manage_versions` → `create_sprints` and `assign_versions` → `manage_sprint_items` |
| `20260223142025_add_view_budgets_to_roles_with_edit_budgets.rb` | `edit_budgets` → `view_budgets` |
| `20260304160505_fix_sprint_role_dependencies.rb` | `manage_sprint_items` → `view_sprints`, `create_sprints` → `view_sprints`, `view_sprints` → `view_work_packages` |
| `20260309130829_add_group_details.rb` | Creates `group_details` and back-fills one row per `users.type = 'Group'` |
| `20260402094525_add_group_detail_index_for_ous.rb` | Index for organizational-unit lookups |
| `20260623113435_nullify_group_details_parent_on_delete.rb` | `parent_id` → `ON DELETE SET NULL` |
| `20260622144833_relax_group_lastname_uniqueness.rb` | Group name uniqueness becomes scoped to the OU parent |

---

## 11. Integrity gaps and quirks

Worth knowing before writing raw SQL or a data migration:

1. **Almost no foreign keys.** The only FKs into the RBAC tables are
   `workflows.role_id`, `custom_fields_roles.role_id`,
   `custom_fields_roles.custom_field_id`, `group_details.{principal_id,parent_id}`
   and `oidc_group_memberships.group_user_id`. `members`, `member_roles`,
   `role_permissions`, `group_users`, `enabled_modules` and
   `custom_actions_roles` have **no** FKs. Referential integrity is enforced by
   Active Record (`dependent:`, `belongs_to`) and by service objects. A raw
   `DELETE FROM roles` will orphan `role_permissions` and `member_roles` rows.
2. **No unique index on `role_permissions (role_id, permission)`.** Duplicate
   grants are possible; `Role#permissions` returns duplicates.
3. **`unique_inherited_role` does not constrain directly-assigned roles.**
   `NULL != NULL` in a PostgreSQL unique index, so `(member_id, role_id, NULL)`
   may repeat.
4. **Global memberships are not DB-unique either.** For a global member
   `project_id IS NULL`, so `index_members_on_user_id_and_project_without_entity`
   cannot deduplicate. `Members::AddRoleService` relies on
   `Member.find_by(user_id:, project_id:, entity:)` to find the existing row.
5. **`role_permissions.permission` can hold orphan strings.** Nothing constrains
   it to the `map.permission` registry; unknown values are silently ignored at
   check time (`Authorization.permissions_for` filters them out) — the query in
   §4 only intersects with declared names.
6. **`roles.builtin` is a loose integer.** There is no check constraint; only
   `Role`'s constants define meaning. A value of `9` would be treated as a
   built-in (non-zero) but would match no query.
7. **`enabled_modules (project_id, name)` has no unique index** — the model
   validates it, the database does not.

---

## 12. Useful SQL recipes

Everything is granted to a user (name it first):

```sql
-- Permissions of one user in one project
SELECT DISTINCT rp.permission
FROM users u
JOIN members m          ON m.user_id = u.id
JOIN member_roles mr    ON mr.member_id = m.id
JOIN role_permissions rp ON rp.role_id = mr.role_id
WHERE u.login = 'jsmith'
  AND m.project_id = 42
  AND m.entity_type IS NULL AND m.entity_id IS NULL
ORDER BY rp.permission;

-- Same, plus the Non-member/Anonymous fallback of a public project
SELECT DISTINCT rp.permission, r.name AS role, r.builtin, mr.inherited_from
FROM roles r
JOIN role_permissions rp ON rp.role_id = r.id
WHERE r.builtin IN (1, 2, 8)
ORDER BY r.builtin, r.name;

-- Which roles does a user hold in a project, and where did each come from?
SELECT r.name, r.type, r.builtin, mr.inherited_from IS NULL AS assigned_directly,
       gm.user_id AS inherited_from_group
FROM member_roles mr
JOIN roles r   ON r.id = mr.role_id
LEFT JOIN member_roles gm ON gm.id = mr.inherited_from
LEFT JOIN members gm_m    ON gm_m.id = gm.member_id
WHERE mr.member_id = (SELECT id FROM members WHERE user_id = 42 AND project_id = 42
                      AND entity_type IS NULL AND entity_id IS NULL);

-- Roles nobody has (or nobody can have), i.e. stale grants after a permission rename
SELECT r.id, r.name, rp.permission
FROM role_permissions rp
JOIN roles r ON r.id = rp.role_id
WHERE rp.permission NOT IN ( /* … the names from OpenProject::AccessControl.permissions … */);

-- Membership counts per scope
SELECT m.project_id, m.entity_type, count(*)
FROM members m GROUP BY m.project_id, m.entity_type ORDER BY 2 DESC, 3 DESC LIMIT 20;

-- Users holding a role both directly and by group inheritance (duplicated work)
SELECT mr.member_id, mr.role_id, count(*)
FROM member_roles mr GROUP BY mr.member_id, mr.role_id HAVING count(*) > 1;

-- Projects where a module is off but a role still grants a module permission
SELECT DISTINCT em.name AS disabled_module, rp.permission, r.name
FROM enabled_modules em
CROSS JOIN roles r
JOIN role_permissions rp ON rp.role_id = r.id
WHERE r.builtin = 0
  AND NOT EXISTS (SELECT 1 FROM enabled_modules e2 WHERE e2.project_id = em.project_id AND e2.name = em.name)
LIMIT 50;

-- Pre-flight before deleting a role (no FKs will save you)
SELECT 'role_permissions' AS t, count(*) FROM role_permissions WHERE role_id = $1
UNION ALL SELECT 'member_roles',        count(*) FROM member_roles        WHERE role_id = $1
UNION ALL SELECT 'workflows',           count(*) FROM workflows           WHERE role_id = $1
UNION ALL SELECT 'custom_fields_roles', count(*) FROM custom_fields_roles WHERE role_id = $1
UNION ALL SELECT 'custom_actions_roles',count(*) FROM custom_actions_roles WHERE role_id = $1;
```