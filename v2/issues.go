package v2

// Record types a plugin can attach extension data to (Extensions). The host
// accepts only the types it has registered; others answer an error.
const (
	EntityTypeIssue = "issue" // tickets, tasks and todos (issueId)
)

// Status categories: what an organization's issue status means, the same for
// every organization.
const (
	IssueStatusCategoryTodo       = "todo"
	IssueStatusCategoryInProgress = "inProgress"
	IssueStatusCategoryWaiting    = "waiting"
	IssueStatusCategoryDone       = "done"
)

// IssueStatus is one of the organization's ticket/task statuses, named as in
// the external API GET /me/issue-statuses.
type IssueStatus struct {
	StatusId    uint   `json:"statusId"`
	Name        string `json:"name"`
	Color       string `json:"color"`
	IsCompleted bool   `json:"isCompleted"` // moving an issue here closes it
	Active      bool   `json:"active"`
	OrderBy     uint   `json:"orderBy"`
	Category    string `json:"category"`
}

// IssuesAPI is the host's tickets and tasks.
type IssuesAPI interface {
	// SetIssueStatus moves one of the context organization's open issues to
	// another active status, for automation such as "merge request opened →
	// review". The status change activity records the plugin and the
	// organization's afterUpdate workflows run as for a person's change.
	// Closing stays a person's action: statuses in the done category,
	// completing statuses and closed issues are refused.
	SetIssueStatus(context RequestContext, issueId uint, statusId uint) error
	// GetIssueStatuses lists the organization's statuses, for plugin settings
	// where the organization picks the status for each plugin stage. Choose by
	// id and category, never by name.
	GetIssueStatuses(context RequestContext) ([]IssueStatus, error)
}

// ExtensionsAPI keeps a plugin's summary on records of any registered type.
type ExtensionsAPI interface {
	// SetEntityExtensionData stores the plugin's summary on one of the context
	// organization's records (entityType: EntityTypeIssue, …), replacing what
	// it stored before; nil removes it. data must marshal to a JSON object (at
	// most 16 KB). It is a plugin-written copy for display and agents: people
	// do not edit it and lists never filter or search it (declare a custom
	// field for that). For an issue the external API and get_issue return it
	// under extensions.<pluginId> = {data, updatedAt}, never in the customer
	// portal, only while the plugin is enabled; the host deletes it with the
	// record.
	SetEntityExtensionData(context RequestContext, entityType string, entityId uint, data any) error
}
