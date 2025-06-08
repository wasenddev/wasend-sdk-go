package wasendcore


// Video message request.
type MessageVideoRequest struct {
	// The content of the message.
	Content *string `field:"required" json:"content" yaml:"content"`
	// The recipient of the message.
	To *string `field:"required" json:"to" yaml:"to"`
	// The video caption.
	Caption *string `field:"optional" json:"caption" yaml:"caption"`
	// The video data in base64 format.
	Data *string `field:"optional" json:"data" yaml:"data"`
	// The video URL.
	Url *string `field:"optional" json:"url" yaml:"url"`
}

