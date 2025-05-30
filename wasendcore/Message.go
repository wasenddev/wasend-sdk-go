package wasendcore


// Message data structure.
type Message struct {
	// The content of the message.
	Content *string `field:"required" json:"content" yaml:"content"`
	// Unique identifier for the message.
	Id *string `field:"required" json:"id" yaml:"id"`
	// Timestamp when the message was sent (ISO string).
	SentAt *string `field:"required" json:"sentAt" yaml:"sentAt"`
	// Message status.
	Status *string `field:"required" json:"status" yaml:"status"`
	// The recipient of the message.
	To *string `field:"required" json:"to" yaml:"to"`
}

