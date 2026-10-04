package perms

import (
	"sort"
	"testing"
)

func TestBuildDefaultCatalog(t *testing.T) {
	svc := NewService()
	perms := svc.GetAllPermissions()

	if len(perms) < 100 {
		t.Fatalf("expected at least 100 permissions in OpenProject catalog, got %d", len(perms))
	}

	for _, p := range perms {
		if p.ID == "" {
			t.Errorf("permission with empty ID found")
		}
		if p.DisplayName == "" {
			t.Errorf("permission %s has empty DisplayName", p.ID)
		}
		if p.Description == "" {
			t.Errorf("permission %s has empty Description", p.ID)
		}
		if p.Module == "" {
			t.Errorf("permission %s has empty Module", p.ID)
		}
		if p.Context == "" {
			t.Errorf("permission %s has empty Context", p.ID)
		}
		if p.Operation == "" {
			t.Errorf("permission %s has empty Operation", p.ID)
		}
	}
}

func TestGetPermission(t *testing.T) {
	svc := NewService()

	p, ok := svc.GetPermission(PermViewWorkPackages)
	if !ok {
		t.Fatalf("expected to find %s in catalog", PermViewWorkPackages)
	}
	if p.ID != PermViewWorkPackages {
		t.Errorf("expected ID %s, got %s", PermViewWorkPackages, p.ID)
	}
	if p.Module != ModuleWorkPackages {
		t.Errorf("expected module %s, got %s", ModuleWorkPackages, p.Module)
	}

	_, notFound := svc.GetPermission("non_existent_permission")
	if notFound {
		t.Errorf("expected false for nonexistent permission")
	}
}

func TestGetAllPermissions(t *testing.T) {
	svc := NewService()
	all := svc.GetAllPermissions()

	if len(all) == 0 {
		t.Fatal("expected non-empty list of permissions")
	}

	// Verify deterministic sorting: Module asc, then ID asc
	for i := 1; i < len(all); i++ {
		prev := all[i-1]
		curr := all[i]
		if prev.Module > curr.Module {
			t.Fatalf("list not sorted by Module: %s > %s", prev.Module, curr.Module)
		}
		if prev.Module == curr.Module && prev.ID > curr.ID {
			t.Fatalf("list not sorted by ID within module %s: %s > %s", prev.Module, prev.ID, curr.ID)
		}
	}
}

func TestGetPermissionsByContext(t *testing.T) {
	svc := NewService()

	globals := svc.GetPermissionsByContext(ContextGlobal)
	if len(globals) == 0 {
		t.Fatal("expected global permissions to be present")
	}
	for _, p := range globals {
		if !p.IsGlobal && p.Context != ContextGlobal {
			t.Errorf("permission %s is not global", p.ID)
		}
	}

	projects := svc.GetPermissionsByContext(ContextProject)
	if len(projects) == 0 {
		t.Fatal("expected project permissions to be present")
	}

	wpSharing := svc.GetPermissionsByContext(ContextWorkPackage)
	if len(wpSharing) == 0 {
		t.Fatal("expected work package sharable permissions to be present")
	}
}

func TestGetPermissionsByModule(t *testing.T) {
	svc := NewService()

	wpPerms := svc.GetPermissionsByModule(ModuleWorkPackages)
	if len(wpPerms) == 0 {
		t.Fatal("expected work package permissions")
	}
	for _, p := range wpPerms {
		if p.Module != ModuleWorkPackages {
			t.Errorf("expected module %s, got %s for perm %s", ModuleWorkPackages, p.Module, p.ID)
		}
	}

	wikiPerms := svc.GetPermissionsByModule(ModuleWiki)
	if len(wikiPerms) == 0 {
		t.Fatal("expected wiki permissions")
	}
	for _, p := range wikiPerms {
		if p.Module != ModuleWiki {
			t.Errorf("expected module %s, got %s for perm %s", ModuleWiki, p.Module, p.ID)
		}
	}
}

func TestResolveImplied(t *testing.T) {
	svc := NewService()

	tests := []struct {
		name     string
		input    []string
		expected []string
	}{
		{
			name:  "Edit Work Packages implies View Work Packages",
			input: []string{PermEditWorkPackages},
			expected: func() []string {
				exp := []string{PermEditWorkPackages, PermViewWorkPackages}
				sort.Strings(exp)
				return exp
			}(),
		},
		{
			name:  "Manage Wiki implies Edit Wiki Pages and View Wiki",
			input: []string{PermManageWiki},
			expected: func() []string {
				exp := []string{PermManageWiki, PermEditWikiPages, PermViewWiki}
				sort.Strings(exp)
				return exp
			}(),
		},
		{
			name:  "Copy Projects implies Edit Project and View Project Attributes",
			input: []string{PermCopyProjects},
			expected: func() []string {
				exp := []string{PermCopyProjects, PermEditProject, PermViewProjectAttributes}
				sort.Strings(exp)
				return exp
			}(),
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			res := svc.ResolveImplied(tc.input)
			if len(res) != len(tc.expected) {
				t.Fatalf("expected %v, got %v", tc.expected, res)
			}
			for i, exp := range tc.expected {
				if res[i] != exp {
					t.Errorf("at index %d: expected %s, got %s", i, exp, res[i])
				}
			}
		})
	}
}

func TestResolveRevocation(t *testing.T) {
	svc := NewService()

	// Initial set: EditWorkPackages and DeleteWorkPackages, both implying ViewWorkPackages
	active := svc.ResolveImplied([]string{PermEditWorkPackages, PermDeleteWorkPackages})

	// Revoke ViewWorkPackages
	remaining := svc.ResolveRevocation(active, PermViewWorkPackages)

	// Since EditWorkPackages and DeleteWorkPackages depend on ViewWorkPackages,
	// all of them must be dropped
	for _, p := range remaining {
		if p == PermEditWorkPackages || p == PermDeleteWorkPackages || p == PermViewWorkPackages {
			t.Errorf("revocation failed: permission %s is still present in remaining list", p)
		}
	}

	if len(remaining) != 0 {
		t.Errorf("expected 0 remaining permissions, got %v", remaining)
	}
}

func TestValidatePermissionsForContext(t *testing.T) {
	svc := NewService()

	mix := []string{
		PermCreateUser,       // Global
		PermAddProject,       // Global
		PermEditProject,      // Project
		PermArchiveProject,   // Project
		PermViewWorkPackages, // Project + WorkPackage
		PermEditWorkPackages, // Project + WorkPackage
	}

	// 1. Global Context: Only global permissions apply
	globalValid := svc.ValidatePermissionsForContext(mix, ContextGlobal)
	for _, id := range globalValid {
		p, _ := svc.GetPermission(id)
		if !p.IsGlobal && p.Context != ContextGlobal {
			t.Errorf("non-global permission %s leaked into Global context", id)
		}
	}
	if len(globalValid) != 2 {
		t.Errorf("expected 2 global permissions, got %d (%v)", len(globalValid), globalValid)
	}

	// 2. Project Context: Project permissions apply
	projValid := svc.ValidatePermissionsForContext(mix, ContextProject)
	for _, id := range projValid {
		p, _ := svc.GetPermission(id)
		if p.Context == ContextGlobal {
			t.Errorf("global permission %s leaked into Project context", id)
		}
	}
	if len(projValid) != 4 {
		t.Errorf("expected 4 project permissions, got %d (%v)", len(projValid), projValid)
	}

	// 3. Work Package Context: Only work-package sharable permissions apply
	wpValid := svc.ValidatePermissionsForContext(mix, ContextWorkPackage)
	for _, id := range wpValid {
		if id != PermViewWorkPackages && id != PermEditWorkPackages {
			t.Errorf("permission %s not valid in WorkPackage context", id)
		}
	}
	if len(wpValid) != 2 {
		t.Errorf("expected 2 work package permissions, got %d (%v)", len(wpValid), wpValid)
	}
}

func TestHasPermission(t *testing.T) {
	svc := NewService()

	granted := []string{PermViewWorkPackages, PermEditWorkPackages}
	if !svc.HasPermission(granted, PermViewWorkPackages) {
		t.Errorf("expected true for granted permission")
	}
	if svc.HasPermission(granted, PermDeleteWorkPackages) {
		t.Errorf("expected false for ungranted permission")
	}
}
