package wasendcore


type GetChatsOverviewOptions struct {
	Ids *[]*string `field:"optional" json:"ids" yaml:"ids"`
	Limit *float64 `field:"optional" json:"limit" yaml:"limit"`
	Offset *float64 `field:"optional" json:"offset" yaml:"offset"`
}

