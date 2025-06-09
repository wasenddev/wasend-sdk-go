package wasendcore


type PaginationOptions struct {
	Limit *float64 `field:"optional" json:"limit" yaml:"limit"`
	Offset *float64 `field:"optional" json:"offset" yaml:"offset"`
}

