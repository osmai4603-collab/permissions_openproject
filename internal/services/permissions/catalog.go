package perms

// ──────────────────────────────────────────────
// Standard OpenProject Permission IDs
// ──────────────────────────────────────────────

const (
	// Global (Platform-wide) permissions
	PermAddProject                 = "add_project"
	PermCreateUser                 = "create_user"
	PermEditUsers                  = "edit_users"
	PermCreateBackup               = "create_backup"
	PermViewAllUsersAndGroups       = "view_all_users_and_groups"
	PermManagePublicProjectLists   = "manage_public_project_lists"
	PermCreatePortfolios           = "create_portfolios"
	PermCreatePrograms             = "create_programs"
	PermEditAttributeHelpTexts     = "edit_attribute_help_texts"
	PermManagePlaceholderUsers     = "manage_placeholder_users"
	PermViewAllUsersMailAddresses  = "view_all_users_mail_addresses"

	// Project Module permissions
	PermArchiveProject               = "archive_project"
	PermEditProject                  = "edit_project"
	PermSelectProjectModules         = "select_project_modules"
	PermViewProjectAttributes        = "view_project_attributes"
	PermExportProjects               = "export_projects"
	PermEditProjectAttributes        = "edit_project_attributes"
	PermSelectProjectAttributes      = "select_project_attributes"
	PermViewProjectPhases            = "view_project_phases"
	PermEditProjectPhases            = "edit_project_phases"
	PermSelectProjectPhases          = "select_project_phases"
	PermManageMembers                = "manage_members"
	PermInviteMembersByEmail         = "invite_members_by_email"
	PermViewMembers                  = "view_members"
	PermManageVersions               = "manage_versions"
	PermSelectTypes                  = "select_types"
	PermSelectCustomFields           = "select_custom_fields"
	PermCreateSubprojects            = "create_subprojects"
	PermCopyProjects                 = "copy_projects"
	PermManageDashboards             = "manage_dashboards"
	PermManageFilesInProject         = "manage_files_in_project"
	PermAutoManagedFoldersReadFiles  = "auto_managed_folders_read_files"
	PermAutoManagedFoldersWriteFiles = "auto_managed_folders_write_files"
	PermAutoManagedFoldersCreateFiles = "auto_managed_folders_create_files"
	PermAutoManagedFoldersDeleteFiles = "auto_managed_folders_delete_files"
	PermAutoManagedFoldersShareFiles = "auto_managed_folders_share_files"

	// Work Packages Module permissions
	PermViewWorkPackages             = "view_work_packages"
	PermAddWorkPackages              = "add_work_packages"
	PermEditWorkPackages             = "edit_work_packages"
	PermMoveWorkPackages             = "move_work_packages"
	PermDuplicateWorkPackages        = "duplicate_work_packages"
	PermAddComments                  = "add_comments"
	PermEditOwnComments              = "edit_own_comments"
	PermModerateComments             = "moderate_comments"
	PermViewInternalComments         = "view_internal_comments"
	PermWriteInternalComments        = "write_internal_comments"
	PermEditOwnInternalComments      = "edit_own_internal_comments"
	PermModerateInternalComments     = "moderate_internal_comments"
	PermAddAttachments               = "add_attachments"
	PermManageWorkPackageCategories  = "manage_work_package_categories"
	PermExportWorkPackages           = "export_work_packages"
	PermDeleteWorkPackages           = "delete_work_packages"
	PermManageWorkPackageRelations   = "manage_work_package_relations"
	PermManageWorkPackageHierarchies = "manage_work_package_hierarchies"
	PermManagePublicViews            = "manage_public_views"
	PermSaveViews                    = "save_views"
	PermViewWatchersList             = "view_watchers_list"
	PermAddWatchers                  = "add_watchers"
	PermDeleteWatchers               = "delete_watchers"
	PermShareWorkPackages            = "share_work_packages"
	PermViewWorkPackageShares        = "view_work_package_shares"
	PermAssignVersions               = "assign_versions"
	PermChangeWorkPackageStatus      = "change_work_package_status"
	PermBecomeAssigneeResponsible    = "become_assignee_responsible"
	PermViewFileLinks                = "view_file_links"
	PermManageFileLinks              = "manage_file_links"
	PermManageWikiPageLinks          = "manage_wiki_page_links"

	// Boards Module permissions
	PermViewBoards   = "view_boards"
	PermManageBoards = "manage_boards"

	// Backlogs Module permissions
	PermViewSprints                   = "view_sprints"
	PermSelectBacklogTypesAndStatuses = "select_backlog_types_and_statuses"
	PermCreateSprints                 = "create_sprints"
	PermStartCompleteSprint           = "start_complete_sprint"
	PermManageSprintItems             = "manage_sprint_items"
	PermShareSprint                   = "share_sprint"

	// Budgets Module permissions
	PermViewBudgets = "view_budgets"
	PermEditBudgets = "edit_budgets"

	// Calendars Module permissions
	PermViewCalendars         = "view_calendars"
	PermEditCalendars         = "edit_calendars"
	PermSubscribeToIcalendars = "subscribe_to_icalendars"

	// Documents Module permissions
	PermViewDocuments   = "view_documents"
	PermManageDocuments = "manage_documents"

	// Forums Module permissions
	PermManageForums     = "manage_forums"
	PermPostMessages     = "post_messages"
	PermEditMessages     = "edit_messages"
	PermEditOwnMessages  = "edit_own_messages"
	PermDeleteMessages   = "delete_messages"
	PermDeleteOwnMessages = "delete_own_messages"

	// GitHub Integration Module permissions
	PermShowGithubContent = "show_github_content"

	// GitLab Integration Module permissions
	PermShowGitlabContent = "show_gitlab_content"

	// Meetings Module permissions
	PermViewMeetings       = "view_meetings"
	PermCreateMeetings     = "create_meetings"
	PermEditMeetings       = "edit_meetings"
	PermDeleteMeetings     = "delete_meetings"
	PermSendMeetingInvites = "send_meeting_invites"
	PermManageAgendas      = "manage_agendas"
	PermManageOutcomes     = "manage_outcomes"

	// News Module permissions
	PermManageNews  = "manage_news"
	PermCommentNews = "comment_news"

	// Team Planner Module permissions
	PermViewTeamPlanner   = "view_team_planner"
	PermManageTeamPlanner = "manage_team_planner"

	// Time & Costs Module permissions
	PermViewSpentTime             = "view_spent_time"
	PermViewOwnSpentTime          = "view_own_spent_time"
	PermLogOwnTime                = "log_own_time"
	PermLogTimeForOtherUsers      = "log_time_for_other_users"
	PermEditOwnTimeLogs           = "edit_own_time_logs"
	PermEditTimeLogsForOtherUsers = "edit_time_logs_for_other_users"
	PermManageProjectActivities   = "manage_project_activities"
	PermViewOwnHourlyRate         = "view_own_hourly_rate"
	PermViewAllHourlyRates        = "view_all_hourly_rates"
	PermEditOwnHourlyRates        = "edit_own_hourly_rates"
	PermEditHourlyRates           = "edit_hourly_rates"
	PermViewCostRates             = "view_cost_rates"
	PermBookUnitCostsForOneself   = "book_unit_costs_for_oneself"
	PermBookUnitCosts             = "book_unit_costs"
	PermEditOwnBookedUnitCosts    = "edit_own_booked_unit_costs"
	PermEditBookedUnitCosts       = "edit_booked_unit_costs"
	PermViewBookedCosts           = "view_booked_costs"
	PermViewOwnBookedCosts        = "view_own_booked_costs"
	PermSavePublicCostReports     = "save_public_cost_reports"
	PermSavePrivateCostReports    = "save_private_cost_reports"

	// Wiki Module permissions
	PermViewWiki        = "view_wiki"
	PermViewWikiHistory = "view_wiki_history"
	PermEditWikiPages   = "edit_wiki_pages"
	PermManageWiki      = "manage_wiki"
)

// rawPermissionDefinitions contains baseline OpenProject permissions.
var rawPermissionDefinitions = []Permission{
	// ──────────────────────────────────────────────
	// 1. Global (platform-wide) permissions (from global.go)
	// ──────────────────────────────────────────────
	{
		ID:            PermAddProject,
		DisplayName:   "Create projects",
		Description:   "Allows users to create new top-level projects.",
		PermissibleOn: []Context{ContextGlobal},
		Module:        ModuleGlobal,
	},
	{
		ID:            PermCreateUser,
		DisplayName:   "Create users",
		Description:   "Allows users to create new user accounts.",
		PermissibleOn: []Context{ContextGlobal},
		Module:        ModuleGlobal,
	},
	{
		ID:            PermEditUsers,
		DisplayName:   "Edit users",
		Description:   "Allows users to edit existing user accounts (excluding administrators).",
		PermissibleOn: []Context{ContextGlobal},
		Module:        ModuleGlobal,
	},
	{
		ID:            PermCreateBackup,
		DisplayName:   "Create backups",
		Description:   "Allows users to create and restore system backups via the web interface.",
		PermissibleOn: []Context{ContextGlobal},
		Module:        ModuleGlobal,
	},
	{
		ID:            PermViewAllUsersAndGroups,
		DisplayName:   "View all users and groups",
		Description:   "Allows users to view all users and groups across the entire platform, bypassing project-scoped visibility.",
		PermissibleOn: []Context{ContextGlobal},
		Module:        ModuleGlobal,
	},
	{
		ID:            PermManagePublicProjectLists,
		DisplayName:   "Manage public project lists",
		Description:   "Allows users to manage and publish shared project lists.",
		PermissibleOn: []Context{ContextGlobal},
		Module:        ModuleGlobal,
	},
	{
		ID:            PermCreatePortfolios,
		DisplayName:   "Create portfolios",
		Description:   "Allows users to create project portfolios.",
		PermissibleOn: []Context{ContextGlobal},
		Module:        ModuleGlobal,
	},
	{
		ID:            PermCreatePrograms,
		DisplayName:   "Create programs",
		Description:   "Allows users to create strategic programs.",
		PermissibleOn: []Context{ContextGlobal},
		Module:        ModuleGlobal,
	},
	{
		ID:            PermEditAttributeHelpTexts,
		DisplayName:   "Edit attribute help texts",
		Description:   "Allows users to edit the explanatory help texts for fields and attributes.",
		PermissibleOn: []Context{ContextGlobal},
		Module:        ModuleGlobal,
	},
	{
		ID:            PermManagePlaceholderUsers,
		DisplayName:   "Manage placeholder users",
		Description:   "Allows users to create, edit, and delete placeholder users for resource planning.",
		PermissibleOn: []Context{ContextGlobal},
		Module:        ModuleGlobal,
	},
	{
		ID:            PermViewAllUsersMailAddresses,
		DisplayName:   "View all users' mail addresses",
		Description:   "Allows users to view email addresses of all users, even when email visibility is restricted.",
		PermissibleOn: []Context{ContextGlobal},
		Module:        ModuleGlobal,
	},

	// ──────────────────────────────────────────────
	// 2. Project & File Storage permissions (from project.go)
	// ──────────────────────────────────────────────
	{
		ID:            PermArchiveProject,
		DisplayName:   "Archive project",
		Description:   "Allows users to archive and restore a project.",
		PermissibleOn: []Context{ContextProject},
		Module:        ModuleProject,
	},
	{
		ID:            PermEditProject,
		DisplayName:   "Edit project",
		Description:   "Allows users to access Project settings and edit the project's configuration.",
		PermissibleOn: []Context{ContextProject},
		Module:        ModuleProject,
	},
	{
		ID:            PermSelectProjectModules,
		DisplayName:   "Select project modules",
		Description:   "Allows users to enable or disable project modules.",
		PermissibleOn: []Context{ContextProject},
		Module:        ModuleProject,
	},
	{
		ID:            PermViewProjectAttributes,
		DisplayName:   "View project attributes",
		Description:   "Allows users to view project information and attributes.",
		PermissibleOn: []Context{ContextProject},
		Module:        ModuleProject,
	},
	{
		ID:            PermExportProjects,
		DisplayName:   "Export projects",
		Description:   "Allows users to export project information.",
		PermissibleOn: []Context{ContextProject},
		Module:        ModuleProject,
	},
	{
		ID:            PermEditProjectAttributes,
		DisplayName:   "Edit project attributes",
		Description:   "Allows users to edit project attributes on the overview page.",
		PermissibleOn: []Context{ContextProject},
		Module:        ModuleProject,
	},
	{
		ID:            PermSelectProjectAttributes,
		DisplayName:   "Select project attributes",
		Description:   "Allows users to configure which project attributes are available.",
		PermissibleOn: []Context{ContextProject},
		Module:        ModuleProject,
	},
	{
		ID:            PermViewProjectPhases,
		DisplayName:   "View project phases",
		Description:   "Allows users to view project life cycle phases.",
		PermissibleOn: []Context{ContextProject},
		Module:        ModuleProject,
	},
	{
		ID:            PermEditProjectPhases,
		DisplayName:   "Edit project phases",
		Description:   "Allows users to edit project life cycle phases.",
		PermissibleOn: []Context{ContextProject},
		Module:        ModuleProject,
	},
	{
		ID:            PermSelectProjectPhases,
		DisplayName:   "Select project phases",
		Description:   "Allows users to activate or deactivate project phases.",
		PermissibleOn: []Context{ContextProject},
		Module:        ModuleProject,
	},
	{
		ID:            PermManageMembers,
		DisplayName:   "Manage members",
		Description:   "Allows users to add, remove and manage project members and their roles.",
		PermissibleOn: []Context{ContextProject},
		Module:        ModuleProject,
	},
	{
		ID:            PermInviteMembersByEmail,
		DisplayName:   "Invite members by email",
		Description:   "Allows users to invite project members by email.",
		PermissibleOn: []Context{ContextProject},
		Module:        ModuleProject,
		Dependencies: []Dependency{
			{PermissionID: PermManageMembers, Description: "Requires Manage members."},
		},
	},
	{
		ID:            PermViewMembers,
		DisplayName:   "View members",
		Description:   "Allows users to view project members.",
		PermissibleOn: []Context{ContextProject},
		Module:        ModuleProject,
	},
	{
		ID:            PermManageVersions,
		DisplayName:   "Manage versions",
		Description:   "Allows users to create, edit and delete project versions.",
		PermissibleOn: []Context{ContextProject},
		Module:        ModuleProject,
	},
	{
		ID:            PermSelectTypes,
		DisplayName:   "Select types",
		Description:   "Allows users to configure the work package types available in the project.",
		PermissibleOn: []Context{ContextProject},
		Module:        ModuleProject,
	},
	{
		ID:            PermSelectCustomFields,
		DisplayName:   "Select custom fields",
		Description:   "Allows users to configure which custom fields are available in the project.",
		PermissibleOn: []Context{ContextProject},
		Module:        ModuleProject,
	},
	{
		ID:            PermCreateSubprojects,
		DisplayName:   "Create subprojects",
		Description:   "Allows users to create subprojects.",
		PermissibleOn: []Context{ContextProject},
		Module:        ModuleProject,
	},
	{
		ID:            PermCopyProjects,
		DisplayName:   "Copy projects",
		Description:   "Allows users to create a new project by copying an existing project.",
		PermissibleOn: []Context{ContextProject},
		Module:        ModuleProject,
		Dependencies: []Dependency{
			{Description: "Users are assigned the configured New role for users that create projects in the copied project. Accessing Copy from Project settings typically also requires Edit project."},
		},
	},
	{
		ID:            PermManageDashboards,
		DisplayName:   "Manage dashboards",
		Description:   "Allows users to create and edit project dashboards.",
		PermissibleOn: []Context{ContextProject},
		Module:        ModuleProject,
	},
	{
		ID:            PermManageFilesInProject,
		DisplayName:   "Manage files in project",
		Description:   "Allows users to manage project file storages.",
		PermissibleOn: []Context{ContextProject},
		Module:        ModuleFileStorage,
	},
	{
		ID:            PermAutoManagedFoldersReadFiles,
		DisplayName:   "Automatically managed project folders: Read files",
		Description:   "Allows users to read files in automatically managed project folders.",
		PermissibleOn: []Context{ContextProject},
		Module:        ModuleFileStorage,
	},
	{
		ID:            PermAutoManagedFoldersWriteFiles,
		DisplayName:   "Automatically managed project folders: Write files",
		Description:   "Allows users to modify files in automatically managed project folders.",
		PermissibleOn: []Context{ContextProject},
		Module:        ModuleFileStorage,
	},
	{
		ID:            PermAutoManagedFoldersCreateFiles,
		DisplayName:   "Automatically managed project folders: Create files",
		Description:   "Allows users to create files in automatically managed project folders.",
		PermissibleOn: []Context{ContextProject},
		Module:        ModuleFileStorage,
	},
	{
		ID:            PermAutoManagedFoldersDeleteFiles,
		DisplayName:   "Automatically managed project folders: Delete files",
		Description:   "Allows users to delete files from automatically managed project folders.",
		PermissibleOn: []Context{ContextProject},
		Module:        ModuleFileStorage,
	},
	{
		ID:            PermAutoManagedFoldersShareFiles,
		DisplayName:   "Automatically managed project folders: Share files",
		Description:   "Allows users to share files from automatically managed project folders.",
		PermissibleOn: []Context{ContextProject},
		Module:        ModuleFileStorage,
	},

	// ──────────────────────────────────────────────
	// 3. Work packages and Gantt chart permissions (from workpackages.go)
	// ──────────────────────────────────────────────
	{
		ID:            PermViewWorkPackages,
		DisplayName:   "View work packages",
		Description:   "Allows users to view work packages.",
		PermissibleOn: []Context{ContextProject, ContextWorkPackage},
		Module:        ModuleWorkPackages,
	},
	{
		ID:            PermAddWorkPackages,
		DisplayName:   "Add work packages",
		Description:   "Allows users to create work packages.",
		PermissibleOn: []Context{ContextProject},
		Module:        ModuleWorkPackages,
	},
	{
		ID:            PermEditWorkPackages,
		DisplayName:   "Edit work packages",
		Description:   "Allows users to edit work packages.",
		PermissibleOn: []Context{ContextProject, ContextWorkPackage},
		Module:        ModuleWorkPackages,
	},
	{
		ID:            PermMoveWorkPackages,
		DisplayName:   "Move work packages",
		Description:   "Allows users to move work packages between projects.",
		PermissibleOn: []Context{ContextProject},
		Module:        ModuleWorkPackages,
	},
	{
		ID:            PermDuplicateWorkPackages,
		DisplayName:   "Duplicate work packages",
		Description:   "Allows users to duplicate work packages.",
		PermissibleOn: []Context{ContextProject},
		Module:        ModuleWorkPackages,
	},
	{
		ID:            PermAddComments,
		DisplayName:   "Add comments",
		Description:   "Allows users to comment on work packages.",
		PermissibleOn: []Context{ContextProject, ContextWorkPackage},
		Module:        ModuleWorkPackages,
	},
	{
		ID:            PermEditOwnComments,
		DisplayName:   "Edit own comments",
		Description:   "Allows users to edit their own comments.",
		PermissibleOn: []Context{ContextProject},
		Module:        ModuleWorkPackages,
	},
	{
		ID:            PermModerateComments,
		DisplayName:   "Moderate comments",
		Description:   "Allows users to edit comments created by any user.",
		PermissibleOn: []Context{ContextProject},
		Module:        ModuleWorkPackages,
	},
	{
		ID:            PermViewInternalComments,
		DisplayName:   "View internal comments",
		Description:   "Allows users to view internal comments.",
		PermissibleOn: []Context{ContextProject},
		Module:        ModuleWorkPackages,
	},
	{
		ID:            PermWriteInternalComments,
		DisplayName:   "Write internal comments",
		Description:   "Allows users to create internal comments.",
		PermissibleOn: []Context{ContextProject},
		Module:        ModuleWorkPackages,
	},
	{
		ID:            PermEditOwnInternalComments,
		DisplayName:   "Edit own internal comments",
		Description:   "Allows users to edit their own internal comments.",
		PermissibleOn: []Context{ContextProject},
		Module:        ModuleWorkPackages,
	},
	{
		ID:            PermModerateInternalComments,
		DisplayName:   "Moderate internal comments",
		Description:   "Allows users to edit internal comments created by any user.",
		PermissibleOn: []Context{ContextProject},
		Module:        ModuleWorkPackages,
	},
	{
		ID:            PermAddAttachments,
		DisplayName:   "Add attachments",
		Description:   "Allows users to upload attachments to work packages.",
		PermissibleOn: []Context{ContextProject, ContextWorkPackage},
		Module:        ModuleWorkPackages,
		Dependencies: []Dependency{
			{Description: "Can be granted independently of Edit work packages."},
		},
	},
	{
		ID:            PermManageWorkPackageCategories,
		DisplayName:   "Manage work package categories",
		Description:   "Allows users to create, edit, and delete work package categories.",
		PermissibleOn: []Context{ContextProject},
		Module:        ModuleWorkPackages,
	},
	{
		ID:            PermExportWorkPackages,
		DisplayName:   "Export work packages",
		Description:   "Allows users to export work packages.",
		PermissibleOn: []Context{ContextProject},
		Module:        ModuleWorkPackages,
	},
	{
		ID:            PermDeleteWorkPackages,
		DisplayName:   "Delete work packages",
		Description:   "Allows users to delete work packages.",
		PermissibleOn: []Context{ContextProject},
		Module:        ModuleWorkPackages,
	},
	{
		ID:            PermManageWorkPackageRelations,
		DisplayName:   "Manage work package relations",
		Description:   "Allows users to create, edit, and remove work package relations.",
		PermissibleOn: []Context{ContextProject},
		Module:        ModuleWorkPackages,
	},
	{
		ID:            PermManageWorkPackageHierarchies,
		DisplayName:   "Manage work package hierarchies",
		Description:   "Allows users to manage parent-child relationships between work packages.",
		PermissibleOn: []Context{ContextProject},
		Module:        ModuleWorkPackages,
	},
	{
		ID:            PermManagePublicViews,
		DisplayName:   "Manage public views",
		Description:   "Allows users to create, edit, and delete public work package views.",
		PermissibleOn: []Context{ContextProject},
		Module:        ModuleWorkPackages,
	},
	{
		ID:            PermSaveViews,
		DisplayName:   "Save views",
		Description:   "Allows users to save personal work package views.",
		PermissibleOn: []Context{ContextProject},
		Module:        ModuleWorkPackages,
	},
	{
		ID:            PermViewWatchersList,
		DisplayName:   "View watchers list",
		Description:   "Allows users to see who is watching a work package.",
		PermissibleOn: []Context{ContextProject},
		Module:        ModuleWorkPackages,
	},
	{
		ID:            PermAddWatchers,
		DisplayName:   "Add watchers",
		Description:   "Allows users to add watchers to work packages.",
		PermissibleOn: []Context{ContextProject},
		Module:        ModuleWorkPackages,
	},
	{
		ID:            PermDeleteWatchers,
		DisplayName:   "Delete watchers",
		Description:   "Allows users to remove watchers from work packages.",
		PermissibleOn: []Context{ContextProject},
		Module:        ModuleWorkPackages,
	},
	{
		ID:            PermShareWorkPackages,
		DisplayName:   "Share work packages",
		Description:   "Allows users to share work packages with other users.",
		PermissibleOn: []Context{ContextProject},
		Module:        ModuleWorkPackages,
	},
	{
		ID:            PermViewWorkPackageShares,
		DisplayName:   "View work package shares",
		Description:   "Allows users to view existing work package shares.",
		PermissibleOn: []Context{ContextProject},
		Module:        ModuleWorkPackages,
	},
	{
		ID:            PermAssignVersions,
		DisplayName:   "Assign versions",
		Description:   "Allows users to assign versions to work packages.",
		PermissibleOn: []Context{ContextProject},
		Module:        ModuleWorkPackages,
	},
	{
		ID:            PermChangeWorkPackageStatus,
		DisplayName:   "Change work package status",
		Description:   "Allows users to change the status of work packages.",
		PermissibleOn: []Context{ContextProject},
		Module:        ModuleWorkPackages,
		Dependencies: []Dependency{
			{Description: "Can be granted independently of Edit work packages."},
		},
	},
	{
		ID:            PermBecomeAssigneeResponsible,
		DisplayName:   "Become assignee/responsible",
		Description:   "Allows work packages to be assigned to users or groups with this role in the project.",
		PermissibleOn: []Context{ContextProject},
		Module:        ModuleWorkPackages,
		Dependencies: []Dependency{
			{Description: "Controls whether a user or group can be selected as assignee or responsible for work packages."},
		},
	},
	{
		ID:            PermViewFileLinks,
		DisplayName:   "View file links",
		Description:   "Allows users to view file links attached to work packages.",
		PermissibleOn: []Context{ContextProject, ContextWorkPackage},
		Module:        ModuleWorkPackages,
	},
	{
		ID:            PermManageFileLinks,
		DisplayName:   "Manage file links",
		Description:   "Allows users to create, edit, and remove file links.",
		PermissibleOn: []Context{ContextProject, ContextWorkPackage},
		Module:        ModuleWorkPackages,
	},
	{
		ID:            PermManageWikiPageLinks,
		DisplayName:   "Manage wiki page links",
		Description:   "Allows users to create, edit, and remove wiki page links.",
		PermissibleOn: []Context{ContextProject},
		Module:        ModuleWorkPackages,
	},

	// ──────────────────────────────────────────────
	// 4. Boards permissions (from boards.go)
	// ──────────────────────────────────────────────
	{
		ID:            PermViewBoards,
		DisplayName:   "View boards",
		Description:   "Allows users to view boards.",
		PermissibleOn: []Context{ContextProject},
		Module:        ModuleBoards,
	},
	{
		ID:            PermManageBoards,
		DisplayName:   "Manage boards",
		Description:   "Allows users to create, edit, and delete boards.",
		PermissibleOn: []Context{ContextProject},
		Module:        ModuleBoards,
	},

	// ──────────────────────────────────────────────
	// 5. Backlogs (Scrum) permissions (from backlogs.go)
	// ──────────────────────────────────────────────
	{
		ID:            PermViewSprints,
		DisplayName:   "View sprints",
		Description:   "Allows users to view sprints.",
		PermissibleOn: []Context{ContextProject},
		Module:        ModuleBacklogs,
	},
	{
		ID:            PermSelectBacklogTypesAndStatuses,
		DisplayName:   "Select backlog types and statuses",
		Description:   "Allows users to configure which work package types appear in the backlog and which statuses are considered completed.",
		PermissibleOn: []Context{ContextProject},
		Module:        ModuleBacklogs,
	},
	{
		ID:            PermCreateSprints,
		DisplayName:   "Create sprints",
		Description:   "Allows users to create sprints.",
		PermissibleOn: []Context{ContextProject},
		Module:        ModuleBacklogs,
	},
	{
		ID:            PermStartCompleteSprint,
		DisplayName:   "Start/complete sprint",
		Description:   "Allows users to start and complete sprints.",
		PermissibleOn: []Context{ContextProject},
		Module:        ModuleBacklogs,
	},
	{
		ID:            PermManageSprintItems,
		DisplayName:   "Manage sprint items",
		Description:   "Allows users to manage sprint contents.",
		PermissibleOn: []Context{ContextProject},
		Module:        ModuleBacklogs,
	},
	{
		ID:            PermShareSprint,
		DisplayName:   "Share sprint",
		Description:   "Allows users to share sprint information.",
		PermissibleOn: []Context{ContextProject},
		Module:        ModuleBacklogs,
	},

	// ──────────────────────────────────────────────
	// 6. Budgets permissions (from budgets.go)
	// ──────────────────────────────────────────────
	{
		ID:            PermViewBudgets,
		DisplayName:   "View budgets",
		Description:   "Allows users to view project budgets.",
		PermissibleOn: []Context{ContextProject},
		Module:        ModuleBudgets,
	},
	{
		ID:            PermEditBudgets,
		DisplayName:   "Edit budgets",
		Description:   "Allows users to create, edit, and delete project budgets.",
		PermissibleOn: []Context{ContextProject},
		Module:        ModuleBudgets,
	},

	// ──────────────────────────────────────────────
	// 7. Calendars permissions (from calendars.go)
	// ──────────────────────────────────────────────
	{
		ID:            PermViewCalendars,
		DisplayName:   "View calendars",
		Description:   "Allows users to view calendars.",
		PermissibleOn: []Context{ContextProject},
		Module:        ModuleCalendars,
	},
	{
		ID:            PermEditCalendars,
		DisplayName:   "Edit calendars",
		Description:   "Allows users to create, edit, and delete calendars.",
		PermissibleOn: []Context{ContextProject},
		Module:        ModuleCalendars,
	},
	{
		ID:            PermSubscribeToIcalendars,
		DisplayName:   "Subscribe to iCalendars",
		Description:   "Allows users to subscribe to calendar feeds.",
		PermissibleOn: []Context{ContextProject},
		Module:        ModuleCalendars,
	},

	// ──────────────────────────────────────────────
	// 8. Documents permissions (from documents.go)
	// ──────────────────────────────────────────────
	{
		ID:            PermViewDocuments,
		DisplayName:   "View documents",
		Description:   "Allows users to view project documents.",
		PermissibleOn: []Context{ContextProject},
		Module:        ModuleDocuments,
	},
	{
		ID:            PermManageDocuments,
		DisplayName:   "Manage documents",
		Description:   "Allows users to create, edit, and delete project documents.",
		PermissibleOn: []Context{ContextProject},
		Module:        ModuleDocuments,
	},

	// ──────────────────────────────────────────────
	// 9. Forums permissions (from forums.go)
	// ──────────────────────────────────────────────
	{
		ID:            PermManageForums,
		DisplayName:   "Manage forums",
		Description:   "Allows users to create, edit, and delete forums.",
		PermissibleOn: []Context{ContextProject},
		Module:        ModuleForums,
	},
	{
		ID:            PermPostMessages,
		DisplayName:   "Post messages",
		Description:   "Allows users to post messages in forums.",
		PermissibleOn: []Context{ContextProject},
		Module:        ModuleForums,
	},
	{
		ID:            PermEditMessages,
		DisplayName:   "Edit messages",
		Description:   "Allows users to edit any forum message.",
		PermissibleOn: []Context{ContextProject},
		Module:        ModuleForums,
	},
	{
		ID:            PermEditOwnMessages,
		DisplayName:   "Edit own messages",
		Description:   "Allows users to edit their own forum messages.",
		PermissibleOn: []Context{ContextProject},
		Module:        ModuleForums,
	},
	{
		ID:            PermDeleteMessages,
		DisplayName:   "Delete messages",
		Description:   "Allows users to delete any forum message.",
		PermissibleOn: []Context{ContextProject},
		Module:        ModuleForums,
	},
	{
		ID:            PermDeleteOwnMessages,
		DisplayName:   "Delete own messages",
		Description:   "Allows users to delete their own forum messages.",
		PermissibleOn: []Context{ContextProject},
		Module:        ModuleForums,
	},

	// ──────────────────────────────────────────────
	// 10. GitHub integration permissions (from github.go)
	// ──────────────────────────────────────────────
	{
		ID:            PermShowGithubContent,
		DisplayName:   "Show GitHub content",
		Description:   "Allows users to see GitHub pull requests and issues linked to work packages.",
		PermissibleOn: []Context{ContextProject},
		Module:        ModuleGitHub,
	},

	// ──────────────────────────────────────────────
	// 11. GitLab integration permissions (from gitlab.go)
	// ──────────────────────────────────────────────
	{
		ID:            PermShowGitlabContent,
		DisplayName:   "Show GitLab content",
		Description:   "Allows users to see GitLab merge requests and issues linked to work packages.",
		PermissibleOn: []Context{ContextProject},
		Module:        ModuleGitLab,
	},

	// ──────────────────────────────────────────────
	// 12. Meetings permissions (from meetings.go)
	// ──────────────────────────────────────────────
	{
		ID:            PermViewMeetings,
		DisplayName:   "View meetings",
		Description:   "Allows users to view meetings in a project.",
		PermissibleOn: []Context{ContextProject},
		Module:        ModuleMeetings,
	},
	{
		ID:            PermCreateMeetings,
		DisplayName:   "Create meetings",
		Description:   "Allows users to create new meetings.",
		PermissibleOn: []Context{ContextProject},
		Module:        ModuleMeetings,
	},
	{
		ID:            PermEditMeetings,
		DisplayName:   "Edit meetings",
		Description:   "Allows users to edit existing meetings.",
		PermissibleOn: []Context{ContextProject},
		Module:        ModuleMeetings,
	},
	{
		ID:            PermDeleteMeetings,
		DisplayName:   "Delete meetings",
		Description:   "Allows users to delete meetings.",
		PermissibleOn: []Context{ContextProject},
		Module:        ModuleMeetings,
	},
	{
		ID:            PermSendMeetingInvites,
		DisplayName:   "Send meeting invites",
		Description:   "Allows users to send email invitations for meetings to participants.",
		PermissibleOn: []Context{ContextProject},
		Module:        ModuleMeetings,
	},
	{
		ID:            PermManageAgendas,
		DisplayName:   "Manage agendas",
		Description:   "Allows users to create, edit, and close meeting agendas.",
		PermissibleOn: []Context{ContextProject},
		Module:        ModuleMeetings,
	},
	{
		ID:            PermManageOutcomes,
		DisplayName:   "Manage outcomes",
		Description:   "Allows users to create, edit, and close meeting outcomes/minutes.",
		PermissibleOn: []Context{ContextProject},
		Module:        ModuleMeetings,
	},

	// ──────────────────────────────────────────────
	// 13. News permissions (from news.go)
	// ──────────────────────────────────────────────
	{
		ID:            PermManageNews,
		DisplayName:   "Manage news",
		Description:   "Allows users to create, edit, and delete news entries.",
		PermissibleOn: []Context{ContextProject},
		Module:        ModuleNews,
	},
	{
		ID:            PermCommentNews,
		DisplayName:   "Comment news",
		Description:   "Allows users to post comments on news entries.",
		PermissibleOn: []Context{ContextProject},
		Module:        ModuleNews,
	},

	// ──────────────────────────────────────────────
	// 14. Team planner permissions (from teamplanner.go)
	// ──────────────────────────────────────────────
	{
		ID:            PermViewTeamPlanner,
		DisplayName:   "View team planner",
		Description:   "Allows users to view team planner views.",
		PermissibleOn: []Context{ContextProject},
		Module:        ModuleTeamPlanner,
	},
	{
		ID:            PermManageTeamPlanner,
		DisplayName:   "Manage team planner",
		Description:   "Allows users to create, edit, and delete team planner views.",
		PermissibleOn: []Context{ContextProject},
		Module:        ModuleTeamPlanner,
	},

	// ──────────────────────────────────────────────
	// 15. Time & Costs permissions (from timecosts.go)
	// ──────────────────────────────────────────────
	{
		ID:            PermViewSpentTime,
		DisplayName:   "View spent time",
		Description:   "Allows users to view time entries logged by all users.",
		PermissibleOn: []Context{ContextProject},
		Module:        ModuleTimeCosts,
	},
	{
		ID:            PermViewOwnSpentTime,
		DisplayName:   "View own spent time",
		Description:   "Allows users to view only their own time entries.",
		PermissibleOn: []Context{ContextProject},
		Module:        ModuleTimeCosts,
	},
	{
		ID:            PermLogOwnTime,
		DisplayName:   "Log own time",
		Description:   "Allows users to log time on work packages for themselves.",
		PermissibleOn: []Context{ContextProject},
		Module:        ModuleTimeCosts,
	},
	{
		ID:            PermLogTimeForOtherUsers,
		DisplayName:   "Log time for other users",
		Description:   "Allows users to log time entries on behalf of other users.",
		PermissibleOn: []Context{ContextProject},
		Module:        ModuleTimeCosts,
	},
	{
		ID:            PermEditOwnTimeLogs,
		DisplayName:   "Edit own time logs",
		Description:   "Allows users to edit their own logged time entries.",
		PermissibleOn: []Context{ContextProject},
		Module:        ModuleTimeCosts,
	},
	{
		ID:            PermEditTimeLogsForOtherUsers,
		DisplayName:   "Edit time logs for other users",
		Description:   "Allows users to edit time entries logged by other users.",
		PermissibleOn: []Context{ContextProject},
		Module:        ModuleTimeCosts,
	},
	{
		ID:            PermManageProjectActivities,
		DisplayName:   "Manage project activities",
		Description:   "Allows users to manage time tracking activity types within a project.",
		PermissibleOn: []Context{ContextProject},
		Module:        ModuleTimeCosts,
	},
	{
		ID:            PermViewOwnHourlyRate,
		DisplayName:   "View own hourly rate",
		Description:   "Allows users to view their own hourly labor rate.",
		PermissibleOn: []Context{ContextProject},
		Module:        ModuleTimeCosts,
	},
	{
		ID:            PermViewAllHourlyRates,
		DisplayName:   "View all hourly rates",
		Description:   "Allows users to view hourly rates of all members.",
		PermissibleOn: []Context{ContextProject},
		Module:        ModuleTimeCosts,
	},
	{
		ID:            PermEditOwnHourlyRates,
		DisplayName:   "Edit own hourly rates",
		Description:   "Allows users to edit their own hourly labor rate.",
		PermissibleOn: []Context{ContextProject},
		Module:        ModuleTimeCosts,
	},
	{
		ID:            PermEditHourlyRates,
		DisplayName:   "Edit hourly rates",
		Description:   "Allows users to edit hourly rates for all members.",
		PermissibleOn: []Context{ContextProject},
		Module:        ModuleTimeCosts,
	},
	{
		ID:            PermViewCostRates,
		DisplayName:   "View cost rates",
		Description:   "Allows users to view cost types and their rates.",
		PermissibleOn: []Context{ContextProject},
		Module:        ModuleTimeCosts,
	},
	{
		ID:            PermBookUnitCostsForOneself,
		DisplayName:   "Book unit costs for oneself",
		Description:   "Allows users to book unit costs on work packages for themselves.",
		PermissibleOn: []Context{ContextProject},
		Module:        ModuleTimeCosts,
	},
	{
		ID:            PermBookUnitCosts,
		DisplayName:   "Book unit costs",
		Description:   "Allows users to book unit costs on work packages for all members.",
		PermissibleOn: []Context{ContextProject},
		Module:        ModuleTimeCosts,
	},
	{
		ID:            PermEditOwnBookedUnitCosts,
		DisplayName:   "Edit own booked unit costs",
		Description:   "Allows users to edit their own booked unit cost entries.",
		PermissibleOn: []Context{ContextProject},
		Module:        ModuleTimeCosts,
	},
	{
		ID:            PermEditBookedUnitCosts,
		DisplayName:   "Edit booked unit costs",
		Description:   "Allows users to edit unit cost entries booked by any member.",
		PermissibleOn: []Context{ContextProject},
		Module:        ModuleTimeCosts,
	},
	{
		ID:            PermViewBookedCosts,
		DisplayName:   "View booked costs",
		Description:   "Allows users to view booked costs for all members.",
		PermissibleOn: []Context{ContextProject},
		Module:        ModuleTimeCosts,
	},
	{
		ID:            PermViewOwnBookedCosts,
		DisplayName:   "View own booked costs",
		Description:   "Allows users to view only their own booked costs.",
		PermissibleOn: []Context{ContextProject},
		Module:        ModuleTimeCosts,
	},
	{
		ID:            PermSavePublicCostReports,
		DisplayName:   "Save public cost reports",
		Description:   "Allows users to save cost reports visible to all project members.",
		PermissibleOn: []Context{ContextProject},
		Module:        ModuleTimeCosts,
	},
	{
		ID:            PermSavePrivateCostReports,
		DisplayName:   "Save private cost reports",
		Description:   "Allows users to save cost reports visible only to themselves.",
		PermissibleOn: []Context{ContextProject},
		Module:        ModuleTimeCosts,
	},

	// ──────────────────────────────────────────────
	// 16. Wiki permissions (from wiki.go)
	// ──────────────────────────────────────────────
	{
		ID:            PermViewWiki,
		DisplayName:   "View wiki",
		Description:   "Allows users to view wiki pages.",
		PermissibleOn: []Context{ContextProject},
		Module:        ModuleWiki,
	},
	{
		ID:            PermViewWikiHistory,
		DisplayName:   "View wiki history",
		Description:   "Allows users to view the edit history of wiki pages.",
		PermissibleOn: []Context{ContextProject},
		Module:        ModuleWiki,
	},
	{
		ID:            PermEditWikiPages,
		DisplayName:   "Edit wiki pages",
		Description:   "Allows users to create and edit wiki pages.",
		PermissibleOn: []Context{ContextProject},
		Module:        ModuleWiki,
	},
	{
		ID:            PermManageWiki,
		DisplayName:   "Manage wiki",
		Description:   "Allows users to manage the wiki, including renaming, deleting pages, and managing the wiki menu.",
		PermissibleOn: []Context{ContextProject},
		Module:        ModuleWiki,
	},
}

// GetDefaultPermissions returns all baseline OpenProject permissions defined in the catalog.
func GetDefaultPermissions() []Permission {
	out := make([]Permission, len(rawPermissionDefinitions))
	copy(out, rawPermissionDefinitions)
	return out
}

