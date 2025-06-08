package wasendcore


// Subject request.
type SubjectRequest struct {
	// Group subject/name.
	Subject *string `field:"required" json:"subject" yaml:"subject"`
}

