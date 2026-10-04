package perms

// ──────────────────────────────────────────────
// News permissions
// ──────────────────────────────────────────────

var (
	ManageNews = Permission{
		ID:          "manage_news",
		DisplayName: "Manage news",
		Description: "Allows users to create, edit, and delete news entries.",
		Context:     ContextProject,
		Module:      ModuleNews,
	}

	CommentNews = Permission{
		ID:          "comment_news",
		DisplayName: "Comment news",
		Description: "Allows users to post comments on news entries.",
		Context:     ContextProject,
		Module:      ModuleNews,
	}
)

// NewsPermissions returns all news permissions.
func NewsPermissions() []Permission {
	return []Permission{
		ManageNews,
		CommentNews,
	}
}
