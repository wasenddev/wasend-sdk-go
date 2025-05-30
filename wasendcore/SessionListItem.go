package wasendcore


// Session information from getAllSessions.
type SessionListItem struct {
	// Downstream connection information.
	Downstream *DownstreamInfo `field:"required" json:"downstream" yaml:"downstream"`
	// Session details.
	Session *SessionDetails `field:"required" json:"session" yaml:"session"`
}

