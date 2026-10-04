package perms

// ──────────────────────────────────────────────
// Team planner permissions
// ──────────────────────────────────────────────

var (
	ViewTeamPlanner = Permission{
		ID:          "view_team_planner",
		DisplayName: "View team planner",
		Description: "Allows users to view team planner views.",
		Context:     ContextProject,
		Module:      ModuleTeamPlanner,
	}

	ManageTeamPlanner = Permission{
		ID:          "manage_team_planner",
		DisplayName: "Manage team planner",
		Description: "Allows users to create, edit, and delete team planner views.",
		Context:     ContextProject,
		Module:      ModuleTeamPlanner,
	}
)

// TeamPlannerPermissions returns all team planner permissions.
func TeamPlannerPermissions() []Permission {
	return []Permission{
		ViewTeamPlanner,
		ManageTeamPlanner,
	}
}
