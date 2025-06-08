package wasendcore


// WhatsApp message response.
type WAMessage struct {
	// The message content.
	Content *string `field:"required" json:"content" yaml:"content"`
	// Whether the message is from me.
	FromMe *bool `field:"required" json:"fromMe" yaml:"fromMe"`
	// The message ID.
	Id *string `field:"required" json:"id" yaml:"id"`
	// The message recipient.
	Recipient *string `field:"required" json:"recipient" yaml:"recipient"`
	// The message sender.
	Sender *string `field:"required" json:"sender" yaml:"sender"`
	// The message status.
	Status *string `field:"required" json:"status" yaml:"status"`
	// The message timestamp.
	Timestamp *string `field:"required" json:"timestamp" yaml:"timestamp"`
	// The recipient's phone number or group JID.
	To *string `field:"required" json:"to" yaml:"to"`
	// The message type.
	Type *string `field:"required" json:"type" yaml:"type"`
}

