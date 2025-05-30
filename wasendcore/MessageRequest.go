package wasendcore


// Message request structure.
type MessageRequest struct {
	// The content of the message.
	Content *string `field:"required" json:"content" yaml:"content"`
	// The recipient of the message.
	To *string `field:"required" json:"to" yaml:"to"`
}

