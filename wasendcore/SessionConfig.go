package wasendcore


// Session configuration.
type SessionConfig struct {
	// Debug mode.
	Debug *bool `field:"required" json:"debug" yaml:"debug"`
	// User metadata.
	Metadata *SessionMetadata `field:"required" json:"metadata" yaml:"metadata"`
	// Noweb configuration.
	Noweb *NowebConfig `field:"required" json:"noweb" yaml:"noweb"`
	// Proxy configuration.
	Proxy interface{} `field:"required" json:"proxy" yaml:"proxy"`
	// Webhook configurations.
	Webhooks *[]interface{} `field:"required" json:"webhooks" yaml:"webhooks"`
}

