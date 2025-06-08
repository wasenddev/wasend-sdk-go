package wasendcore


// Get contact query parameters.
type GetContactQueryParams struct {
	// Contact ID (phone number with.
	//
	// Example:
	//   "11111111111@c.us"
	//
	ContactId *string `field:"required" json:"contactId" yaml:"contactId"`
}

