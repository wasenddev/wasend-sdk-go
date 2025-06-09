package wasendcore


type Chat struct {
	Id *string `field:"required" json:"id" yaml:"id"`
	Archived *bool `field:"optional" json:"archived" yaml:"archived"`
	IsGroup *bool `field:"optional" json:"isGroup" yaml:"isGroup"`
	Name *string `field:"optional" json:"name" yaml:"name"`
	Pinned *bool `field:"optional" json:"pinned" yaml:"pinned"`
	Timestamp *float64 `field:"optional" json:"timestamp" yaml:"timestamp"`
	UnreadCount *float64 `field:"optional" json:"unreadCount" yaml:"unreadCount"`
}

