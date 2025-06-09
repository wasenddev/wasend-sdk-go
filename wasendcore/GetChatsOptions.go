package wasendcore


type GetChatsOptions struct {
	Limit *float64 `field:"optional" json:"limit" yaml:"limit"`
	Offset *float64 `field:"optional" json:"offset" yaml:"offset"`
	SortBy *string `field:"optional" json:"sortBy" yaml:"sortBy"`
	SortOrder *string `field:"optional" json:"sortOrder" yaml:"sortOrder"`
}

