package perms

// Context represents the scope level at which a permission applies.
type Context string

const (
	// ContextProject indicates a project-level permission.
	ContextProject Context = "project"
	// ContextGlobal indicates a global/system-level permission.
	ContextGlobal Context = "global"
)

// Module represents the functional module a permission belongs to.
type Module string

const (
	ModuleProject      Module = "project"
	ModuleWorkPackages Module = "work_packages"
	ModuleGantt        Module = "gantt"
	ModuleBoards       Module = "boards"
	ModuleBacklogs     Module = "backlogs"
	ModuleBudgets      Module = "budgets"
	ModuleCalendars    Module = "calendars"
	ModuleDocuments    Module = "documents"
	ModuleForums       Module = "forums"
	ModuleGitHub       Module = "github"
	ModuleGitLab       Module = "gitlab"
	ModuleMeetings     Module = "meetings"
	ModuleNews         Module = "news"
	ModuleTeamPlanner  Module = "team_planner"
	ModuleTimeCosts    Module = "time_and_costs"
	ModuleWiki         Module = "wiki"
	ModuleFileStorage  Module = "file_storage"
)

// Dependency describes a prerequisite or behavioral note for a permission.
type Dependency struct {
	// PermissionID is the ID of the permission this depends on (empty if none).
	PermissionID string
	// Description explains the nature of the dependency.
	Description string
}

// Permission defines a single permission entry in OpenProject.
type Permission struct {
	// ID is a unique snake_case identifier for this permission.
	ID string
	// DisplayName is the human-readable name shown in the UI.
	DisplayName string
	// Description explains what this permission allows.
	Description string
	// Context indicates whether this is a project-level or global permission.
	Context Context
	// Module is the functional module this permission belongs to.
	Module Module
	// Dependencies lists any permissions or conditions this permission depends on.
	Dependencies []Dependency
}
