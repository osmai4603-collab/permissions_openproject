package perms

import (
	"reflect"
	"sync"
	"testing"
)

// TestComplexTransitiveClosure_DeepChains tests deep transitive implication trees,
// graph idempotency, and safe handling of duplicate or overlapping inputs.
func TestComplexTransitiveClosure_DeepChains(t *testing.T) {
	svc := NewService()

	t.Run("Deep Chain: EditTimeLogsForOtherUsers -> EditOwnTimeLogs -> ViewOwnSpentTime", func(t *testing.T) {
		input := []string{PermEditTimeLogsForOtherUsers}
		resolved := svc.ResolveImplied(input)

		expectedSubset := []string{
			PermEditTimeLogsForOtherUsers,
			PermEditOwnTimeLogs,
			PermViewSpentTime,
			PermViewOwnSpentTime,
		}

		for _, exp := range expectedSubset {
			found := false
			for _, r := range resolved {
				if r == exp {
					found = true
					break
				}
			}
			if !found {
				t.Errorf("expected permission %s to be resolved in deep chain, but was missing in %v", exp, resolved)
			}
		}
	})

	t.Run("Multi-Root Overlapping Inputs", func(t *testing.T) {
		// Both EditWorkPackages and DeleteWorkPackages require ViewWorkPackages
		input := []string{PermEditWorkPackages, PermDeleteWorkPackages}
		resolved := svc.ResolveImplied(input)

		// Check for duplicates
		seen := make(map[string]int)
		for _, p := range resolved {
			seen[p]++
			if seen[p] > 1 {
				t.Errorf("duplicate permission %s found in resolved implied list", p)
			}
		}

		if seen[PermViewWorkPackages] != 1 {
			t.Errorf("expected exactly 1 instance of %s, got %d", PermViewWorkPackages, seen[PermViewWorkPackages])
		}
	})

	t.Run("Idempotency: Resolve(Resolve(X)) == Resolve(X)", func(t *testing.T) {
		inputs := [][]string{
			{PermEditWorkPackages},
			{PermManageWiki},
			{PermCopyProjects, PermEditWorkPackages},
			{PermEditTimeLogsForOtherUsers, PermEditHourlyRates},
		}

		for _, input := range inputs {
			pass1 := svc.ResolveImplied(input)
			pass2 := svc.ResolveImplied(pass1)

			if !reflect.DeepEqual(pass1, pass2) {
				t.Errorf("resolve operation not idempotent: pass1=%v, pass2=%v", pass1, pass2)
			}
		}
	})
}

// TestScopeIsolationIntegrity verifies that permissions cannot leak across scopes.
func TestScopeIsolationIntegrity(t *testing.T) {
	svc := NewService()
	allPerms := svc.GetAllPermissions()
	allIDs := make([]string, len(allPerms))
	for i, p := range allPerms {
		allIDs[i] = p.ID
	}

	t.Run("ScopeGlobal Strictly Contains Only Global Permissions", func(t *testing.T) {
		res := svc.ValidatePermissionsForScope(allIDs, ScopeGlobal)
		for _, id := range res {
			p, _ := svc.GetPermission(id)
			if !p.IsGlobal && p.Scope != ScopeGlobal {
				t.Fatalf("non-global permission %s leaked into ScopeGlobal", id)
			}
		}
	})

	t.Run("ScopeProject Does Not Contain Pure Global Permissions", func(t *testing.T) {
		res := svc.ValidatePermissionsForScope(allIDs, ScopeProject)
		for _, id := range res {
			p, _ := svc.GetPermission(id)
			if p.Scope == ScopeGlobal {
				t.Fatalf("global permission %s leaked into ScopeProject", id)
			}
		}
	})

	t.Run("ScopeWorkPackage Strictly Contains Only WorkPackage Sharable Permissions", func(t *testing.T) {
		res := svc.ValidatePermissionsForScope(allIDs, ScopeWorkPackage)
		for _, id := range res {
			p, _ := svc.GetPermission(id)
			isSharable := false
			for _, allowed := range p.AllowedScopes {
				if allowed == ScopeWorkPackage {
					isSharable = true
					break
				}
			}
			if !isSharable && p.Scope != ScopeWorkPackage {
				t.Fatalf("non-work-package permission %s leaked into ScopeWorkPackage", id)
			}
		}
	})

	t.Run("Empty Input Slice Handles Safely", func(t *testing.T) {
		res := svc.ValidatePermissionsForScope([]string{}, ScopeProject)
		if len(res) != 0 {
			t.Errorf("expected empty slice, got %v", res)
		}
	})
}

// TestConcurrentAccessAndThreadSafety ensures that Service is completely race-free
// and thread-safe when accessed by multiple concurrent goroutines.
func TestConcurrentAccessAndThreadSafety(t *testing.T) {
	svc := NewService()
	allPerms := svc.GetAllPermissions()

	const workers = 50
	const iterations = 50

	var wg sync.WaitGroup
	wg.Add(workers)

	for w := 0; w < workers; w++ {
		go func(workerID int) {
			defer wg.Done()
			for i := 0; i < iterations; i++ {
				// Concurrent catalog reads
				idx := (workerID + i) % len(allPerms)
				targetID := allPerms[idx].ID

				p, ok := svc.GetPermission(targetID)
				if !ok || p.ID != targetID {
					t.Errorf("worker %d: expected to find permission %s", workerID, targetID)
				}

				// Concurrent Implied Resolution
				implied := svc.ResolveImplied([]string{targetID})
				if len(implied) == 0 {
					t.Errorf("worker %d: implied resolution produced empty set for %s", workerID, targetID)
				}

				// Concurrent Revocation Resolution
				remaining := svc.ResolveRevocation(implied, targetID)
				for _, rem := range remaining {
					if rem == targetID {
						t.Errorf("worker %d: target %s was not revoked", workerID, targetID)
					}
				}

				// Concurrent Scope Filtering
				scoped := svc.ValidatePermissionsForScope(implied, ScopeProject)
				for _, sc := range scoped {
					perm, _ := svc.GetPermission(sc)
					if perm.Scope == ScopeGlobal {
						t.Errorf("worker %d: global perm %s leaked into project scope", workerID, sc)
					}
				}
			}
		}(w)
	}

	wg.Wait()
}

// Benchmarks

func BenchmarkResolveImplied(b *testing.B) {
	svc := NewService()
	input := []string{PermCopyProjects, PermEditWorkPackages, PermEditTimeLogsForOtherUsers}

	for b.Loop() {
		_ = svc.ResolveImplied(input)
	}
}

func BenchmarkResolveRevocation(b *testing.B) {
	svc := NewService()
	active := svc.ResolveImplied([]string{PermEditWorkPackages, PermDeleteWorkPackages})

	for b.Loop() {
		_ = svc.ResolveRevocation(active, PermViewWorkPackages)
	}
}

func BenchmarkValidatePermissionsForScope(b *testing.B) {
	svc := NewService()
	allPerms := svc.GetAllPermissions()
	ids := make([]string, len(allPerms))
	for i, p := range allPerms {
		ids[i] = p.ID
	}

	for b.Loop() {
		_ = svc.ValidatePermissionsForScope(ids, ScopeProject)
	}
}
