package perms

// ──────────────────────────────────────────────
// Meetings permissions
// ──────────────────────────────────────────────

var (
	ViewMeetings = Permission{
		ID:            "view_meetings",
		DisplayName:   "View meetings",
		Description:   "Allows users to view meetings in a project.",
		PermissibleOn: []Context{ContextProject},
		Module:        ModuleMeetings,
	}

	CreateMeetings = Permission{
		ID:            "create_meetings",
		DisplayName:   "Create meetings",
		Description:   "Allows users to create new meetings.",
		PermissibleOn: []Context{ContextProject},
		Module:        ModuleMeetings,
	}

	EditMeetings = Permission{
		ID:            "edit_meetings",
		DisplayName:   "Edit meetings",
		Description:   "Allows users to edit existing meetings.",
		PermissibleOn: []Context{ContextProject},
		Module:        ModuleMeetings,
	}

	DeleteMeetings = Permission{
		ID:            "delete_meetings",
		DisplayName:   "Delete meetings",
		Description:   "Allows users to delete meetings.",
		PermissibleOn: []Context{ContextProject},
		Module:        ModuleMeetings,
	}

	SendMeetingInvites = Permission{
		ID:            "send_meeting_invites",
		DisplayName:   "Send meeting invites",
		Description:   "Allows users to send email invitations for meetings to participants.",
		PermissibleOn: []Context{ContextProject},
		Module:        ModuleMeetings,
	}

	ManageAgendas = Permission{
		ID:            "manage_agendas",
		DisplayName:   "Manage agendas",
		Description:   "Allows users to create, edit, and close meeting agendas.",
		PermissibleOn: []Context{ContextProject},
		Module:        ModuleMeetings,
	}

	ManageOutcomes = Permission{
		ID:            "manage_outcomes",
		DisplayName:   "Manage outcomes",
		Description:   "Allows users to create, edit, and close meeting outcomes/minutes.",
		PermissibleOn: []Context{ContextProject},
		Module:        ModuleMeetings,
	}
)

// MeetingsPermissions returns all meetings permissions.
func MeetingsPermissions() []Permission {
	return []Permission{
		ViewMeetings,
		CreateMeetings,
		EditMeetings,
		DeleteMeetings,
		SendMeetingInvites,
		ManageAgendas,
		ManageOutcomes,
	}
}
