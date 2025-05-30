package wasendcore


// Session information from createSession and other operations.
type Session struct {
	// Creation timestamp.
	CreatedAt *string `field:"required" json:"createdAt" yaml:"createdAt"`
	// Session configuration flags.
	EnableAccountProtection *bool `field:"required" json:"enableAccountProtection" yaml:"enableAccountProtection"`
	EnableMessageLogging *bool `field:"required" json:"enableMessageLogging" yaml:"enableMessageLogging"`
	EnableWebhook *bool `field:"required" json:"enableWebhook" yaml:"enableWebhook"`
	// MongoDB ID of the session.
	Id *string `field:"required" json:"id" yaml:"id"`
	// Phone number associated with the session.
	PhoneNumber *string `field:"required" json:"phoneNumber" yaml:"phoneNumber"`
	// Name of the session.
	SessionName *string `field:"required" json:"sessionName" yaml:"sessionName"`
	// Unique session identifier.
	UniqueSessionId *string `field:"required" json:"uniqueSessionId" yaml:"uniqueSessionId"`
	// Last update timestamp.
	UpdatedAt *string `field:"required" json:"updatedAt" yaml:"updatedAt"`
	// User ID who owns the session.
	UserId *string `field:"required" json:"userId" yaml:"userId"`
	WebhookUrl *string `field:"optional" json:"webhookUrl" yaml:"webhookUrl"`
	// Session configuration.
	Config *SessionConfig `field:"required" json:"config" yaml:"config"`
	// Engine status information.
	Engine *EngineStatus `field:"required" json:"engine" yaml:"engine"`
	// Name of the session.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Current status of the session.
	Status *string `field:"required" json:"status" yaml:"status"`
}

