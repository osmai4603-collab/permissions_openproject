package perms

// ──────────────────────────────────────────────
// Wiki permissions
// ──────────────────────────────────────────────

var (
	ViewWiki = Permission{
		ID:          "view_wiki",
		DisplayName: "View wiki",
		Description: "Allows users to view wiki pages.",
		Context:     ContextProject,
		Module:      ModuleWiki,
	}

	ViewWikiHistory = Permission{
		ID:          "view_wiki_history",
		DisplayName: "View wiki history",
		Description: "Allows users to view the edit history of wiki pages.",
		Context:     ContextProject,
		Module:      ModuleWiki,
	}

	EditWikiPages = Permission{
		ID:          "edit_wiki_pages",
		DisplayName: "Edit wiki pages",
		Description: "Allows users to create and edit wiki pages.",
		Context:     ContextProject,
		Module:      ModuleWiki,
	}

	ManageWiki = Permission{
		ID:          "manage_wiki",
		DisplayName: "Manage wiki",
		Description: "Allows users to manage the wiki, including renaming, deleting pages, and managing the wiki menu.",
		Context:     ContextProject,
		Module:      ModuleWiki,
	}
)

// WikiPermissions returns all wiki permissions.
func WikiPermissions() []Permission {
	return []Permission{
		ViewWiki,
		ViewWikiHistory,
		EditWikiPages,
		ManageWiki,
	}
}
