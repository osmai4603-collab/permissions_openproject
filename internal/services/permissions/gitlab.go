package perms

// ──────────────────────────────────────────────
// GitLab integration permissions
// ──────────────────────────────────────────────

var (
	ShowGitLabContent = Permission{
		ID:          "show_gitlab_content",
		DisplayName: "Show GitLab content",
		Description: "Allows users to see GitLab merge requests and issues linked to work packages.",
		Context:     ContextProject,
		Module:      ModuleGitLab,
	}
)

// GitLabPermissions returns all GitLab integration permissions.
func GitLabPermissions() []Permission {
	return []Permission{
		ShowGitLabContent,
	}
}
