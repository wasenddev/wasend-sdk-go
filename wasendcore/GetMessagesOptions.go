package wasendcore


type GetMessagesOptions struct {
	Limit *float64 `field:"required" json:"limit" yaml:"limit"`
	DownloadMedia *bool `field:"optional" json:"downloadMedia" yaml:"downloadMedia"`
	Filter *GetMessagesFilterOptions `field:"optional" json:"filter" yaml:"filter"`
	Offset *float64 `field:"optional" json:"offset" yaml:"offset"`
}

