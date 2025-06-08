package wasendcore


// Get groups query parameters.
type GetGroupsQueryParams struct {
	// Fields to exclude from response.
	Exclude *[]*string `field:"optional" json:"exclude" yaml:"exclude"`
	// Maximum number of results.
	// Default: 50.
	//
	Limit *float64 `field:"optional" json:"limit" yaml:"limit"`
	// Number of results to skip.
	// Default: 0.
	//
	Offset *float64 `field:"optional" json:"offset" yaml:"offset"`
	// Search by name or description.
	Search *string `field:"optional" json:"search" yaml:"search"`
	// Sort by field.
	SortBy *string `field:"optional" json:"sortBy" yaml:"sortBy"`
	// Sort order.
	SortOrder *string `field:"optional" json:"sortOrder" yaml:"sortOrder"`
	// Filter by status.
	Status GroupStatus `field:"optional" json:"status" yaml:"status"`
	// Filter by tags.
	Tags *[]*string `field:"optional" json:"tags" yaml:"tags"`
}

