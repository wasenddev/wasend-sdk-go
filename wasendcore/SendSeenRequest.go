package wasendcore


// Send seen request.
type SendSeenRequest struct {
	// The session ID.
	Session *string `field:"required" json:"session" yaml:"session"`
	// The recipient's phone number or group JID.
	//
	// Example:
	//   "+1234567890" or "1234567890-12345678@g.us"
	//
	To *string `field:"required" json:"to" yaml:"to"`
	// The message ID to mark as seen.
	MessageId *string `field:"required" json:"messageId" yaml:"messageId"`
}

