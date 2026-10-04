package perms

// ──────────────────────────────────────────────
// Global (platform-wide) permissions
// ──────────────────────────────────────────────

var (
	AddProject = Permission{
		ID:            "add_project",
		DisplayName:   "Create projects",
		Description:   "Allows users to create new top-level projects.",
		PermissibleOn: []Context{ContextGlobal},
		Module:        ModuleGlobal,
	}

	CreateUser = Permission{
		ID:            "create_user",
		DisplayName:   "Create users",
		Description:   "Allows users to create new user accounts.",
		PermissibleOn: []Context{ContextGlobal},
		Module:        ModuleGlobal,
	}

	EditUsers = Permission{
		ID:            "edit_users",
		DisplayName:   "Edit users",
		Description:   "Allows users to edit existing user accounts (excluding administrators).",
		PermissibleOn: []Context{ContextGlobal},
		Module:        ModuleGlobal,
	}

	CreateBackup = Permission{
		ID:            "create_backup",
		DisplayName:   "Create backups",
		Description:   "Allows users to create and restore system backups via the web interface.",
		PermissibleOn: []Context{ContextGlobal},
		Module:        ModuleGlobal,
	}

	ViewAllUsersAndGroups = Permission{
		ID:            "view_all_users_and_groups",
		DisplayName:   "View all users and groups",
		Description:   "Allows users to view all users and groups across the entire platform, bypassing project-scoped visibility.",
		PermissibleOn: []Context{ContextGlobal},
		Module:        ModuleGlobal,
	}

	ManagePublicProjectLists = Permission{
		ID:            "manage_public_project_lists",
		DisplayName:   "Manage public project lists",
		Description:   "Allows users to manage and publish shared project lists.",
		PermissibleOn: []Context{ContextGlobal},
		Module:        ModuleGlobal,
	}

	CreatePortfolios = Permission{
		ID:            "create_portfolios",
		DisplayName:   "Create portfolios",
		Description:   "Allows users to create project portfolios.",
		PermissibleOn: []Context{ContextGlobal},
		Module:        ModuleGlobal,
	}

	CreatePrograms = Permission{
		ID:            "create_programs",
		DisplayName:   "Create programs",
		Description:   "Allows users to create strategic programs.",
		PermissibleOn: []Context{ContextGlobal},
		Module:        ModuleGlobal,
	}

	EditAttributeHelpTexts = Permission{
		ID:            "edit_attribute_help_texts",
		DisplayName:   "Edit attribute help texts",
		Description:   "Allows users to edit the explanatory help texts for fields and attributes.",
		PermissibleOn: []Context{ContextGlobal},
		Module:        ModuleGlobal,
	}

	ManagePlaceholderUsers = Permission{
		ID:            "manage_placeholder_users",
		DisplayName:   "Manage placeholder users",
		Description:   "Allows users to create, edit, and delete placeholder users for resource planning.",
		PermissibleOn: []Context{ContextGlobal},
		Module:        ModuleGlobal,
	}

	ViewAllUsersMailAddresses = Permission{
		ID:            "view_all_users_mail_addresses",
		DisplayName:   "View all users' mail addresses",
		Description:   "Allows users to view email addresses of all users, even when email visibility is restricted.",
		PermissibleOn: []Context{ContextGlobal},
		Module:        ModuleGlobal,
	}
)

// GlobalPermissions returns all global (platform-wide) permissions.
func GlobalPermissions() []Permission {
	return []Permission{
		AddProject,
		CreateUser,
		EditUsers,
		CreateBackup,
		ViewAllUsersAndGroups,
		ManagePublicProjectLists,
		CreatePortfolios,
		CreatePrograms,
		EditAttributeHelpTexts,
		ManagePlaceholderUsers,
		ViewAllUsersMailAddresses,
	}
}
