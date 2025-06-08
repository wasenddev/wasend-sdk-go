package wasendcore


// Participants request.
type ParticipantsRequest struct {
	// List of participant IDs.
	//
	// Example:
	//   ["+919545251359", "+919545251360"]
	//
	Participants *[]*string `field:"required" json:"participants" yaml:"participants"`
	// Whether to notify participants.
	// Default: true.
	//
	Notify *bool `field:"optional" json:"notify" yaml:"notify"`
}

