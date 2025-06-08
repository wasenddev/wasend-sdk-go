package wasendcore


// Join group response.
type JoinGroupResponse struct {
	// Group ID.
	Id *string `field:"required" json:"id" yaml:"id"`
	// Group name.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Group participants count.
	ParticipantsCount *float64 `field:"required" json:"participantsCount" yaml:"participantsCount"`
	// Group description.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Group picture URL.
	PictureUrl *string `field:"optional" json:"pictureUrl" yaml:"pictureUrl"`
}

