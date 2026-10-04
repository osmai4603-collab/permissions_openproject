package perms

// ──────────────────────────────────────────────
// Documents permissions
// ──────────────────────────────────────────────

var (
	ViewDocuments = Permission{
		ID:          "view_documents",
		DisplayName: "View documents",
		Description: "Allows users to view project documents.",
		Context:     ContextProject,
		Module:      ModuleDocuments,
	}

	ManageDocuments = Permission{
		ID:          "manage_documents",
		DisplayName: "Manage documents",
		Description: "Allows users to create, edit, and delete project documents.",
		Context:     ContextProject,
		Module:      ModuleDocuments,
	}
)

// DocumentsPermissions returns all documents permissions.
func DocumentsPermissions() []Permission {
	return []Permission{
		ViewDocuments,
		ManageDocuments,
	}
}
