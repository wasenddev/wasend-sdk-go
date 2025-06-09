package wasendcore


type ReadMessagesOptions struct {
	Days *float64 `field:"optional" json:"days" yaml:"days"`
	Messages *float64 `field:"optional" json:"messages" yaml:"messages"`
}

