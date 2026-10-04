package perms

// ──────────────────────────────────────────────
// Time tracking and cost reporting permissions
// ──────────────────────────────────────────────

var (
	ViewSpentTime = Permission{
		ID:            "view_spent_time",
		DisplayName:   "View spent time",
		Description:   "Allows users to view time entries logged by all users.",
		PermissibleOn: []Context{ContextProject},
		Module:        ModuleTimeCosts,
	}

	ViewOwnSpentTime = Permission{
		ID:            "view_own_spent_time",
		DisplayName:   "View own spent time",
		Description:   "Allows users to view only their own time entries.",
		PermissibleOn: []Context{ContextProject},
		Module:        ModuleTimeCosts,
	}

	LogOwnTime = Permission{
		ID:            "log_own_time",
		DisplayName:   "Log own time",
		Description:   "Allows users to log time on work packages for themselves.",
		PermissibleOn: []Context{ContextProject},
		Module:        ModuleTimeCosts,
	}

	LogTimeForOtherUsers = Permission{
		ID:            "log_time_for_other_users",
		DisplayName:   "Log time for other users",
		Description:   "Allows users to log time entries on behalf of other users.",
		PermissibleOn: []Context{ContextProject},
		Module:        ModuleTimeCosts,
	}

	EditOwnTimeLogs = Permission{
		ID:            "edit_own_time_logs",
		DisplayName:   "Edit own time logs",
		Description:   "Allows users to edit their own logged time entries.",
		PermissibleOn: []Context{ContextProject},
		Module:        ModuleTimeCosts,
	}

	EditTimeLogsForOtherUsers = Permission{
		ID:            "edit_time_logs_for_other_users",
		DisplayName:   "Edit time logs for other users",
		Description:   "Allows users to edit time entries logged by other users.",
		PermissibleOn: []Context{ContextProject},
		Module:        ModuleTimeCosts,
	}

	ManageProjectActivities = Permission{
		ID:            "manage_project_activities",
		DisplayName:   "Manage project activities",
		Description:   "Allows users to manage time tracking activity types within a project.",
		PermissibleOn: []Context{ContextProject},
		Module:        ModuleTimeCosts,
	}

	ViewOwnHourlyRate = Permission{
		ID:            "view_own_hourly_rate",
		DisplayName:   "View own hourly rate",
		Description:   "Allows users to view their own hourly labor rate.",
		PermissibleOn: []Context{ContextProject},
		Module:        ModuleTimeCosts,
	}

	ViewAllHourlyRates = Permission{
		ID:            "view_all_hourly_rates",
		DisplayName:   "View all hourly rates",
		Description:   "Allows users to view hourly rates of all members.",
		PermissibleOn: []Context{ContextProject},
		Module:        ModuleTimeCosts,
	}

	EditOwnHourlyRates = Permission{
		ID:            "edit_own_hourly_rates",
		DisplayName:   "Edit own hourly rates",
		Description:   "Allows users to edit their own hourly labor rate.",
		PermissibleOn: []Context{ContextProject},
		Module:        ModuleTimeCosts,
	}

	EditHourlyRates = Permission{
		ID:            "edit_hourly_rates",
		DisplayName:   "Edit hourly rates",
		Description:   "Allows users to edit hourly rates for all members.",
		PermissibleOn: []Context{ContextProject},
		Module:        ModuleTimeCosts,
	}

	ViewCostRates = Permission{
		ID:            "view_cost_rates",
		DisplayName:   "View cost rates",
		Description:   "Allows users to view cost types and their rates.",
		PermissibleOn: []Context{ContextProject},
		Module:        ModuleTimeCosts,
	}

	BookUnitCostsForOneself = Permission{
		ID:            "book_unit_costs_for_oneself",
		DisplayName:   "Book unit costs for oneself",
		Description:   "Allows users to book unit costs on work packages for themselves.",
		PermissibleOn: []Context{ContextProject},
		Module:        ModuleTimeCosts,
	}

	BookUnitCosts = Permission{
		ID:            "book_unit_costs",
		DisplayName:   "Book unit costs",
		Description:   "Allows users to book unit costs on work packages for all members.",
		PermissibleOn: []Context{ContextProject},
		Module:        ModuleTimeCosts,
	}

	EditOwnBookedUnitCosts = Permission{
		ID:            "edit_own_booked_unit_costs",
		DisplayName:   "Edit own booked unit costs",
		Description:   "Allows users to edit their own booked unit cost entries.",
		PermissibleOn: []Context{ContextProject},
		Module:        ModuleTimeCosts,
	}

	EditBookedUnitCosts = Permission{
		ID:            "edit_booked_unit_costs",
		DisplayName:   "Edit booked unit costs",
		Description:   "Allows users to edit unit cost entries booked by any member.",
		PermissibleOn: []Context{ContextProject},
		Module:        ModuleTimeCosts,
	}

	ViewBookedCosts = Permission{
		ID:            "view_booked_costs",
		DisplayName:   "View booked costs",
		Description:   "Allows users to view booked costs for all members.",
		PermissibleOn: []Context{ContextProject},
		Module:        ModuleTimeCosts,
	}

	ViewOwnBookedCosts = Permission{
		ID:            "view_own_booked_costs",
		DisplayName:   "View own booked costs",
		Description:   "Allows users to view only their own booked costs.",
		PermissibleOn: []Context{ContextProject},
		Module:        ModuleTimeCosts,
	}

	SavePublicCostReports = Permission{
		ID:            "save_public_cost_reports",
		DisplayName:   "Save public cost reports",
		Description:   "Allows users to save cost reports visible to all project members.",
		PermissibleOn: []Context{ContextProject},
		Module:        ModuleTimeCosts,
	}

	SavePrivateCostReports = Permission{
		ID:            "save_private_cost_reports",
		DisplayName:   "Save private cost reports",
		Description:   "Allows users to save cost reports visible only to themselves.",
		PermissibleOn: []Context{ContextProject},
		Module:        ModuleTimeCosts,
	}
)

// TimeCostsPermissions returns all time tracking and cost reporting permissions.
func TimeCostsPermissions() []Permission {
	return []Permission{
		ViewSpentTime,
		ViewOwnSpentTime,
		LogOwnTime,
		LogTimeForOtherUsers,
		EditOwnTimeLogs,
		EditTimeLogsForOtherUsers,
		ManageProjectActivities,
		ViewOwnHourlyRate,
		ViewAllHourlyRates,
		EditOwnHourlyRates,
		EditHourlyRates,
		ViewCostRates,
		BookUnitCostsForOneself,
		BookUnitCosts,
		EditOwnBookedUnitCosts,
		EditBookedUnitCosts,
		ViewBookedCosts,
		ViewOwnBookedCosts,
		SavePublicCostReports,
		SavePrivateCostReports,
	}
}
