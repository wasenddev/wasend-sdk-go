package wasendcore


// Group participant information.
type GroupParticipant struct {
	// Whether the participant is an admin.
	IsAdmin *bool `field:"required" json:"isAdmin" yaml:"isAdmin"`
	// Whether the participant is a super admin.
	IsSuperAdmin *bool `field:"required" json:"isSuperAdmin" yaml:"isSuperAdmin"`
	// Participant JID.
	Jid *string `field:"required" json:"jid" yaml:"jid"`
	// Participant phone number.
	PhoneNumber *string `field:"required" json:"phoneNumber" yaml:"phoneNumber"`
	// Add request information (if any).
	AddRequest interface{} `field:"optional" json:"addRequest" yaml:"addRequest"`
	// Participant display name.
	DisplayName *string `field:"optional" json:"displayName" yaml:"displayName"`
	// Error code (if any).
	Error *float64 `field:"optional" json:"error" yaml:"error"`
	// Participant LID (if available).
	Lid *string `field:"optional" json:"lid" yaml:"lid"`
}

