package perms

// ──────────────────────────────────────────────
// Project permissions
// ──────────────────────────────────────────────

var (
	ArchiveProject = Permission{
		ID:            "archive_project",
		DisplayName:   "Archive project",
		Description:   "Allows users to archive and restore a project.",
		PermissibleOn: []Context{ContextProject},
		Module:        ModuleProject,
	}

	EditProject = Permission{
		ID:            "edit_project",
		DisplayName:   "Edit project",
		Description:   "Allows users to access Project settings and edit the project's configuration.",
		PermissibleOn: []Context{ContextProject},
		Module:        ModuleProject,
	}

	SelectProjectModules = Permission{
		ID:            "select_project_modules",
		DisplayName:   "Select project modules",
		Description:   "Allows users to enable or disable project modules.",
		PermissibleOn: []Context{ContextProject},
		Module:        ModuleProject,
	}

	ViewProjectAttributes = Permission{
		ID:            "view_project_attributes",
		DisplayName:   "View project attributes",
		Description:   "Allows users to view project information and attributes.",
		PermissibleOn: []Context{ContextProject},
		Module:        ModuleProject,
	}

	ExportProjects = Permission{
		ID:            "export_projects",
		DisplayName:   "Export projects",
		Description:   "Allows users to export project information.",
		PermissibleOn: []Context{ContextProject},
		Module:        ModuleProject,
	}

	EditProjectAttributes = Permission{
		ID:            "edit_project_attributes",
		DisplayName:   "Edit project attributes",
		Description:   "Allows users to edit project attributes on the overview page.",
		PermissibleOn: []Context{ContextProject},
		Module:        ModuleProject,
	}

	SelectProjectAttributes = Permission{
		ID:            "select_project_attributes",
		DisplayName:   "Select project attributes",
		Description:   "Allows users to configure which project attributes are available.",
		PermissibleOn: []Context{ContextProject},
		Module:        ModuleProject,
	}

	ViewProjectPhases = Permission{
		ID:            "view_project_phases",
		DisplayName:   "View project phases",
		Description:   "Allows users to view project life cycle phases.",
		PermissibleOn: []Context{ContextProject},
		Module:        ModuleProject,
	}

	EditProjectPhases = Permission{
		ID:            "edit_project_phases",
		DisplayName:   "Edit project phases",
		Description:   "Allows users to edit project life cycle phases.",
		PermissibleOn: []Context{ContextProject},
		Module:        ModuleProject,
	}

	SelectProjectPhases = Permission{
		ID:            "select_project_phases",
		DisplayName:   "Select project phases",
		Description:   "Allows users to activate or deactivate project phases.",
		PermissibleOn: []Context{ContextProject},
		Module:        ModuleProject,
	}

	ManageMembers = Permission{
		ID:            "manage_members",
		DisplayName:   "Manage members",
		Description:   "Allows users to add, remove and manage project members and their roles.",
		PermissibleOn: []Context{ContextProject},
		Module:        ModuleProject,
	}

	InviteMembersByEmail = Permission{
		ID:            "invite_members_by_email",
		DisplayName:   "Invite members by email",
		Description:   "Allows users to invite project members by email.",
		PermissibleOn: []Context{ContextProject},
		Module:        ModuleProject,
		Dependencies: []Dependency{
			{PermissionID: "manage_members", Description: "Requires Manage members."},
		},
	}

	ViewMembers = Permission{
		ID:            "view_members",
		DisplayName:   "View members",
		Description:   "Allows users to view project members.",
		PermissibleOn: []Context{ContextProject},
		Module:        ModuleProject,
	}

	ManageVersions = Permission{
		ID:            "manage_versions",
		DisplayName:   "Manage versions",
		Description:   "Allows users to create, edit and delete project versions.",
		PermissibleOn: []Context{ContextProject},
		Module:        ModuleProject,
	}

	SelectTypes = Permission{
		ID:            "select_types",
		DisplayName:   "Select types",
		Description:   "Allows users to configure the work package types available in the project.",
		PermissibleOn: []Context{ContextProject},
		Module:        ModuleProject,
	}

	SelectCustomFields = Permission{
		ID:            "select_custom_fields",
		DisplayName:   "Select custom fields",
		Description:   "Allows users to configure which custom fields are available in the project.",
		PermissibleOn: []Context{ContextProject},
		Module:        ModuleProject,
	}

	CreateSubprojects = Permission{
		ID:            "create_subprojects",
		DisplayName:   "Create subprojects",
		Description:   "Allows users to create subprojects.",
		PermissibleOn: []Context{ContextProject},
		Module:        ModuleProject,
	}

	CopyProjects = Permission{
		ID:            "copy_projects",
		DisplayName:   "Copy projects",
		Description:   "Allows users to create a new project by copying an existing project.",
		PermissibleOn: []Context{ContextProject},
		Module:        ModuleProject,
		Dependencies: []Dependency{
			{Description: "Users are assigned the configured New role for users that create projects in the copied project. Accessing Copy from Project settings typically also requires Edit project."},
		},
	}

	ManageDashboards = Permission{
		ID:            "manage_dashboards",
		DisplayName:   "Manage dashboards",
		Description:   "Allows users to create and edit project dashboards.",
		PermissibleOn: []Context{ContextProject},
		Module:        ModuleProject,
	}

	ManageFilesInProject = Permission{
		ID:            "manage_files_in_project",
		DisplayName:   "Manage files in project",
		Description:   "Allows users to manage project file storages.",
		PermissibleOn: []Context{ContextProject},
		Module:        ModuleFileStorage,
	}

	AutoManagedFoldersReadFiles = Permission{
		ID:            "auto_managed_folders_read_files",
		DisplayName:   "Automatically managed project folders: Read files",
		Description:   "Allows users to read files in automatically managed project folders.",
		PermissibleOn: []Context{ContextProject},
		Module:        ModuleFileStorage,
	}

	AutoManagedFoldersWriteFiles = Permission{
		ID:            "auto_managed_folders_write_files",
		DisplayName:   "Automatically managed project folders: Write files",
		Description:   "Allows users to modify files in automatically managed project folders.",
		PermissibleOn: []Context{ContextProject},
		Module:        ModuleFileStorage,
	}

	AutoManagedFoldersCreateFiles = Permission{
		ID:            "auto_managed_folders_create_files",
		DisplayName:   "Automatically managed project folders: Create files",
		Description:   "Allows users to create files in automatically managed project folders.",
		PermissibleOn: []Context{ContextProject},
		Module:        ModuleFileStorage,
	}

	AutoManagedFoldersDeleteFiles = Permission{
		ID:            "auto_managed_folders_delete_files",
		DisplayName:   "Automatically managed project folders: Delete files",
		Description:   "Allows users to delete files from automatically managed project folders.",
		PermissibleOn: []Context{ContextProject},
		Module:        ModuleFileStorage,
	}

	AutoManagedFoldersShareFiles = Permission{
		ID:            "auto_managed_folders_share_files",
		DisplayName:   "Automatically managed project folders: Share files",
		Description:   "Allows users to share files from automatically managed project folders.",
		PermissibleOn: []Context{ContextProject},
		Module:        ModuleFileStorage,
	}
)

// ProjectPermissions returns all project-level permissions.
func ProjectPermissions() []Permission {
	return []Permission{
		ArchiveProject,
		EditProject,
		SelectProjectModules,
		ViewProjectAttributes,
		ExportProjects,
		EditProjectAttributes,
		SelectProjectAttributes,
		ViewProjectPhases,
		EditProjectPhases,
		SelectProjectPhases,
		ManageMembers,
		InviteMembersByEmail,
		ViewMembers,
		ManageVersions,
		SelectTypes,
		SelectCustomFields,
		CreateSubprojects,
		CopyProjects,
		ManageDashboards,
		ManageFilesInProject,
		AutoManagedFoldersReadFiles,
		AutoManagedFoldersWriteFiles,
		AutoManagedFoldersCreateFiles,
		AutoManagedFoldersDeleteFiles,
		AutoManagedFoldersShareFiles,
	}
}
