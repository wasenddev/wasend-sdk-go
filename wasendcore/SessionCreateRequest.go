package wasendcore


// Session creation request body.
type SessionCreateRequest struct {
	// Phone number for the WhatsApp session.
	PhoneNumber *string `field:"required" json:"phoneNumber" yaml:"phoneNumber"`
	// Name of the session.
	SessionName *string `field:"required" json:"sessionName" yaml:"sessionName"`
	// Enable account protection features.
	// Default: false.
	//
	EnableAccountProtection *bool `field:"optional" json:"enableAccountProtection" yaml:"enableAccountProtection"`
	// Enable message logging.
	// Default: false.
	//
	EnableMessageLogging *bool `field:"optional" json:"enableMessageLogging" yaml:"enableMessageLogging"`
	// Enable webhook notifications.
	// Default: false.
	//
	EnableWebhook *bool `field:"optional" json:"enableWebhook" yaml:"enableWebhook"`
	// Webhook URL for notifications Required if enableWebhook is true.
	WebhookUrl *string `field:"optional" json:"webhookUrl" yaml:"webhookUrl"`
}

