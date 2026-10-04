package perms

// ──────────────────────────────────────────────
// Boards permissions
// ──────────────────────────────────────────────

var (
	ViewBoards = Permission{
		ID:            "view_boards",
		DisplayName:   "View boards",
		Description:   "Allows users to view boards.",
		PermissibleOn: []Context{ContextProject},
		Module:        ModuleBoards,
	}

	ManageBoards = Permission{
		ID:            "manage_boards",
		DisplayName:   "Manage boards",
		Description:   "Allows users to create, edit, and delete boards.",
		PermissibleOn: []Context{ContextProject},
		Module:        ModuleBoards,
	}
)

// BoardsPermissions returns all boards permissions.
func BoardsPermissions() []Permission {
	return []Permission{
		ViewBoards,
		ManageBoards,
	}
}
