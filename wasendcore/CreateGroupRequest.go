package wasendcore


// Create group request.
type CreateGroupRequest struct {
	// Group name.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Group participants.
	Participants *[]*CreateGroupParticipant `field:"required" json:"participants" yaml:"participants"`
	// Group description (optional).
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Group picture URL (optional).
	PictureUrl *string `field:"optional" json:"pictureUrl" yaml:"pictureUrl"`
	// Group tags (optional).
	Tags *[]*string `field:"optional" json:"tags" yaml:"tags"`
}

