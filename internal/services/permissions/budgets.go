package perms

// ──────────────────────────────────────────────
// Budgets permissions
// ──────────────────────────────────────────────

var (
	ViewBudgets = Permission{
		ID:            "view_budgets",
		DisplayName:   "View budgets",
		Description:   "Allows users to view project budgets.",
		PermissibleOn: []Context{ContextProject},
		Module:        ModuleBudgets,
	}

	EditBudgets = Permission{
		ID:            "edit_budgets",
		DisplayName:   "Edit budgets",
		Description:   "Allows users to create, edit, and delete project budgets.",
		PermissibleOn: []Context{ContextProject},
		Module:        ModuleBudgets,
	}
)

// BudgetsPermissions returns all budgets permissions.
func BudgetsPermissions() []Permission {
	return []Permission{
		ViewBudgets,
		EditBudgets,
	}
}
