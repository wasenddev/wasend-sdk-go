package wasendcore


// Settings security change info.
type SettingsSecurityChangeInfo struct {
	// Whether the setting is enabled.
	Enabled *bool `field:"required" json:"enabled" yaml:"enabled"`
	// When the setting was last changed.
	ChangedAt *string `field:"optional" json:"changedAt" yaml:"changedAt"`
	// Who changed the setting.
	ChangedBy *string `field:"optional" json:"changedBy" yaml:"changedBy"`
}

