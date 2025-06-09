package wasendcore


type ChatOverview struct {
	Id *string `field:"required" json:"id" yaml:"id"`
	IsGroup *bool `field:"optional" json:"isGroup" yaml:"isGroup"`
	Name *string `field:"optional" json:"name" yaml:"name"`
	Timestamp *float64 `field:"optional" json:"timestamp" yaml:"timestamp"`
	UnreadCount *float64 `field:"optional" json:"unreadCount" yaml:"unreadCount"`
}

