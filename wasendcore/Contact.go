package wasendcore


// Contact information.
type Contact struct {
	// Contact ID (phone number with.
	//
	// Example:
	//   "11111111111@c.us"
	//
	Id *string `field:"required" json:"id" yaml:"id"`
	// Contact phone number.
	PhoneNumber *string `field:"required" json:"phoneNumber" yaml:"phoneNumber"`
	// Contact business profile.
	BusinessProfile *map[string]*string `field:"optional" json:"businessProfile" yaml:"businessProfile"`
	// Contact name.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Contact profile picture URL.
	ProfilePictureUrl *string `field:"optional" json:"profilePictureUrl" yaml:"profilePictureUrl"`
	// Contact push name.
	PushName *string `field:"optional" json:"pushName" yaml:"pushName"`
	// Contact status.
	Status *string `field:"optional" json:"status" yaml:"status"`
}

