package perms

// ──────────────────────────────────────────────
// Work packages and Gantt chart permissions
// ──────────────────────────────────────────────

var (
	ViewWorkPackages = Permission{
		ID:            "view_work_packages",
		DisplayName:   "View work packages",
		Description:   "Allows users to view work packages.",
		PermissibleOn: []Context{ContextProject, ContextWorkPackage},
		Module:        ModuleWorkPackages,
	}

	AddWorkPackages = Permission{
		ID:            "add_work_packages",
		DisplayName:   "Add work packages",
		Description:   "Allows users to create work packages.",
		PermissibleOn: []Context{ContextProject},
		Module:        ModuleWorkPackages,
	}

	EditWorkPackages = Permission{
		ID:            "edit_work_packages",
		DisplayName:   "Edit work packages",
		Description:   "Allows users to edit work packages.",
		PermissibleOn: []Context{ContextProject, ContextWorkPackage},
		Module:        ModuleWorkPackages,
	}

	MoveWorkPackages = Permission{
		ID:            "move_work_packages",
		DisplayName:   "Move work packages",
		Description:   "Allows users to move work packages between projects.",
		PermissibleOn: []Context{ContextProject},
		Module:        ModuleWorkPackages,
	}

	DuplicateWorkPackages = Permission{
		ID:            "duplicate_work_packages",
		DisplayName:   "Duplicate work packages",
		Description:   "Allows users to duplicate work packages.",
		PermissibleOn: []Context{ContextProject},
		Module:        ModuleWorkPackages,
	}

	AddComments = Permission{
		ID:            "add_comments",
		DisplayName:   "Add comments",
		Description:   "Allows users to comment on work packages.",
		PermissibleOn: []Context{ContextProject, ContextWorkPackage},
		Module:        ModuleWorkPackages,
	}

	EditOwnComments = Permission{
		ID:            "edit_own_comments",
		DisplayName:   "Edit own comments",
		Description:   "Allows users to edit their own comments.",
		PermissibleOn: []Context{ContextProject},
		Module:        ModuleWorkPackages,
	}

	ModerateComments = Permission{
		ID:            "moderate_comments",
		DisplayName:   "Moderate comments",
		Description:   "Allows users to edit comments created by any user.",
		PermissibleOn: []Context{ContextProject},
		Module:        ModuleWorkPackages,
	}

	ViewInternalComments = Permission{
		ID:            "view_internal_comments",
		DisplayName:   "View internal comments",
		Description:   "Allows users to view internal comments.",
		PermissibleOn: []Context{ContextProject},
		Module:        ModuleWorkPackages,
	}

	WriteInternalComments = Permission{
		ID:            "write_internal_comments",
		DisplayName:   "Write internal comments",
		Description:   "Allows users to create internal comments.",
		PermissibleOn: []Context{ContextProject},
		Module:        ModuleWorkPackages,
	}

	EditOwnInternalComments = Permission{
		ID:            "edit_own_internal_comments",
		DisplayName:   "Edit own internal comments",
		Description:   "Allows users to edit their own internal comments.",
		PermissibleOn: []Context{ContextProject},
		Module:        ModuleWorkPackages,
	}

	ModerateInternalComments = Permission{
		ID:            "moderate_internal_comments",
		DisplayName:   "Moderate internal comments",
		Description:   "Allows users to edit internal comments created by any user.",
		PermissibleOn: []Context{ContextProject},
		Module:        ModuleWorkPackages,
	}

	AddAttachments = Permission{
		ID:            "add_attachments",
		DisplayName:   "Add attachments",
		Description:   "Allows users to upload attachments to work packages.",
		PermissibleOn: []Context{ContextProject, ContextWorkPackage},
		Module:        ModuleWorkPackages,
		Dependencies: []Dependency{
			{Description: "Can be granted independently of Edit work packages."},
		},
	}

	ManageWorkPackageCategories = Permission{
		ID:            "manage_work_package_categories",
		DisplayName:   "Manage work package categories",
		Description:   "Allows users to create, edit, and delete work package categories.",
		PermissibleOn: []Context{ContextProject},
		Module:        ModuleWorkPackages,
	}

	ExportWorkPackages = Permission{
		ID:            "export_work_packages",
		DisplayName:   "Export work packages",
		Description:   "Allows users to export work packages.",
		PermissibleOn: []Context{ContextProject},
		Module:        ModuleWorkPackages,
	}

	DeleteWorkPackages = Permission{
		ID:            "delete_work_packages",
		DisplayName:   "Delete work packages",
		Description:   "Allows users to delete work packages.",
		PermissibleOn: []Context{ContextProject},
		Module:        ModuleWorkPackages,
	}

	ManageWorkPackageRelations = Permission{
		ID:            "manage_work_package_relations",
		DisplayName:   "Manage work package relations",
		Description:   "Allows users to create, edit, and remove work package relations.",
		PermissibleOn: []Context{ContextProject},
		Module:        ModuleWorkPackages,
	}

	ManageWorkPackageHierarchies = Permission{
		ID:            "manage_work_package_hierarchies",
		DisplayName:   "Manage work package hierarchies",
		Description:   "Allows users to manage parent-child relationships between work packages.",
		PermissibleOn: []Context{ContextProject},
		Module:        ModuleWorkPackages,
	}

	ManagePublicViews = Permission{
		ID:            "manage_public_views",
		DisplayName:   "Manage public views",
		Description:   "Allows users to create, edit, and delete public work package views.",
		PermissibleOn: []Context{ContextProject},
		Module:        ModuleWorkPackages,
	}

	SaveViews = Permission{
		ID:            "save_views",
		DisplayName:   "Save views",
		Description:   "Allows users to save personal work package views.",
		PermissibleOn: []Context{ContextProject},
		Module:        ModuleWorkPackages,
	}

	ViewWatchersList = Permission{
		ID:            "view_watchers_list",
		DisplayName:   "View watchers list",
		Description:   "Allows users to see who is watching a work package.",
		PermissibleOn: []Context{ContextProject},
		Module:        ModuleWorkPackages,
	}

	AddWatchers = Permission{
		ID:            "add_watchers",
		DisplayName:   "Add watchers",
		Description:   "Allows users to add watchers to work packages.",
		PermissibleOn: []Context{ContextProject},
		Module:        ModuleWorkPackages,
	}

	DeleteWatchers = Permission{
		ID:            "delete_watchers",
		DisplayName:   "Delete watchers",
		Description:   "Allows users to remove watchers from work packages.",
		PermissibleOn: []Context{ContextProject},
		Module:        ModuleWorkPackages,
	}

	ShareWorkPackages = Permission{
		ID:            "share_work_packages",
		DisplayName:   "Share work packages",
		Description:   "Allows users to share work packages with other users.",
		PermissibleOn: []Context{ContextProject},
		Module:        ModuleWorkPackages,
	}

	ViewWorkPackageShares = Permission{
		ID:            "view_work_package_shares",
		DisplayName:   "View work package shares",
		Description:   "Allows users to view existing work package shares.",
		PermissibleOn: []Context{ContextProject},
		Module:        ModuleWorkPackages,
	}

	AssignVersions = Permission{
		ID:            "assign_versions",
		DisplayName:   "Assign versions",
		Description:   "Allows users to assign versions to work packages.",
		PermissibleOn: []Context{ContextProject},
		Module:        ModuleWorkPackages,
	}

	ChangeWorkPackageStatus = Permission{
		ID:            "change_work_package_status",
		DisplayName:   "Change work package status",
		Description:   "Allows users to change the status of work packages.",
		PermissibleOn: []Context{ContextProject},
		Module:        ModuleWorkPackages,
		Dependencies: []Dependency{
			{Description: "Can be granted independently of Edit work packages."},
		},
	}

	BecomeAssigneeResponsible = Permission{
		ID:            "become_assignee_responsible",
		DisplayName:   "Become assignee/responsible",
		Description:   "Allows work packages to be assigned to users or groups with this role in the project.",
		PermissibleOn: []Context{ContextProject},
		Module:        ModuleWorkPackages,
		Dependencies: []Dependency{
			{Description: "Controls whether a user or group can be selected as assignee or responsible for work packages."},
		},
	}

	ViewFileLinks = Permission{
		ID:            "view_file_links",
		DisplayName:   "View file links",
		Description:   "Allows users to view file links attached to work packages.",
		PermissibleOn: []Context{ContextProject, ContextWorkPackage},
		Module:        ModuleWorkPackages,
	}

	ManageFileLinks = Permission{
		ID:            "manage_file_links",
		DisplayName:   "Manage file links",
		Description:   "Allows users to create, edit, and remove file links.",
		PermissibleOn: []Context{ContextProject, ContextWorkPackage},
		Module:        ModuleWorkPackages,
	}

	ManageWikiPageLinks = Permission{
		ID:            "manage_wiki_page_links",
		DisplayName:   "Manage wiki page links",
		Description:   "Allows users to create, edit, and remove wiki page links.",
		PermissibleOn: []Context{ContextProject},
		Module:        ModuleWorkPackages,
	}
)

// WorkPackagePermissions returns all work packages and Gantt chart permissions.
func WorkPackagePermissions() []Permission {
	return []Permission{
		ViewWorkPackages,
		AddWorkPackages,
		EditWorkPackages,
		MoveWorkPackages,
		DuplicateWorkPackages,
		AddComments,
		EditOwnComments,
		ModerateComments,
		ViewInternalComments,
		WriteInternalComments,
		EditOwnInternalComments,
		ModerateInternalComments,
		AddAttachments,
		ManageWorkPackageCategories,
		ExportWorkPackages,
		DeleteWorkPackages,
		ManageWorkPackageRelations,
		ManageWorkPackageHierarchies,
		ManagePublicViews,
		SaveViews,
		ViewWatchersList,
		AddWatchers,
		DeleteWatchers,
		ShareWorkPackages,
		ViewWorkPackageShares,
		AssignVersions,
		ChangeWorkPackageStatus,
		BecomeAssigneeResponsible,
		ViewFileLinks,
		ManageFileLinks,
		ManageWikiPageLinks,
	}
}
