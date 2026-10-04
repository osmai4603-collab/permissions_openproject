package perms

// ──────────────────────────────────────────────
// Forums permissions
// ──────────────────────────────────────────────

var (
	ManageForums = Permission{
		ID:            "manage_forums",
		DisplayName:   "Manage forums",
		Description:   "Allows users to create, edit, and delete forums.",
		PermissibleOn: []Context{ContextProject},
		Module:        ModuleForums,
	}

	PostMessages = Permission{
		ID:            "post_messages",
		DisplayName:   "Post messages",
		Description:   "Allows users to post messages in forums.",
		PermissibleOn: []Context{ContextProject},
		Module:        ModuleForums,
	}

	EditMessages = Permission{
		ID:            "edit_messages",
		DisplayName:   "Edit messages",
		Description:   "Allows users to edit any forum message.",
		PermissibleOn: []Context{ContextProject},
		Module:        ModuleForums,
	}

	EditOwnMessages = Permission{
		ID:            "edit_own_messages",
		DisplayName:   "Edit own messages",
		Description:   "Allows users to edit their own forum messages.",
		PermissibleOn: []Context{ContextProject},
		Module:        ModuleForums,
	}

	DeleteMessages = Permission{
		ID:            "delete_messages",
		DisplayName:   "Delete messages",
		Description:   "Allows users to delete any forum message.",
		PermissibleOn: []Context{ContextProject},
		Module:        ModuleForums,
	}

	DeleteOwnMessages = Permission{
		ID:            "delete_own_messages",
		DisplayName:   "Delete own messages",
		Description:   "Allows users to delete their own forum messages.",
		PermissibleOn: []Context{ContextProject},
		Module:        ModuleForums,
	}
)

// ForumsPermissions returns all forums permissions.
func ForumsPermissions() []Permission {
	return []Permission{
		ManageForums,
		PostMessages,
		EditMessages,
		EditOwnMessages,
		DeleteMessages,
		DeleteOwnMessages,
	}
}
