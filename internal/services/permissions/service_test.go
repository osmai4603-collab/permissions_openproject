package perms

import (
	"slices"
	"testing"
)

func TestAllPermissions_NoDuplicates(t *testing.T) {
	all := AllPermissions()
	if len(all) == 0 {
		t.Fatal("expected permissions to be registered, got 0")
	}

	seen := make(map[string]bool)
	for _, p := range all {
		if p.ID == "" {
			t.Errorf("found permission with empty ID: %+v", p)
		}
		if seen[p.ID] {
			t.Errorf("duplicate permission ID found: %s", p.ID)
		}
		seen[p.ID] = true

		if len(p.PermissibleOn) == 0 {
			t.Errorf("permission %s has empty PermissibleOn", p.ID)
		}
	}
}

func TestGlobalPermissions(t *testing.T) {
	globals := GlobalPermissions()
	if len(globals) != 11 {
		t.Fatalf("expected 11 global permissions, got %d", len(globals))
	}

	for _, p := range globals {
		if p.Module != ModuleGlobal {
			t.Errorf("expected module %s for global permission %s, got %s", ModuleGlobal, p.ID, p.Module)
		}
		if !slices.Contains(p.PermissibleOn, ContextGlobal) {
			t.Errorf("permission %s should be permissible on ContextGlobal", p.ID)
		}
	}
}

func TestWorkPackages_MultiContext(t *testing.T) {
	expectedMulti := []string{
		"view_work_packages",
		"edit_work_packages",
		"add_comments",
		"add_attachments",
		"view_file_links",
		"manage_file_links",
	}

	for _, id := range expectedMulti {
		p, ok := PermissionByID(id)
		if !ok {
			t.Errorf("expected permission %s to exist", id)
			continue
		}
		if !p.AllowsContext(ContextProject) {
			t.Errorf("permission %s must allow ContextProject", id)
		}
		if !p.AllowsContext(ContextWorkPackage) {
			t.Errorf("permission %s must allow ContextWorkPackage", id)
		}
	}
}

func TestValidatePermissionsForContext(t *testing.T) {
	svc := NewService()

	// Global context test
	input := []string{"add_project", "view_work_packages", "create_user", "invalid_id"}
	validGlobal := svc.ValidatePermissionsForContext(input, ContextGlobal)
	expectedGlobal := []string{"add_project", "create_user"}
	if !slices.Equal(validGlobal, expectedGlobal) {
		t.Errorf("expected %v for ContextGlobal, got %v", expectedGlobal, validGlobal)
	}

	// WorkPackage context test
	inputWP := []string{"add_project", "view_work_packages", "edit_work_packages", "manage_versions"}
	validWP := svc.ValidatePermissionsForContext(inputWP, ContextWorkPackage)
	expectedWP := []string{"edit_work_packages", "view_work_packages"}
	if !slices.Equal(validWP, expectedWP) {
		t.Errorf("expected %v for ContextWorkPackage, got %v", expectedWP, validWP)
	}
}

func TestResolveRevocation(t *testing.T) {
	svc := NewService()

	// "invite_members_by_email" depends on "manage_members"
	active := []string{"manage_members", "invite_members_by_email", "view_work_packages"}
	remaining := svc.ResolveRevocation(active, "manage_members")

	// Revoking "manage_members" should also revoke "invite_members_by_email"
	expected := []string{"view_work_packages"}
	if !slices.Equal(remaining, expected) {
		t.Errorf("expected %v after revoking manage_members, got %v", expected, remaining)
	}
}

func TestPermissionsByContext(t *testing.T) {
	svc := NewService()

	globals := svc.GetPermissionsByContext(ContextGlobal)
	if len(globals) != 11 {
		t.Errorf("expected 11 permissions for ContextGlobal, got %d", len(globals))
	}

	workPackages := svc.GetPermissionsByContext(ContextWorkPackage)
	if len(workPackages) != 6 {
		t.Errorf("expected 6 permissions for ContextWorkPackage, got %d", len(workPackages))
	}
}

func TestCatalogCompleteness(t *testing.T) {
	catalogPerms := []string{
		PermAddProject, PermCreateUser, PermEditUsers, PermCreateBackup,
		PermViewAllUsersAndGroups, PermManagePublicProjectLists, PermCreatePortfolios,
		PermCreatePrograms, PermEditAttributeHelpTexts, PermManagePlaceholderUsers,
		PermViewAllUsersMailAddresses,
		PermArchiveProject, PermEditProject, PermSelectProjectModules,
		PermViewProjectAttributes, PermExportProjects, PermEditProjectAttributes,
		PermSelectProjectAttributes, PermViewProjectPhases, PermEditProjectPhases,
		PermSelectProjectPhases, PermManageMembers, PermInviteMembersByEmail,
		PermViewMembers, PermManageVersions, PermSelectTypes,
		PermSelectCustomFields, PermCreateSubprojects, PermCopyProjects,
		PermManageDashboards, PermManageFilesInProject, PermAutoManagedFoldersReadFiles,
		PermAutoManagedFoldersWriteFiles, PermAutoManagedFoldersCreateFiles,
		PermAutoManagedFoldersDeleteFiles, PermAutoManagedFoldersShareFiles,
		PermViewWorkPackages, PermAddWorkPackages, PermEditWorkPackages,
		PermMoveWorkPackages, PermDuplicateWorkPackages, PermAddComments,
		PermEditOwnComments, PermModerateComments, PermViewInternalComments,
		PermWriteInternalComments, PermEditOwnInternalComments, PermModerateInternalComments,
		PermAddAttachments, PermManageWorkPackageCategories, PermExportWorkPackages,
		PermDeleteWorkPackages, PermManageWorkPackageRelations, PermManageWorkPackageHierarchies,
		PermManagePublicViews, PermSaveViews, PermViewWatchersList,
		PermAddWatchers, PermDeleteWatchers, PermShareWorkPackages,
		PermViewWorkPackageShares, PermAssignVersions, PermChangeWorkPackageStatus,
		PermBecomeAssigneeResponsible, PermViewFileLinks, PermManageFileLinks,
		PermManageWikiPageLinks,
		PermViewBoards, PermManageBoards,
		PermViewSprints, PermSelectBacklogTypesAndStatuses, PermCreateSprints,
		PermStartCompleteSprint, PermManageSprintItems, PermShareSprint,
		PermViewBudgets, PermEditBudgets,
		PermViewCalendars, PermEditCalendars, PermSubscribeToIcalendars,
		PermViewDocuments, PermManageDocuments,
		PermManageForums, PermPostMessages, PermEditMessages,
		PermEditOwnMessages, PermDeleteMessages, PermDeleteOwnMessages,
		PermShowGithubContent,
		PermShowGitlabContent,
		PermViewMeetings, PermCreateMeetings, PermEditMeetings,
		PermDeleteMeetings, PermSendMeetingInvites, PermManageAgendas,
		PermManageOutcomes,
		PermManageNews, PermCommentNews,
		PermViewTeamPlanner, PermManageTeamPlanner,
		PermViewSpentTime, PermViewOwnSpentTime, PermLogOwnTime,
		PermLogTimeForOtherUsers, PermEditOwnTimeLogs, PermEditTimeLogsForOtherUsers,
		PermManageProjectActivities, PermViewOwnHourlyRate, PermViewAllHourlyRates,
		PermEditOwnHourlyRates, PermEditHourlyRates, PermViewCostRates,
		PermBookUnitCostsForOneself, PermBookUnitCosts, PermEditOwnBookedUnitCosts,
		PermEditBookedUnitCosts, PermViewBookedCosts, PermViewOwnBookedCosts,
		PermSavePublicCostReports, PermSavePrivateCostReports,
		PermViewWiki, PermViewWikiHistory, PermEditWikiPages, PermManageWiki,
	}

	all := AllPermissions()
	if len(catalogPerms) != len(all) {
		t.Fatalf("catalog defines %d keys, but registry has %d permissions", len(catalogPerms), len(all))
	}

	for _, key := range catalogPerms {
		if _, ok := PermissionByID(key); !ok {
			t.Errorf("catalog key %q not found in registered permissions", key)
		}
	}
}

