package perms

import "sort"

// Standard OpenProject Permission Keys as defined across OpenProject Core and Modules
const (

	// ModuleSystem
	PermAddProject = "add_project"
	PermCreateUser = "create_user"
	PermEditUsers = "edit_users"
	PermDeleteUser = "delete_user"
	PermCreateBackup = "create_backup"
	PermRestoreBackup = "restore_backup"
	PermViewAllUsersAndGroups = "view_all_users_and_groups"
	PermManagePublicProjectLists = "manage_public_project_lists"
	PermManageGlobalCustomFields = "manage_global_custom_fields"
	PermManageRoles = "manage_roles"

	// ModuleProject
	PermArchiveProject = "archive_project"
	PermEditProject = "edit_project"
	PermSelectProjectModules = "select_project_modules"
	PermViewProjectAttributes = "view_project_attributes"
	PermExportProjects = "export_projects"
	PermEditProjectAttributes = "edit_project_attributes"
	PermSelectProjectAttributes = "select_project_attributes"
	PermViewProjectPhases = "view_project_phases"
	PermEditProjectPhases = "edit_project_phases"
	PermSelectProjectPhases = "select_project_phases"
	PermManageMembers = "manage_members"
	PermInviteMembersByEmail = "invite_members_by_email"
	PermViewMembers = "view_members"
	PermManageVersions = "manage_versions"
	PermSelectTypes = "select_types"
	PermSelectCustomFields = "select_custom_fields"
	PermCreateSubprojects = "create_subprojects"
	PermCopyProjects = "copy_projects"
	PermManageDashboards = "manage_dashboards"
	PermManageFilesInProject = "manage_files_in_project"
	PermAutoManagedFoldersReadFiles = "auto_managed_folders_read_files"
	PermAutoManagedFoldersWriteFiles = "auto_managed_folders_write_files"
	PermAutoManagedFoldersCreateFiles = "auto_managed_folders_create_files"
	PermAutoManagedFoldersDeleteFiles = "auto_managed_folders_delete_files"
	PermAutoManagedFoldersShareFiles = "auto_managed_folders_share_files"

	// ModuleWorkPackages
	PermViewWorkPackages = "view_work_packages"
	PermAddWorkPackages = "add_work_packages"
	PermEditWorkPackages = "edit_work_packages"
	PermMoveWorkPackages = "move_work_packages"
	PermDuplicateWorkPackages = "duplicate_work_packages"
	PermAddComments = "add_comments"
	PermEditOwnComments = "edit_own_comments"
	PermModerateComments = "moderate_comments"
	PermViewInternalComments = "view_internal_comments"
	PermWriteInternalComments = "write_internal_comments"
	PermEditOwnInternalComments = "edit_own_internal_comments"
	PermModerateInternalComments = "moderate_internal_comments"
	PermAddAttachments = "add_attachments"
	PermManageWorkPackageCategories = "manage_work_package_categories"
	PermExportWorkPackages = "export_work_packages"
	PermDeleteWorkPackages = "delete_work_packages"
	PermManageWorkPackageRelations = "manage_work_package_relations"
	PermManageWorkPackageHierarchies = "manage_work_package_hierarchies"
	PermManagePublicViews = "manage_public_views"
	PermSaveViews = "save_views"
	PermViewWatchersList = "view_watchers_list"
	PermAddWatchers = "add_watchers"
	PermDeleteWatchers = "delete_watchers"
	PermShareWorkPackages = "share_work_packages"
	PermViewWorkPackageShares = "view_work_package_shares"
	PermAssignVersions = "assign_versions"
	PermChangeWorkPackageStatus = "change_work_package_status"
	PermBecomeAssigneeResponsible = "become_assignee_responsible"
	PermViewFileLinks = "view_file_links"
	PermManageFileLinks = "manage_file_links"
	PermManageWikiPageLinks = "manage_wiki_page_links"

	// ModuleBoards
	PermViewBoards = "view_boards"
	PermManageBoards = "manage_boards"

	// ModuleBacklogs
	PermViewSprints = "view_sprints"
	PermSelectBacklogTypesAndStatuses = "select_backlog_types_and_statuses"
	PermCreateSprints = "create_sprints"
	PermStartCompleteSprint = "start_complete_sprint"
	PermManageSprintItems = "manage_sprint_items"
	PermShareSprint = "share_sprint"

	// ModuleBudgets
	PermViewBudgets = "view_budgets"
	PermEditBudgets = "edit_budgets"

	// ModuleCalendars
	PermViewCalendars = "view_calendars"
	PermEditCalendars = "edit_calendars"
	PermSubscribeToIcalendars = "subscribe_to_icalendars"

	// ModuleDocuments
	PermViewDocuments = "view_documents"
	PermManageDocuments = "manage_documents"

	// ModuleForums
	PermManageForums = "manage_forums"
	PermPostMessages = "post_messages"
	PermEditMessages = "edit_messages"
	PermEditOwnMessages = "edit_own_messages"
	PermDeleteMessages = "delete_messages"
	PermDeleteOwnMessages = "delete_own_messages"

	// ModuleGitHub
	PermShowGithubContent = "show_github_content"

	// ModuleGitLab
	PermShowGitlabContent = "show_gitlab_content"

	// ModuleMeetings
	PermViewMeetings = "view_meetings"
	PermCreateMeetings = "create_meetings"
	PermEditMeetings = "edit_meetings"
	PermDeleteMeetings = "delete_meetings"
	PermSendMeetingInvites = "send_meeting_invites"
	PermManageAgendas = "manage_agendas"
	PermManageOutcomes = "manage_outcomes"

	// ModuleNews
	PermManageNews = "manage_news"
	PermCommentNews = "comment_news"

	// ModuleTeamPlanner
	PermViewTeamPlanner = "view_team_planner"
	PermManageTeamPlanner = "manage_team_planner"

	// ModuleTimeCosts
	PermViewSpentTime = "view_spent_time"
	PermViewOwnSpentTime = "view_own_spent_time"
	PermLogOwnTime = "log_own_time"
	PermLogTimeForOtherUsers = "log_time_for_other_users"
	PermEditOwnTimeLogs = "edit_own_time_logs"
	PermEditTimeLogsForOtherUsers = "edit_time_logs_for_other_users"
	PermManageProjectActivities = "manage_project_activities"
	PermViewOwnHourlyRate = "view_own_hourly_rate"
	PermViewAllHourlyRates = "view_all_hourly_rates"
	PermEditOwnHourlyRates = "edit_own_hourly_rates"
	PermEditHourlyRates = "edit_hourly_rates"
	PermViewCostRates = "view_cost_rates"
	PermBookUnitCostsForOneself = "book_unit_costs_for_oneself"
	PermBookUnitCosts = "book_unit_costs"
	PermEditOwnBookedUnitCosts = "edit_own_booked_unit_costs"
	PermEditBookedUnitCosts = "edit_booked_unit_costs"
	PermViewBookedCosts = "view_booked_costs"
	PermViewOwnBookedCosts = "view_own_booked_costs"
	PermSavePublicCostReports = "save_public_cost_reports"
	PermSavePrivateCostReports = "save_private_cost_reports"

	// ModuleWiki
	PermViewWiki = "view_wiki"
	PermViewWikiHistory = "view_wiki_history"
	PermEditWikiPages = "edit_wiki_pages"
	PermManageWiki = "manage_wiki"
)

// rawPermissionDefinitions contains baseline OpenProject permissions.
var rawPermissionDefinitions = []Permission{
	{
		ID:          PermAddProject,
		DisplayName: "Add project",
		Description: "Allows users to create new projects globally across the platform.",
		IsGlobal:    true,
		Module:      ModuleSystem,
		Context:     ContextGlobal,
		Operation:   OpCreate,
	},
	{
		ID:          PermCreateUser,
		DisplayName: "Create user",
		Description: "Allows administrators to provision new user accounts globally.",
		IsGlobal:    true,
		Module:      ModuleSystem,
		Context:     ContextGlobal,
		Operation:   OpCreate,
		ImpliedPerms: []string{PermViewAllUsersAndGroups},
	},
	{
		ID:          PermEditUsers,
		DisplayName: "Edit users",
		Description: "Allows administrators to edit user accounts, profiles, and authentication settings.",
		IsGlobal:    true,
		Module:      ModuleSystem,
		Context:     ContextGlobal,
		Operation:   OpUpdate,
		ImpliedPerms: []string{PermViewAllUsersAndGroups},
	},
	{
		ID:          PermDeleteUser,
		DisplayName: "Delete user",
		Description: "Allows administrators to lock or delete user accounts.",
		IsGlobal:    true,
		Module:      ModuleSystem,
		Context:     ContextGlobal,
		Operation:   OpDelete,
		ImpliedPerms: []string{PermViewAllUsersAndGroups},
	},
	{
		ID:          PermCreateBackup,
		DisplayName: "Create system backup",
		Description: "Allows administrators to trigger and download full system backups.",
		IsGlobal:    true,
		Module:      ModuleSystem,
		Context:     ContextGlobal,
		Operation:   OpAdmin,
	},
	{
		ID:          PermRestoreBackup,
		DisplayName: "Restore system backup",
		Description: "Allows administrators to restore system databases and assets from backups.",
		IsGlobal:    true,
		Module:      ModuleSystem,
		Context:     ContextGlobal,
		Operation:   OpAdmin,
		ImpliedPerms: []string{PermCreateBackup},
	},
	{
		ID:          PermViewAllUsersAndGroups,
		DisplayName: "View all users and groups",
		Description: "Allows viewing global directory of all registered users and groups.",
		IsGlobal:    true,
		Module:      ModuleSystem,
		Context:     ContextGlobal,
		Operation:   OpRead,
	},
	{
		ID:          PermManagePublicProjectLists,
		DisplayName: "Manage public project lists",
		Description: "Allows managing and publishing system-wide public project query views.",
		IsGlobal:    true,
		Module:      ModuleSystem,
		Context:     ContextGlobal,
		Operation:   OpManage,
	},
	{
		ID:          PermManageGlobalCustomFields,
		DisplayName: "Manage global custom fields",
		Description: "Allows creating and configuring custom fields applied across the system.",
		IsGlobal:    true,
		Module:      ModuleSystem,
		Context:     ContextGlobal,
		Operation:   OpAdmin,
	},
	{
		ID:          PermManageRoles,
		DisplayName: "Manage roles and permissions",
		Description: "Allows creating, configuring, and modifying role permission matrices.",
		IsGlobal:    true,
		Module:      ModuleSystem,
		Context:     ContextGlobal,
		Operation:   OpAdmin,
	},
	{
		ID:          PermArchiveProject,
		DisplayName: "Archive project",
		Description: "Allows users to archive and restore a project.",
		IsGlobal:    false,
		Module:      ModuleProject,
		Context:     ContextProject,
		Operation:   OpManage,
		ImpliedPerms: []string{PermEditProject},
	},
	{
		ID:          PermEditProject,
		DisplayName: "Edit project",
		Description: "Allows users to access Project settings and edit the project's configuration.",
		IsGlobal:    false,
		Module:      ModuleProject,
		Context:     ContextProject,
		Operation:   OpUpdate,
		ImpliedPerms: []string{PermViewProjectAttributes},
	},
	{
		ID:          PermSelectProjectModules,
		DisplayName: "Select project modules",
		Description: "Allows users to enable or disable project modules.",
		IsGlobal:    false,
		Module:      ModuleProject,
		Context:     ContextProject,
		Operation:   OpUpdate,
		ImpliedPerms: []string{PermEditProject},
	},
	{
		ID:          PermViewProjectAttributes,
		DisplayName: "View project attributes",
		Description: "Allows users to view project information and attributes.",
		IsGlobal:    false,
		Module:      ModuleProject,
		Context:     ContextProject,
		Operation:   OpRead,
	},
	{
		ID:          PermExportProjects,
		DisplayName: "Export projects",
		Description: "Allows users to export project information.",
		IsGlobal:    false,
		Module:      ModuleProject,
		Context:     ContextProject,
		Operation:   OpRead,
	},
	{
		ID:          PermEditProjectAttributes,
		DisplayName: "Edit project attributes",
		Description: "Allows users to edit project attributes on the overview page.",
		IsGlobal:    false,
		Module:      ModuleProject,
		Context:     ContextProject,
		Operation:   OpUpdate,
		ImpliedPerms: []string{PermViewProjectAttributes},
	},
	{
		ID:          PermSelectProjectAttributes,
		DisplayName: "Select project attributes",
		Description: "Allows users to configure which project attributes are available.",
		IsGlobal:    false,
		Module:      ModuleProject,
		Context:     ContextProject,
		Operation:   OpUpdate,
		ImpliedPerms: []string{PermViewProjectAttributes},
	},
	{
		ID:          PermViewProjectPhases,
		DisplayName: "View project phases",
		Description: "Allows users to view project life cycle phases.",
		IsGlobal:    false,
		Module:      ModuleProject,
		Context:     ContextProject,
		Operation:   OpRead,
	},
	{
		ID:          PermEditProjectPhases,
		DisplayName: "Edit project phases",
		Description: "Allows users to edit project life cycle phases.",
		IsGlobal:    false,
		Module:      ModuleProject,
		Context:     ContextProject,
		Operation:   OpUpdate,
		ImpliedPerms: []string{PermViewProjectPhases},
	},
	{
		ID:          PermSelectProjectPhases,
		DisplayName: "Select project phases",
		Description: "Allows users to activate or deactivate project phases.",
		IsGlobal:    false,
		Module:      ModuleProject,
		Context:     ContextProject,
		Operation:   OpUpdate,
		ImpliedPerms: []string{PermViewProjectPhases},
	},
	{
		ID:          PermManageMembers,
		DisplayName: "Manage members",
		Description: "Allows users to add, remove and manage project members and their roles.",
		IsGlobal:    false,
		Module:      ModuleProject,
		Context:     ContextProject,
		Operation:   OpManage,
		ImpliedPerms: []string{PermViewMembers},
	},
	{
		ID:          PermInviteMembersByEmail,
		DisplayName: "Invite members by email",
		Description: "Allows users to invite project members by email.",
		IsGlobal:    false,
		Module:      ModuleProject,
		Context:     ContextProject,
		Operation:   OpShare,
		ImpliedPerms: []string{PermManageMembers, PermViewMembers},
	},
	{
		ID:          PermViewMembers,
		DisplayName: "View members",
		Description: "Allows users to view project members.",
		IsGlobal:    false,
		Module:      ModuleProject,
		Context:     ContextProject,
		Operation:   OpRead,
	},
	{
		ID:          PermManageVersions,
		DisplayName: "Manage versions",
		Description: "Allows users to create, edit and delete project versions.",
		IsGlobal:    false,
		Module:      ModuleProject,
		Context:     ContextProject,
		Operation:   OpManage,
	},
	{
		ID:          PermSelectTypes,
		DisplayName: "Select types",
		Description: "Allows users to configure the work package types available in the project.",
		IsGlobal:    false,
		Module:      ModuleProject,
		Context:     ContextProject,
		Operation:   OpUpdate,
		ImpliedPerms: []string{PermEditProject},
	},
	{
		ID:          PermSelectCustomFields,
		DisplayName: "Select custom fields",
		Description: "Allows users to configure which custom fields are available in the project.",
		IsGlobal:    false,
		Module:      ModuleProject,
		Context:     ContextProject,
		Operation:   OpUpdate,
		ImpliedPerms: []string{PermEditProject},
	},
	{
		ID:          PermCreateSubprojects,
		DisplayName: "Create subprojects",
		Description: "Allows users to create subprojects.",
		IsGlobal:    false,
		Module:      ModuleProject,
		Context:     ContextProject,
		Operation:   OpCreate,
		ImpliedPerms: []string{PermViewProjectAttributes},
	},
	{
		ID:          PermCopyProjects,
		DisplayName: "Copy projects",
		Description: "Allows users to create a new project by copying an existing project.",
		IsGlobal:    false,
		Module:      ModuleProject,
		Context:     ContextProject,
		Operation:   OpSpecial,
		ImpliedPerms: []string{PermEditProject},
	},
	{
		ID:          PermManageDashboards,
		DisplayName: "Manage dashboards",
		Description: "Allows users to create and edit project dashboards.",
		IsGlobal:    false,
		Module:      ModuleProject,
		Context:     ContextProject,
		Operation:   OpManage,
		ImpliedPerms: []string{PermViewProjectAttributes},
	},
	{
		ID:          PermManageFilesInProject,
		DisplayName: "Manage files in project",
		Description: "Allows users to manage project file storages.",
		IsGlobal:    false,
		Module:      ModuleProject,
		Context:     ContextProject,
		Operation:   OpManage,
		ImpliedPerms: []string{PermAutoManagedFoldersReadFiles},
	},
	{
		ID:          PermAutoManagedFoldersReadFiles,
		DisplayName: "Automatically managed project folders: Read files",
		Description: "Allows users to read files in automatically managed project folders.",
		IsGlobal:    false,
		Module:      ModuleProject,
		Context:     ContextProject,
		Operation:   OpSpecial,
	},
	{
		ID:          PermAutoManagedFoldersWriteFiles,
		DisplayName: "Automatically managed project folders: Write files",
		Description: "Allows users to modify files in automatically managed project folders.",
		IsGlobal:    false,
		Module:      ModuleProject,
		Context:     ContextProject,
		Operation:   OpUpdate,
		ImpliedPerms: []string{PermAutoManagedFoldersReadFiles},
	},
	{
		ID:          PermAutoManagedFoldersCreateFiles,
		DisplayName: "Automatically managed project folders: Create files",
		Description: "Allows users to create files in automatically managed project folders.",
		IsGlobal:    false,
		Module:      ModuleProject,
		Context:     ContextProject,
		Operation:   OpCreate,
		ImpliedPerms: []string{PermAutoManagedFoldersReadFiles},
	},
	{
		ID:          PermAutoManagedFoldersDeleteFiles,
		DisplayName: "Automatically managed project folders: Delete files",
		Description: "Allows users to delete files from automatically managed project folders.",
		IsGlobal:    false,
		Module:      ModuleProject,
		Context:     ContextProject,
		Operation:   OpDelete,
		ImpliedPerms: []string{PermAutoManagedFoldersReadFiles},
	},
	{
		ID:          PermAutoManagedFoldersShareFiles,
		DisplayName: "Automatically managed project folders: Share files",
		Description: "Allows users to share files from automatically managed project folders.",
		IsGlobal:    false,
		Module:      ModuleProject,
		Context:     ContextProject,
		Operation:   OpShare,
		ImpliedPerms: []string{PermAutoManagedFoldersReadFiles},
	},
	{
		ID:          PermViewWorkPackages,
		DisplayName: "View work packages",
		Description: "Allows users to view work packages.",
		IsGlobal:    false,
		Module:      ModuleWorkPackages,
		Context:     ContextProject,
		PermissibleOn: []Context{ContextProject, ContextWorkPackage},
		Operation:   OpRead,
	},
	{
		ID:          PermAddWorkPackages,
		DisplayName: "Add work packages",
		Description: "Allows users to create work packages.",
		IsGlobal:    false,
		Module:      ModuleWorkPackages,
		Context:     ContextProject,
		Operation:   OpCreate,
		ImpliedPerms: []string{PermViewWorkPackages},
	},
	{
		ID:          PermEditWorkPackages,
		DisplayName: "Edit work packages",
		Description: "Allows users to edit work packages.",
		IsGlobal:    false,
		Module:      ModuleWorkPackages,
		Context:     ContextProject,
		PermissibleOn: []Context{ContextProject, ContextWorkPackage},
		Operation:   OpUpdate,
		ImpliedPerms: []string{PermViewWorkPackages},
	},
	{
		ID:          PermMoveWorkPackages,
		DisplayName: "Move work packages",
		Description: "Allows users to move work packages between projects.",
		IsGlobal:    false,
		Module:      ModuleWorkPackages,
		Context:     ContextProject,
		Operation:   OpSpecial,
		ImpliedPerms: []string{PermViewWorkPackages},
	},
	{
		ID:          PermDuplicateWorkPackages,
		DisplayName: "Duplicate work packages",
		Description: "Allows users to duplicate work packages.",
		IsGlobal:    false,
		Module:      ModuleWorkPackages,
		Context:     ContextProject,
		Operation:   OpSpecial,
		ImpliedPerms: []string{PermViewWorkPackages},
	},
	{
		ID:          PermAddComments,
		DisplayName: "Add comments",
		Description: "Allows users to comment on work packages.",
		IsGlobal:    false,
		Module:      ModuleWorkPackages,
		Context:     ContextProject,
		PermissibleOn: []Context{ContextProject, ContextWorkPackage},
		Operation:   OpCreate,
		ImpliedPerms: []string{PermViewWorkPackages},
	},
	{
		ID:          PermEditOwnComments,
		DisplayName: "Edit own comments",
		Description: "Allows users to edit their own comments.",
		IsGlobal:    false,
		Module:      ModuleWorkPackages,
		Context:     ContextProject,
		PermissibleOn: []Context{ContextProject, ContextWorkPackage},
		Operation:   OpUpdate,
		ImpliedPerms: []string{PermViewWorkPackages, PermAddComments},
	},
	{
		ID:          PermModerateComments,
		DisplayName: "Moderate comments",
		Description: "Allows users to edit comments created by any user.",
		IsGlobal:    false,
		Module:      ModuleWorkPackages,
		Context:     ContextProject,
		Operation:   OpManage,
		ImpliedPerms: []string{PermViewWorkPackages, PermAddComments},
	},
	{
		ID:          PermViewInternalComments,
		DisplayName: "View internal comments",
		Description: "Allows users to view internal comments.",
		IsGlobal:    false,
		Module:      ModuleWorkPackages,
		Context:     ContextProject,
		Operation:   OpRead,
		ImpliedPerms: []string{PermViewWorkPackages},
	},
	{
		ID:          PermWriteInternalComments,
		DisplayName: "Write internal comments",
		Description: "Allows users to create internal comments.",
		IsGlobal:    false,
		Module:      ModuleWorkPackages,
		Context:     ContextProject,
		Operation:   OpSpecial,
		ImpliedPerms: []string{PermViewInternalComments, PermViewWorkPackages},
	},
	{
		ID:          PermEditOwnInternalComments,
		DisplayName: "Edit own internal comments",
		Description: "Allows users to edit their own internal comments.",
		IsGlobal:    false,
		Module:      ModuleWorkPackages,
		Context:     ContextProject,
		Operation:   OpUpdate,
		ImpliedPerms: []string{PermWriteInternalComments, PermViewInternalComments},
	},
	{
		ID:          PermModerateInternalComments,
		DisplayName: "Moderate internal comments",
		Description: "Allows users to edit internal comments created by any user.",
		IsGlobal:    false,
		Module:      ModuleWorkPackages,
		Context:     ContextProject,
		Operation:   OpManage,
		ImpliedPerms: []string{PermViewInternalComments},
	},
	{
		ID:          PermAddAttachments,
		DisplayName: "Add attachments",
		Description: "Allows users to upload attachments to work packages.",
		IsGlobal:    false,
		Module:      ModuleWorkPackages,
		Context:     ContextProject,
		PermissibleOn: []Context{ContextProject, ContextWorkPackage},
		Operation:   OpCreate,
		ImpliedPerms: []string{PermViewWorkPackages},
	},
	{
		ID:          PermManageWorkPackageCategories,
		DisplayName: "Manage work package categories",
		Description: "Allows users to create, edit, and delete work package categories.",
		IsGlobal:    false,
		Module:      ModuleWorkPackages,
		Context:     ContextProject,
		Operation:   OpManage,
		ImpliedPerms: []string{PermViewWorkPackages},
	},
	{
		ID:          PermExportWorkPackages,
		DisplayName: "Export work packages",
		Description: "Allows users to export work packages.",
		IsGlobal:    false,
		Module:      ModuleWorkPackages,
		Context:     ContextProject,
		Operation:   OpRead,
		ImpliedPerms: []string{PermViewWorkPackages},
	},
	{
		ID:          PermDeleteWorkPackages,
		DisplayName: "Delete work packages",
		Description: "Allows users to delete work packages.",
		IsGlobal:    false,
		Module:      ModuleWorkPackages,
		Context:     ContextProject,
		Operation:   OpDelete,
		ImpliedPerms: []string{PermViewWorkPackages},
	},
	{
		ID:          PermManageWorkPackageRelations,
		DisplayName: "Manage work package relations",
		Description: "Allows users to create, edit, and remove work package relations.",
		IsGlobal:    false,
		Module:      ModuleWorkPackages,
		Context:     ContextProject,
		Operation:   OpManage,
		ImpliedPerms: []string{PermViewWorkPackages},
	},
	{
		ID:          PermManageWorkPackageHierarchies,
		DisplayName: "Manage work package hierarchies",
		Description: "Allows users to manage parent-child relationships between work packages.",
		IsGlobal:    false,
		Module:      ModuleWorkPackages,
		Context:     ContextProject,
		Operation:   OpManage,
		ImpliedPerms: []string{PermViewWorkPackages},
	},
	{
		ID:          PermManagePublicViews,
		DisplayName: "Manage public views",
		Description: "Allows users to create, edit, and delete public work package views.",
		IsGlobal:    false,
		Module:      ModuleWorkPackages,
		Context:     ContextProject,
		PermissibleOn: []Context{ContextProject, ContextProjectQuery},
		Operation:   OpManage,
		ImpliedPerms: []string{PermSaveViews, PermViewWorkPackages},
	},
	{
		ID:          PermSaveViews,
		DisplayName: "Save views",
		Description: "Allows users to save personal work package views.",
		IsGlobal:    false,
		Module:      ModuleWorkPackages,
		Context:     ContextProject,
		PermissibleOn: []Context{ContextProject, ContextProjectQuery},
		Operation:   OpCreate,
		ImpliedPerms: []string{PermViewWorkPackages},
	},
	{
		ID:          PermViewWatchersList,
		DisplayName: "View watchers list",
		Description: "Allows users to see who is watching a work package.",
		IsGlobal:    false,
		Module:      ModuleWorkPackages,
		Context:     ContextProject,
		Operation:   OpRead,
		ImpliedPerms: []string{PermViewWorkPackages},
	},
	{
		ID:          PermAddWatchers,
		DisplayName: "Add watchers",
		Description: "Allows users to add watchers to work packages.",
		IsGlobal:    false,
		Module:      ModuleWorkPackages,
		Context:     ContextProject,
		Operation:   OpCreate,
		ImpliedPerms: []string{PermViewWatchersList, PermViewWorkPackages},
	},
	{
		ID:          PermDeleteWatchers,
		DisplayName: "Delete watchers",
		Description: "Allows users to remove watchers from work packages.",
		IsGlobal:    false,
		Module:      ModuleWorkPackages,
		Context:     ContextProject,
		Operation:   OpDelete,
		ImpliedPerms: []string{PermViewWatchersList, PermViewWorkPackages},
	},
	{
		ID:          PermShareWorkPackages,
		DisplayName: "Share work packages",
		Description: "Allows users to share work packages with other users.",
		IsGlobal:    false,
		Module:      ModuleWorkPackages,
		Context:     ContextProject,
		PermissibleOn: []Context{ContextProject, ContextWorkPackage},
		Operation:   OpShare,
		ImpliedPerms: []string{PermViewWorkPackages},
	},
	{
		ID:          PermViewWorkPackageShares,
		DisplayName: "View work package shares",
		Description: "Allows users to view existing work package shares.",
		IsGlobal:    false,
		Module:      ModuleWorkPackages,
		Context:     ContextProject,
		PermissibleOn: []Context{ContextProject, ContextWorkPackage},
		Operation:   OpRead,
		ImpliedPerms: []string{PermViewWorkPackages},
	},
	{
		ID:          PermAssignVersions,
		DisplayName: "Assign versions",
		Description: "Allows users to assign versions to work packages.",
		IsGlobal:    false,
		Module:      ModuleWorkPackages,
		Context:     ContextProject,
		Operation:   OpUpdate,
		ImpliedPerms: []string{PermViewWorkPackages},
	},
	{
		ID:          PermChangeWorkPackageStatus,
		DisplayName: "Change work package status",
		Description: "Allows users to change the status of work packages.",
		IsGlobal:    false,
		Module:      ModuleWorkPackages,
		Context:     ContextProject,
		PermissibleOn: []Context{ContextProject, ContextWorkPackage},
		Operation:   OpUpdate,
		ImpliedPerms: []string{PermViewWorkPackages},
	},
	{
		ID:          PermBecomeAssigneeResponsible,
		DisplayName: "Become assignee/responsible",
		Description: "Allows work packages to be assigned to users or groups with this role in the project.",
		IsGlobal:    false,
		Module:      ModuleWorkPackages,
		Context:     ContextProject,
		Operation:   OpSpecial,
		ImpliedPerms: []string{PermViewWorkPackages},
	},
	{
		ID:          PermViewFileLinks,
		DisplayName: "View file links",
		Description: "Allows users to view file links attached to work packages.",
		IsGlobal:    false,
		Module:      ModuleWorkPackages,
		Context:     ContextProject,
		Operation:   OpRead,
		ImpliedPerms: []string{PermViewWorkPackages},
	},
	{
		ID:          PermManageFileLinks,
		DisplayName: "Manage file links",
		Description: "Allows users to create, edit, and remove file links.",
		IsGlobal:    false,
		Module:      ModuleWorkPackages,
		Context:     ContextProject,
		Operation:   OpManage,
		ImpliedPerms: []string{PermViewFileLinks, PermViewWorkPackages},
	},
	{
		ID:          PermManageWikiPageLinks,
		DisplayName: "Manage wiki page links",
		Description: "Allows users to create, edit, and remove wiki page links.",
		IsGlobal:    false,
		Module:      ModuleWorkPackages,
		Context:     ContextProject,
		Operation:   OpManage,
		ImpliedPerms: []string{PermViewWorkPackages},
	},
	{
		ID:          PermViewBoards,
		DisplayName: "View boards",
		Description: "Allows users to view boards.",
		IsGlobal:    false,
		Module:      ModuleBoards,
		Context:     ContextProject,
		Operation:   OpRead,
	},
	{
		ID:          PermManageBoards,
		DisplayName: "Manage boards",
		Description: "Allows users to create, edit, and delete boards.",
		IsGlobal:    false,
		Module:      ModuleBoards,
		Context:     ContextProject,
		Operation:   OpManage,
		ImpliedPerms: []string{PermViewBoards},
	},
	{
		ID:          PermViewSprints,
		DisplayName: "View sprints",
		Description: "Allows users to view sprints.",
		IsGlobal:    false,
		Module:      ModuleBacklogs,
		Context:     ContextProject,
		Operation:   OpRead,
	},
	{
		ID:          PermSelectBacklogTypesAndStatuses,
		DisplayName: "Select backlog types and statuses",
		Description: "Allows users to configure which work package types appear in the backlog and which statuses are considered completed.",
		IsGlobal:    false,
		Module:      ModuleBacklogs,
		Context:     ContextProject,
		Operation:   OpUpdate,
		ImpliedPerms: []string{PermViewSprints},
	},
	{
		ID:          PermCreateSprints,
		DisplayName: "Create sprints",
		Description: "Allows users to create sprints.",
		IsGlobal:    false,
		Module:      ModuleBacklogs,
		Context:     ContextProject,
		Operation:   OpCreate,
		ImpliedPerms: []string{PermViewSprints},
	},
	{
		ID:          PermStartCompleteSprint,
		DisplayName: "Start/complete sprint",
		Description: "Allows users to start and complete sprints.",
		IsGlobal:    false,
		Module:      ModuleBacklogs,
		Context:     ContextProject,
		Operation:   OpUpdate,
		ImpliedPerms: []string{PermViewSprints},
	},
	{
		ID:          PermManageSprintItems,
		DisplayName: "Manage sprint items",
		Description: "Allows users to manage sprint contents.",
		IsGlobal:    false,
		Module:      ModuleBacklogs,
		Context:     ContextProject,
		Operation:   OpManage,
		ImpliedPerms: []string{PermViewSprints},
	},
	{
		ID:          PermShareSprint,
		DisplayName: "Share sprint",
		Description: "Allows users to share sprint information.",
		IsGlobal:    false,
		Module:      ModuleBacklogs,
		Context:     ContextProject,
		Operation:   OpShare,
		ImpliedPerms: []string{PermViewSprints},
	},
	{
		ID:          PermViewBudgets,
		DisplayName: "View budgets",
		Description: "Allows users to view project budgets.",
		IsGlobal:    false,
		Module:      ModuleBudgets,
		Context:     ContextProject,
		Operation:   OpRead,
	},
	{
		ID:          PermEditBudgets,
		DisplayName: "Edit budgets",
		Description: "Allows users to create, edit, and delete project budgets.",
		IsGlobal:    false,
		Module:      ModuleBudgets,
		Context:     ContextProject,
		Operation:   OpUpdate,
		ImpliedPerms: []string{PermViewBudgets},
	},
	{
		ID:          PermViewCalendars,
		DisplayName: "View calendars",
		Description: "Allows users to view calendars.",
		IsGlobal:    false,
		Module:      ModuleCalendars,
		Context:     ContextProject,
		Operation:   OpRead,
	},
	{
		ID:          PermEditCalendars,
		DisplayName: "Edit calendars",
		Description: "Allows users to create, edit, and delete calendars.",
		IsGlobal:    false,
		Module:      ModuleCalendars,
		Context:     ContextProject,
		Operation:   OpUpdate,
		ImpliedPerms: []string{PermViewCalendars},
	},
	{
		ID:          PermSubscribeToIcalendars,
		DisplayName: "Subscribe to iCalendars",
		Description: "Allows users to subscribe to calendar feeds.",
		IsGlobal:    false,
		Module:      ModuleCalendars,
		Context:     ContextProject,
		Operation:   OpRead,
		ImpliedPerms: []string{PermViewCalendars},
	},
	{
		ID:          PermViewDocuments,
		DisplayName: "View documents",
		Description: "Allows users to view project documents.",
		IsGlobal:    false,
		Module:      ModuleDocuments,
		Context:     ContextProject,
		Operation:   OpRead,
	},
	{
		ID:          PermManageDocuments,
		DisplayName: "Manage documents",
		Description: "Allows users to create, edit, and delete project documents.",
		IsGlobal:    false,
		Module:      ModuleDocuments,
		Context:     ContextProject,
		Operation:   OpManage,
		ImpliedPerms: []string{PermViewDocuments},
	},
	{
		ID:          PermManageForums,
		DisplayName: "Manage forums",
		Description: "Allows users to create, edit, and delete forums.",
		IsGlobal:    false,
		Module:      ModuleForums,
		Context:     ContextProject,
		Operation:   OpManage,
		ImpliedPerms: []string{PermPostMessages},
	},
	{
		ID:          PermPostMessages,
		DisplayName: "Post messages",
		Description: "Allows users to post messages in forums.",
		IsGlobal:    false,
		Module:      ModuleForums,
		Context:     ContextProject,
		Operation:   OpCreate,
		ImpliedPerms: []string{PermEditOwnMessages},
	},
	{
		ID:          PermEditMessages,
		DisplayName: "Edit messages",
		Description: "Allows users to edit any forum message.",
		IsGlobal:    false,
		Module:      ModuleForums,
		Context:     ContextProject,
		Operation:   OpUpdate,
		ImpliedPerms: []string{PermEditOwnMessages},
	},
	{
		ID:          PermEditOwnMessages,
		DisplayName: "Edit own messages",
		Description: "Allows users to edit their own forum messages.",
		IsGlobal:    false,
		Module:      ModuleForums,
		Context:     ContextProject,
		Operation:   OpUpdate,
	},
	{
		ID:          PermDeleteMessages,
		DisplayName: "Delete messages",
		Description: "Allows users to delete any forum message.",
		IsGlobal:    false,
		Module:      ModuleForums,
		Context:     ContextProject,
		Operation:   OpDelete,
		ImpliedPerms: []string{PermDeleteOwnMessages},
	},
	{
		ID:          PermDeleteOwnMessages,
		DisplayName: "Delete own messages",
		Description: "Allows users to delete their own forum messages.",
		IsGlobal:    false,
		Module:      ModuleForums,
		Context:     ContextProject,
		Operation:   OpDelete,
	},
	{
		ID:          PermShowGithubContent,
		DisplayName: "Show GitHub content",
		Description: "Allows users to see GitHub pull requests and issues linked to work packages.",
		IsGlobal:    false,
		Module:      ModuleGitHub,
		Context:     ContextProject,
		Operation:   OpRead,
	},
	{
		ID:          PermShowGitlabContent,
		DisplayName: "Show GitLab content",
		Description: "Allows users to see GitLab merge requests and issues linked to work packages.",
		IsGlobal:    false,
		Module:      ModuleGitLab,
		Context:     ContextProject,
		Operation:   OpRead,
	},
	{
		ID:          PermViewMeetings,
		DisplayName: "View meetings",
		Description: "Allows users to view meetings in a project.",
		IsGlobal:    false,
		Module:      ModuleMeetings,
		Context:     ContextProject,
		Operation:   OpRead,
	},
	{
		ID:          PermCreateMeetings,
		DisplayName: "Create meetings",
		Description: "Allows users to create new meetings.",
		IsGlobal:    false,
		Module:      ModuleMeetings,
		Context:     ContextProject,
		Operation:   OpCreate,
		ImpliedPerms: []string{PermViewMeetings},
	},
	{
		ID:          PermEditMeetings,
		DisplayName: "Edit meetings",
		Description: "Allows users to edit existing meetings.",
		IsGlobal:    false,
		Module:      ModuleMeetings,
		Context:     ContextProject,
		Operation:   OpUpdate,
		ImpliedPerms: []string{PermViewMeetings},
	},
	{
		ID:          PermDeleteMeetings,
		DisplayName: "Delete meetings",
		Description: "Allows users to delete meetings.",
		IsGlobal:    false,
		Module:      ModuleMeetings,
		Context:     ContextProject,
		Operation:   OpDelete,
		ImpliedPerms: []string{PermViewMeetings},
	},
	{
		ID:          PermSendMeetingInvites,
		DisplayName: "Send meeting invites",
		Description: "Allows users to send email invitations for meetings to participants.",
		IsGlobal:    false,
		Module:      ModuleMeetings,
		Context:     ContextProject,
		Operation:   OpShare,
		ImpliedPerms: []string{PermViewMeetings},
	},
	{
		ID:          PermManageAgendas,
		DisplayName: "Manage agendas",
		Description: "Allows users to create, edit, and close meeting agendas.",
		IsGlobal:    false,
		Module:      ModuleMeetings,
		Context:     ContextProject,
		Operation:   OpManage,
		ImpliedPerms: []string{PermViewMeetings},
	},
	{
		ID:          PermManageOutcomes,
		DisplayName: "Manage outcomes",
		Description: "Allows users to create, edit, and close meeting outcomes/minutes.",
		IsGlobal:    false,
		Module:      ModuleMeetings,
		Context:     ContextProject,
		Operation:   OpManage,
		ImpliedPerms: []string{PermViewMeetings},
	},
	{
		ID:          PermManageNews,
		DisplayName: "Manage news",
		Description: "Allows users to create, edit, and delete news entries.",
		IsGlobal:    false,
		Module:      ModuleNews,
		Context:     ContextProject,
		Operation:   OpManage,
		ImpliedPerms: []string{PermCommentNews},
	},
	{
		ID:          PermCommentNews,
		DisplayName: "Comment news",
		Description: "Allows users to post comments on news entries.",
		IsGlobal:    false,
		Module:      ModuleNews,
		Context:     ContextProject,
		Operation:   OpSpecial,
	},
	{
		ID:          PermViewTeamPlanner,
		DisplayName: "View team planner",
		Description: "Allows users to view team planner views.",
		IsGlobal:    false,
		Module:      ModuleTeamPlanner,
		Context:     ContextProject,
		Operation:   OpRead,
	},
	{
		ID:          PermManageTeamPlanner,
		DisplayName: "Manage team planner",
		Description: "Allows users to create, edit, and delete team planner views.",
		IsGlobal:    false,
		Module:      ModuleTeamPlanner,
		Context:     ContextProject,
		Operation:   OpManage,
		ImpliedPerms: []string{PermViewTeamPlanner},
	},
	{
		ID:          PermViewSpentTime,
		DisplayName: "View spent time",
		Description: "Allows users to view time entries logged by all users.",
		IsGlobal:    false,
		Module:      ModuleTimeCosts,
		Context:     ContextProject,
		Operation:   OpRead,
		ImpliedPerms: []string{PermViewOwnSpentTime},
	},
	{
		ID:          PermViewOwnSpentTime,
		DisplayName: "View own spent time",
		Description: "Allows users to view only their own time entries.",
		IsGlobal:    false,
		Module:      ModuleTimeCosts,
		Context:     ContextProject,
		Operation:   OpRead,
	},
	{
		ID:          PermLogOwnTime,
		DisplayName: "Log own time",
		Description: "Allows users to log time on work packages for themselves.",
		IsGlobal:    false,
		Module:      ModuleTimeCosts,
		Context:     ContextProject,
		Operation:   OpCreate,
		ImpliedPerms: []string{PermViewOwnSpentTime},
	},
	{
		ID:          PermLogTimeForOtherUsers,
		DisplayName: "Log time for other users",
		Description: "Allows users to log time entries on behalf of other users.",
		IsGlobal:    false,
		Module:      ModuleTimeCosts,
		Context:     ContextProject,
		Operation:   OpCreate,
		ImpliedPerms: []string{PermLogOwnTime},
	},
	{
		ID:          PermEditOwnTimeLogs,
		DisplayName: "Edit own time logs",
		Description: "Allows users to edit their own logged time entries.",
		IsGlobal:    false,
		Module:      ModuleTimeCosts,
		Context:     ContextProject,
		Operation:   OpUpdate,
		ImpliedPerms: []string{PermViewOwnSpentTime},
	},
	{
		ID:          PermEditTimeLogsForOtherUsers,
		DisplayName: "Edit time logs for other users",
		Description: "Allows users to edit time entries logged by other users.",
		IsGlobal:    false,
		Module:      ModuleTimeCosts,
		Context:     ContextProject,
		Operation:   OpUpdate,
		ImpliedPerms: []string{PermEditOwnTimeLogs, PermViewSpentTime},
	},
	{
		ID:          PermManageProjectActivities,
		DisplayName: "Manage project activities",
		Description: "Allows users to manage time tracking activity types within a project.",
		IsGlobal:    false,
		Module:      ModuleTimeCosts,
		Context:     ContextProject,
		Operation:   OpManage,
	},
	{
		ID:          PermViewOwnHourlyRate,
		DisplayName: "View own hourly rate",
		Description: "Allows users to view their own hourly labor rate.",
		IsGlobal:    false,
		Module:      ModuleTimeCosts,
		Context:     ContextProject,
		Operation:   OpRead,
	},
	{
		ID:          PermViewAllHourlyRates,
		DisplayName: "View all hourly rates",
		Description: "Allows users to view hourly rates of all members.",
		IsGlobal:    false,
		Module:      ModuleTimeCosts,
		Context:     ContextProject,
		Operation:   OpRead,
		ImpliedPerms: []string{PermViewOwnHourlyRate},
	},
	{
		ID:          PermEditOwnHourlyRates,
		DisplayName: "Edit own hourly rates",
		Description: "Allows users to edit their own hourly labor rate.",
		IsGlobal:    false,
		Module:      ModuleTimeCosts,
		Context:     ContextProject,
		Operation:   OpUpdate,
		ImpliedPerms: []string{PermViewOwnHourlyRate},
	},
	{
		ID:          PermEditHourlyRates,
		DisplayName: "Edit hourly rates",
		Description: "Allows users to edit hourly rates for all members.",
		IsGlobal:    false,
		Module:      ModuleTimeCosts,
		Context:     ContextProject,
		Operation:   OpUpdate,
		ImpliedPerms: []string{PermViewAllHourlyRates, PermEditOwnHourlyRates},
	},
	{
		ID:          PermViewCostRates,
		DisplayName: "View cost rates",
		Description: "Allows users to view cost types and their rates.",
		IsGlobal:    false,
		Module:      ModuleTimeCosts,
		Context:     ContextProject,
		Operation:   OpRead,
	},
	{
		ID:          PermBookUnitCostsForOneself,
		DisplayName: "Book unit costs for oneself",
		Description: "Allows users to book unit costs on work packages for themselves.",
		IsGlobal:    false,
		Module:      ModuleTimeCosts,
		Context:     ContextProject,
		Operation:   OpCreate,
		ImpliedPerms: []string{PermViewOwnBookedCosts},
	},
	{
		ID:          PermBookUnitCosts,
		DisplayName: "Book unit costs",
		Description: "Allows users to book unit costs on work packages for all members.",
		IsGlobal:    false,
		Module:      ModuleTimeCosts,
		Context:     ContextProject,
		Operation:   OpCreate,
		ImpliedPerms: []string{PermBookUnitCostsForOneself},
	},
	{
		ID:          PermEditOwnBookedUnitCosts,
		DisplayName: "Edit own booked unit costs",
		Description: "Allows users to edit their own booked unit cost entries.",
		IsGlobal:    false,
		Module:      ModuleTimeCosts,
		Context:     ContextProject,
		Operation:   OpUpdate,
		ImpliedPerms: []string{PermViewOwnBookedCosts},
	},
	{
		ID:          PermEditBookedUnitCosts,
		DisplayName: "Edit booked unit costs",
		Description: "Allows users to edit unit cost entries booked by any member.",
		IsGlobal:    false,
		Module:      ModuleTimeCosts,
		Context:     ContextProject,
		Operation:   OpUpdate,
		ImpliedPerms: []string{PermEditOwnBookedUnitCosts, PermViewBookedCosts},
	},
	{
		ID:          PermViewBookedCosts,
		DisplayName: "View booked costs",
		Description: "Allows users to view booked costs for all members.",
		IsGlobal:    false,
		Module:      ModuleTimeCosts,
		Context:     ContextProject,
		Operation:   OpRead,
		ImpliedPerms: []string{PermViewOwnBookedCosts},
	},
	{
		ID:          PermViewOwnBookedCosts,
		DisplayName: "View own booked costs",
		Description: "Allows users to view only their own booked costs.",
		IsGlobal:    false,
		Module:      ModuleTimeCosts,
		Context:     ContextProject,
		Operation:   OpRead,
	},
	{
		ID:          PermSavePublicCostReports,
		DisplayName: "Save public cost reports",
		Description: "Allows users to save cost reports visible to all project members.",
		IsGlobal:    false,
		Module:      ModuleTimeCosts,
		Context:     ContextProject,
		Operation:   OpCreate,
		ImpliedPerms: []string{PermSavePrivateCostReports},
	},
	{
		ID:          PermSavePrivateCostReports,
		DisplayName: "Save private cost reports",
		Description: "Allows users to save cost reports visible only to themselves.",
		IsGlobal:    false,
		Module:      ModuleTimeCosts,
		Context:     ContextProject,
		Operation:   OpCreate,
	},
	{
		ID:          PermViewWiki,
		DisplayName: "View wiki",
		Description: "Allows users to view wiki pages.",
		IsGlobal:    false,
		Module:      ModuleWiki,
		Context:     ContextProject,
		Operation:   OpRead,
	},
	{
		ID:          PermViewWikiHistory,
		DisplayName: "View wiki history",
		Description: "Allows users to view the edit history of wiki pages.",
		IsGlobal:    false,
		Module:      ModuleWiki,
		Context:     ContextProject,
		Operation:   OpRead,
		ImpliedPerms: []string{PermViewWiki},
	},
	{
		ID:          PermEditWikiPages,
		DisplayName: "Edit wiki pages",
		Description: "Allows users to create and edit wiki pages.",
		IsGlobal:    false,
		Module:      ModuleWiki,
		Context:     ContextProject,
		Operation:   OpUpdate,
		ImpliedPerms: []string{PermViewWiki},
	},
	{
		ID:          PermManageWiki,
		DisplayName: "Manage wiki",
		Description: "Allows users to manage the wiki, including renaming, deleting pages, and managing the wiki menu.",
		IsGlobal:    false,
		Module:      ModuleWiki,
		Context:     ContextProject,
		Operation:   OpManage,
		ImpliedPerms: []string{PermEditWikiPages, PermViewWiki},
	},
}

// BuildDefaultCatalog builds the indexed permission catalog and inverts implied
// permissions to compute dependent permissions.
// If Permission A implies Permission B, then Permission B has dependent Permission A.
func BuildDefaultCatalog() map[string]Permission {
	catalog := make(map[string]Permission, len(rawPermissionDefinitions))
	for _, p := range rawPermissionDefinitions {
		catalog[p.ID] = p
	}

	// Compute dependent permissions
	dependentsMap := make(map[string]map[string]struct{})
	for _, p := range rawPermissionDefinitions {
		for _, impliedID := range p.ImpliedPerms {
			if dependentsMap[impliedID] == nil {
				dependentsMap[impliedID] = make(map[string]struct{})
			}
			dependentsMap[impliedID][p.ID] = struct{}{}
		}
	}

	for id, deps := range dependentsMap {
		if p, ok := catalog[id]; ok {
			depList := make([]string, 0, len(deps))
			for depID := range deps {
				depList = append(depList, depID)
			}
			sort.Strings(depList)
			p.DependentPerms = depList
			catalog[id] = p
		}
	}

	return catalog
}
