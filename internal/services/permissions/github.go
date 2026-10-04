package perms

// ──────────────────────────────────────────────
// GitHub integration permissions
// ──────────────────────────────────────────────

var (
	ShowGitHubContent = Permission{
		ID:          "show_github_content",
		DisplayName: "Show GitHub content",
		Description: "Allows users to see GitHub pull requests and issues linked to work packages.",
		Context:     ContextProject,
		Module:      ModuleGitHub,
	}
)

// GitHubPermissions returns all GitHub integration permissions.
func GitHubPermissions() []Permission {
	return []Permission{
		ShowGitHubContent,
	}
}
