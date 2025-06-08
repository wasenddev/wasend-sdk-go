package wasendcore


// Group settings.
type GroupSettings struct {
	// Whether only admins can edit group info.
	InfoAdminOnly *bool `field:"required" json:"infoAdminOnly" yaml:"infoAdminOnly"`
	// Whether only admins can send messages.
	MessagesAdminOnly *bool `field:"required" json:"messagesAdminOnly" yaml:"messagesAdminOnly"`
	// Whether group is archived.
	Archived *bool `field:"optional" json:"archived" yaml:"archived"`
	// Whether group is muted.
	Muted *bool `field:"optional" json:"muted" yaml:"muted"`
	// Whether group is pinned.
	Pinned *bool `field:"optional" json:"pinned" yaml:"pinned"`
}

