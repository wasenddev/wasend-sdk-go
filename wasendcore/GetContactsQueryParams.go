package wasendcore


// Get contacts query parameters.
type GetContactsQueryParams struct {
	// Maximum number of results.
	// Default: 10.
	//
	Limit *float64 `field:"optional" json:"limit" yaml:"limit"`
	// Number of results to skip.
	// Default: 0.
	//
	Offset *float64 `field:"optional" json:"offset" yaml:"offset"`
	// Sort by field.
	SortBy *string `field:"optional" json:"sortBy" yaml:"sortBy"`
	// Sort order.
	SortOrder *string `field:"optional" json:"sortOrder" yaml:"sortOrder"`
}

