package wasendcore


// Text message request.
type MessageTextRequest struct {
	// The content of the message.
	Content *string `field:"required" json:"content" yaml:"content"`
	// The recipient of the message.
	To *string `field:"required" json:"to" yaml:"to"`
	// The text content of the message.
	Text *string `field:"required" json:"text" yaml:"text"`
}

