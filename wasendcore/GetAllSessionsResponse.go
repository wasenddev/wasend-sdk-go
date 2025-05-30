package wasendcore


// Response wrapper for getAllSessions.
type GetAllSessionsResponse struct {
	// Array of sessions.
	Sessions *[]*SessionListItem `field:"required" json:"sessions" yaml:"sessions"`
}

