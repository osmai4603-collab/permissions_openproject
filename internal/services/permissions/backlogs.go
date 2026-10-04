package perms

// ──────────────────────────────────────────────
// Backlogs (Scrum) permissions
// ──────────────────────────────────────────────

var (
	ViewSprints = Permission{
		ID:          "view_sprints",
		DisplayName: "View sprints",
		Description: "Allows users to view sprints.",
		Context:     ContextProject,
		Module:      ModuleBacklogs,
	}

	SelectBacklogTypesAndStatuses = Permission{
		ID:          "select_backlog_types_and_statuses",
		DisplayName: "Select backlog types and statuses",
		Description: "Allows users to configure which work package types appear in the backlog and which statuses are considered completed.",
		Context:     ContextProject,
		Module:      ModuleBacklogs,
	}

	CreateSprints = Permission{
		ID:          "create_sprints",
		DisplayName: "Create sprints",
		Description: "Allows users to create sprints.",
		Context:     ContextProject,
		Module:      ModuleBacklogs,
	}

	StartCompleteSprint = Permission{
		ID:          "start_complete_sprint",
		DisplayName: "Start/complete sprint",
		Description: "Allows users to start and complete sprints.",
		Context:     ContextProject,
		Module:      ModuleBacklogs,
	}

	ManageSprintItems = Permission{
		ID:          "manage_sprint_items",
		DisplayName: "Manage sprint items",
		Description: "Allows users to manage sprint contents.",
		Context:     ContextProject,
		Module:      ModuleBacklogs,
	}

	ShareSprint = Permission{
		ID:          "share_sprint",
		DisplayName: "Share sprint",
		Description: "Allows users to share sprint information.",
		Context:     ContextProject,
		Module:      ModuleBacklogs,
	}
)

// BacklogsPermissions returns all backlogs (Scrum) permissions.
func BacklogsPermissions() []Permission {
	return []Permission{
		ViewSprints,
		SelectBacklogTypesAndStatuses,
		CreateSprints,
		StartCompleteSprint,
		ManageSprintItems,
		ShareSprint,
	}
}
