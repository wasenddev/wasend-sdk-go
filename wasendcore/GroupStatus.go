package wasendcore


// Group status.
type GroupStatus string

const (
	// Group is active.
	GroupStatus_ACTIVE GroupStatus = "ACTIVE"
	// Group is archived.
	GroupStatus_ARCHIVED GroupStatus = "ARCHIVED"
	// Group is deleted.
	GroupStatus_DELETED GroupStatus = "DELETED"
)

