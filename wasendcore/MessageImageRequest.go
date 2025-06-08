package wasendcore


// Image message request.
type MessageImageRequest struct {
	// The content of the message.
	Content *string `field:"required" json:"content" yaml:"content"`
	// The recipient of the message.
	To *string `field:"required" json:"to" yaml:"to"`
	// The image caption.
	Caption *string `field:"optional" json:"caption" yaml:"caption"`
	// The image data in base64 format.
	Data *string `field:"optional" json:"data" yaml:"data"`
	// The image URL.
	Url *string `field:"optional" json:"url" yaml:"url"`
}

