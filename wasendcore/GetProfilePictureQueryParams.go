package wasendcore


// Get profile picture query parameters.
type GetProfilePictureQueryParams struct {
	// Contact ID (phone number with.
	//
	// Example:
	//   "11111111111@c.us"
	//
	ContactId *string `field:"required" json:"contactId" yaml:"contactId"`
	// Refresh the picture from the server (24h cache by default).
	// Default: false.
	//
	Refresh *bool `field:"optional" json:"refresh" yaml:"refresh"`
}

