package wasendcore


// Check contact exists query parameters.
type CheckContactExistsQueryParams struct {
	// The phone number to check.
	//
	// Example:
	//   "1213213213"
	//
	Phone *string `field:"required" json:"phone" yaml:"phone"`
}

