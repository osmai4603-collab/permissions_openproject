package perms

import (
	"sort"
)

// IPermissionService defines the interface for managing and evaluating OpenProject permissions.
type IPermissionService interface {
	GetPermission(id string) *Permission
	GetAllPermissions() []Permission
	GetPermissionsByContext(ctx Context) []Permission
	GetPermissionsByModule(module Module) []Permission
	ResolveRevocation(activePermissionIDs []string, permissionToRemove string) []string
	ValidatePermissionsForContext(permissionIDs []string, targetContext Context) []string
	HasPermission(grantedPermissions []string, requiredPermission string) bool
}

func GetPermission(id string) *Permission {
	perms := GetDefaultPermissions()
	for _, perm := range perms {
		if perm.ID == id {
			return &perm
		}
	}
	return nil
}

// Service provides OpenProject permission management, validation, and evaluation.
type Service struct {
	catalog     map[string]Permission
	listOrdered []Permission
}

// NewService instantiates a new permissions service populated with the OpenProject catalog.
func NewService() *Service {
	all := GetDefaultPermissions()
	cat := make(map[string]Permission, len(all))
	for _, p := range all {
		cat[p.ID] = p
	}
	list := make([]Permission, len(all))
	copy(list, all)

	sort.Slice(list, func(i, j int) bool {
		if list[i].Module != list[j].Module {
			return list[i].Module < list[j].Module
		}
		return list[i].ID < list[j].ID
	})

	return &Service{
		catalog:     cat,
		listOrdered: list,
	}
}

// GetPermission returns a permission definition by its ID.
func (s *Service) GetPermission(id string) *Permission {
	p, ok := s.catalog[id]
	if ok {
		return &p
	}
	return nil
}

// GetAllPermissions returns all registered permissions in deterministic order.
func (s *Service) GetAllPermissions() []Permission {
	out := make([]Permission, len(s.listOrdered))
	copy(out, s.listOrdered)
	return out
}

// GetPermissionsByContext retrieves all permissions applicable to a specific context level.
func (s *Service) GetPermissionsByContext(ctx Context) []Permission {
	var res []Permission
	for _, p := range s.listOrdered {
		if p.AllowsContext(ctx) || s.isAllowedInContext(p, ctx) {
			res = append(res, p)
		}
	}
	return res
}

// GetPermissionsByModule retrieves permissions belonging to a specific functional module.
func (s *Service) GetPermissionsByModule(module Module) []Permission {
	var res []Permission
	for _, p := range s.listOrdered {
		if p.Module == module {
			res = append(res, p)
		}
	}
	return res
}

// ResolveRevocation determines which permissions remain after revoking a target permission.
// When removing a permission that other permissions depend on, dependent permissions are transitively revoked.
func (s *Service) ResolveRevocation(activePermissionIDs []string, permissionToRemove string) []string {
	activeSet := make(map[string]struct{}, len(activePermissionIDs))
	for _, id := range activePermissionIDs {
		activeSet[id] = struct{}{}
	}

	// Queue permissions to delete (starting with the target)
	toDelete := make(map[string]struct{})
	queue := []string{permissionToRemove}
	toDelete[permissionToRemove] = struct{}{}

	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]

		// Find any permission that directly or transitively depends on curr
		for _, p := range s.catalog {
			for _, dep := range p.Dependencies {
				if dep.PermissionID == curr {
					if _, alreadyMarked := toDelete[p.ID]; !alreadyMarked {
						toDelete[p.ID] = struct{}{}
						queue = append(queue, p.ID)
					}
				}
			}
		}
	}

	var remaining []string
	for id := range activeSet {
		if _, deleted := toDelete[id]; !deleted {
			remaining = append(remaining, id)
		}
	}
	sort.Strings(remaining)
	return remaining
}

// ValidatePermissionsForContext filters permissions based on OpenProject scope context isolation rules:
// - ContextGlobal: Only global permissions apply (GlobalRole).
// - ContextProject: Project-level permissions apply (ProjectRole).
// - ContextWorkPackage: Only permissions permissible on WorkPackage entity sharing apply (WorkPackageRole).
// - ContextProjectQuery: Only permissions permissible on saved ProjectQuery views apply.
func (s *Service) ValidatePermissionsForContext(permissionIDs []string, targetContext Context) []string {
	var valid []string
	for _, id := range permissionIDs {
		p, ok := s.catalog[id]
		if !ok {
			continue
		}

		switch targetContext {
		case ContextGlobal:
			if p.AllowsContext(ContextGlobal) {
				valid = append(valid, id)
			}
		case ContextProject:
			if p.AllowsContext(ContextProject) || s.isAllowedInContext(p, ContextProject) {
				valid = append(valid, id)
			}
		case ContextWorkPackage:
			if p.AllowsContext(ContextWorkPackage) || s.isAllowedInContext(p, ContextWorkPackage) {
				valid = append(valid, id)
			}
		case ContextProjectQuery:
			if p.AllowsContext(ContextProjectQuery) || s.isAllowedInContext(p, ContextProjectQuery) {
				valid = append(valid, id)
			}
		}
	}
	sort.Strings(valid)
	return valid
}

// HasPermission checks if the required permission is present in the list of granted permissions.
func (s *Service) HasPermission(grantedPermissions []string, requiredPermission string) bool {
	for _, p := range grantedPermissions {
		if p == requiredPermission {
			return true
		}
	}
	return false
}

func (s *Service) isAllowedInContext(p Permission, ctx Context) bool {
	for _, permissible := range p.PermissibleOn {
		if permissible == ctx {
			return true
		}
	}
	return false
}
