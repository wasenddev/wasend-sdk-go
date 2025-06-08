package wasendcore


// Create group request participant.
type CreateGroupParticipant struct {
	// Participant ID (phone number with country code).
	//
	// Example:
	//   "+919545251359"
	//
	Id *string `field:"required" json:"id" yaml:"id"`
	// Whether to make the participant an admin.
	// Default: false.
	//
	IsAdmin *bool `field:"optional" json:"isAdmin" yaml:"isAdmin"`
}

