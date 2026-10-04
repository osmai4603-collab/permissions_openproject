package perms

// Context represents the scope/context level at which a permission applies in OpenProject.
type Context string

const (
	// ContextProject indicates a project-level permission.
	ContextProject Context = "project"
	// ContextGlobal indicates a platform-wide system-level permission.
	ContextGlobal Context = "global"
	// ContextWorkPackage indicates a granular work-package level permission (e.g. Work Package Sharing).
	ContextWorkPackage Context = "work_package"
	// ContextProjectQuery indicates a granular saved project query permission.
	ContextProjectQuery Context = "project_query"
)

// Module represents the functional module a permission belongs to in OpenProject.
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

// OperationType defines the operational or CRUD action being performed.
type OperationType string

const (
	OpCreate  OperationType = "CREATE"
	OpRead    OperationType = "READ"
	OpUpdate  OperationType = "UPDATE"
	OpDelete  OperationType = "DELETE"
	OpManage  OperationType = "MANAGE"
	OpShare   OperationType = "SHARE"
	OpAdmin   OperationType = "ADMIN"
	OpSpecial OperationType = "SPECIAL"
)

// Permission defines a single authorization rule in OpenProject.
type Permission struct {
	ID             string        `json:"id"`
	DisplayName    string        `json:"display_name"`
	Description    string        `json:"description"`
	IsGlobal       bool          `json:"is_global"`
	Module         Module        `json:"module"`
	Context        Context       `json:"context"`
	PermissibleOn  []Context     `json:"permissible_on,omitempty"`
	Operation      OperationType `json:"operation"`
	ImpliedPerms   []string      `json:"implied_perms,omitempty"`
	DependentPerms []string      `json:"dependent_perms,omitempty"`
}
