package perms

import "slices"

// ──────────────────────────────────────────────
// Permission registry — aggregation and lookup
// ──────────────────────────────────────────────

// AllPermissions returns every registered permission across all modules.
func AllPermissions() []Permission {
	var all []Permission
	all = append(all, GlobalPermissions()...)
	all = append(all, ProjectPermissions()...)
	all = append(all, WorkPackagePermissions()...)
	all = append(all, BoardsPermissions()...)
	all = append(all, BacklogsPermissions()...)
	all = append(all, BudgetsPermissions()...)
	all = append(all, CalendarsPermissions()...)
	all = append(all, DocumentsPermissions()...)
	all = append(all, ForumsPermissions()...)
	all = append(all, GitHubPermissions()...)
	all = append(all, GitLabPermissions()...)
	all = append(all, MeetingsPermissions()...)
	all = append(all, NewsPermissions()...)
	all = append(all, TeamPlannerPermissions()...)
	all = append(all, TimeCostsPermissions()...)
	all = append(all, WikiPermissions()...)
	return all
}

// PermissionByID looks up a single permission by its unique ID.
// Returns the permission and true if found, or a zero Permission and false otherwise.
func PermissionByID(id string) (Permission, bool) {
	for _, p := range AllPermissions() {
		if p.ID == id {
			return p, true
		}
	}
	return Permission{}, false
}

// PermissionsByModule returns all permissions belonging to the given module.
func PermissionsByModule(module Module) []Permission {
	var result []Permission
	for _, p := range AllPermissions() {
		if p.Module == module {
			result = append(result, p)
		}
	}
	return result
}

// PermissionsByContext returns all permissions for the given context (project or global).
func PermissionsByContext(ctx Context) []Permission {
	var result []Permission
	for _, p := range AllPermissions() {
		if slices.Contains(p.PermissibleOn, ctx) {
			result = append(result, p)
		}
	}
	return result
}
