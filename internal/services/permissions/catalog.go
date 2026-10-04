package perms

// ──────────────────────────────────────────────
// Standard OpenProject Permission IDs
// ──────────────────────────────────────────────

const (
	// Global (Platform-wide) permissions
	PermAddProject                 = "add_project"
	PermCreateUser                 = "create_user"
	PermEditUsers                  = "edit_users"
	PermCreateBackup               = "create_backup"
	PermViewAllUsersAndGroups       = "view_all_users_and_groups"
	PermManagePublicProjectLists   = "manage_public_project_lists"
	PermCreatePortfolios           = "create_portfolios"
	PermCreatePrograms             = "create_programs"
	PermEditAttributeHelpTexts     = "edit_attribute_help_texts"
	PermManagePlaceholderUsers     = "manage_placeholder_users"
	PermViewAllUsersMailAddresses  = "view_all_users_mail_addresses"

	// Project Module permissions
	PermArchiveProject               = "archive_project"
	PermEditProject                  = "edit_project"
	PermSelectProjectModules         = "select_project_modules"
	PermViewProjectAttributes        = "view_project_attributes"
	PermExportProjects               = "export_projects"
	PermEditProjectAttributes        = "edit_project_attributes"
	PermSelectProjectAttributes      = "select_project_attributes"
	PermViewProjectPhases            = "view_project_phases"
	PermEditProjectPhases            = "edit_project_phases"
	PermSelectProjectPhases          = "select_project_phases"
	PermManageMembers                = "manage_members"
	PermInviteMembersByEmail         = "invite_members_by_email"
	PermViewMembers                  = "view_members"
	PermManageVersions               = "manage_versions"
	PermSelectTypes                  = "select_types"
	PermSelectCustomFields           = "select_custom_fields"
	PermCreateSubprojects            = "create_subprojects"
	PermCopyProjects                 = "copy_projects"
	PermManageDashboards             = "manage_dashboards"
	PermManageFilesInProject         = "manage_files_in_project"
	PermAutoManagedFoldersReadFiles  = "auto_managed_folders_read_files"
	PermAutoManagedFoldersWriteFiles = "auto_managed_folders_write_files"
	PermAutoManagedFoldersCreateFiles = "auto_managed_folders_create_files"
	PermAutoManagedFoldersDeleteFiles = "auto_managed_folders_delete_files"
	PermAutoManagedFoldersShareFiles = "auto_managed_folders_share_files"

	// Work Packages Module permissions
	PermViewWorkPackages             = "view_work_packages"
	PermAddWorkPackages              = "add_work_packages"
	PermEditWorkPackages             = "edit_work_packages"
	PermMoveWorkPackages             = "move_work_packages"
	PermDuplicateWorkPackages        = "duplicate_work_packages"
	PermAddComments                  = "add_comments"
	PermEditOwnComments              = "edit_own_comments"
	PermModerateComments             = "moderate_comments"
	PermViewInternalComments         = "view_internal_comments"
	PermWriteInternalComments        = "write_internal_comments"
	PermEditOwnInternalComments      = "edit_own_internal_comments"
	PermModerateInternalComments     = "moderate_internal_comments"
	PermAddAttachments               = "add_attachments"
	PermManageWorkPackageCategories  = "manage_work_package_categories"
	PermExportWorkPackages           = "export_work_packages"
	PermDeleteWorkPackages           = "delete_work_packages"
	PermManageWorkPackageRelations   = "manage_work_package_relations"
	PermManageWorkPackageHierarchies = "manage_work_package_hierarchies"
	PermManagePublicViews            = "manage_public_views"
	PermSaveViews                    = "save_views"
	PermViewWatchersList             = "view_watchers_list"
	PermAddWatchers                  = "add_watchers"
	PermDeleteWatchers               = "delete_watchers"
	PermShareWorkPackages            = "share_work_packages"
	PermViewWorkPackageShares        = "view_work_package_shares"
	PermAssignVersions               = "assign_versions"
	PermChangeWorkPackageStatus      = "change_work_package_status"
	PermBecomeAssigneeResponsible    = "become_assignee_responsible"
	PermViewFileLinks                = "view_file_links"
	PermManageFileLinks              = "manage_file_links"
	PermManageWikiPageLinks          = "manage_wiki_page_links"

	// Boards Module permissions
	PermViewBoards   = "view_boards"
	PermManageBoards = "manage_boards"

	// Backlogs Module permissions
	PermViewSprints                   = "view_sprints"
	PermSelectBacklogTypesAndStatuses = "select_backlog_types_and_statuses"
	PermCreateSprints                 = "create_sprints"
	PermStartCompleteSprint           = "start_complete_sprint"
	PermManageSprintItems             = "manage_sprint_items"
	PermShareSprint                   = "share_sprint"

	// Budgets Module permissions
	PermViewBudgets = "view_budgets"
	PermEditBudgets = "edit_budgets"

	// Calendars Module permissions
	PermViewCalendars         = "view_calendars"
	PermEditCalendars         = "edit_calendars"
	PermSubscribeToIcalendars = "subscribe_to_icalendars"

	// Documents Module permissions
	PermViewDocuments   = "view_documents"
	PermManageDocuments = "manage_documents"

	// Forums Module permissions
	PermManageForums     = "manage_forums"
	PermPostMessages     = "post_messages"
	PermEditMessages     = "edit_messages"
	PermEditOwnMessages  = "edit_own_messages"
	PermDeleteMessages   = "delete_messages"
	PermDeleteOwnMessages = "delete_own_messages"

	// GitHub Integration Module permissions
	PermShowGithubContent = "show_github_content"

	// GitLab Integration Module permissions
	PermShowGitlabContent = "show_gitlab_content"

	// Meetings Module permissions
	PermViewMeetings       = "view_meetings"
	PermCreateMeetings     = "create_meetings"
	PermEditMeetings       = "edit_meetings"
	PermDeleteMeetings     = "delete_meetings"
	PermSendMeetingInvites = "send_meeting_invites"
	PermManageAgendas      = "manage_agendas"
	PermManageOutcomes     = "manage_outcomes"

	// News Module permissions
	PermManageNews  = "manage_news"
	PermCommentNews = "comment_news"

	// Team Planner Module permissions
	PermViewTeamPlanner   = "view_team_planner"
	PermManageTeamPlanner = "manage_team_planner"

	// Time & Costs Module permissions
	PermViewSpentTime             = "view_spent_time"
	PermViewOwnSpentTime          = "view_own_spent_time"
	PermLogOwnTime                = "log_own_time"
	PermLogTimeForOtherUsers      = "log_time_for_other_users"
	PermEditOwnTimeLogs           = "edit_own_time_logs"
	PermEditTimeLogsForOtherUsers = "edit_time_logs_for_other_users"
	PermManageProjectActivities   = "manage_project_activities"
	PermViewOwnHourlyRate         = "view_own_hourly_rate"
	PermViewAllHourlyRates        = "view_all_hourly_rates"
	PermEditOwnHourlyRates        = "edit_own_hourly_rates"
	PermEditHourlyRates           = "edit_hourly_rates"
	PermViewCostRates             = "view_cost_rates"
	PermBookUnitCostsForOneself   = "book_unit_costs_for_oneself"
	PermBookUnitCosts             = "book_unit_costs"
	PermEditOwnBookedUnitCosts    = "edit_own_booked_unit_costs"
	PermEditBookedUnitCosts       = "edit_booked_unit_costs"
	PermViewBookedCosts           = "view_booked_costs"
	PermViewOwnBookedCosts        = "view_own_booked_costs"
	PermSavePublicCostReports     = "save_public_cost_reports"
	PermSavePrivateCostReports    = "save_private_cost_reports"

	// Wiki Module permissions
	PermViewWiki        = "view_wiki"
	PermViewWikiHistory = "view_wiki_history"
	PermEditWikiPages   = "edit_wiki_pages"
	PermManageWiki      = "manage_wiki"
)
