package perms

// ──────────────────────────────────────────────
// Calendars permissions
// ──────────────────────────────────────────────

var (
	ViewCalendars = Permission{
		ID:          "view_calendars",
		DisplayName: "View calendars",
		Description: "Allows users to view calendars.",
		Context:     ContextProject,
		Module:      ModuleCalendars,
	}

	EditCalendars = Permission{
		ID:          "edit_calendars",
		DisplayName: "Edit calendars",
		Description: "Allows users to create, edit, and delete calendars.",
		Context:     ContextProject,
		Module:      ModuleCalendars,
	}

	SubscribeToICalendars = Permission{
		ID:          "subscribe_to_icalendars",
		DisplayName: "Subscribe to iCalendars",
		Description: "Allows users to subscribe to calendar feeds.",
		Context:     ContextProject,
		Module:      ModuleCalendars,
	}
)

// CalendarsPermissions returns all calendars permissions.
func CalendarsPermissions() []Permission {
	return []Permission{
		ViewCalendars,
		EditCalendars,
		SubscribeToICalendars,
	}
}
