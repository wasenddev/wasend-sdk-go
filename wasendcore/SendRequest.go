package wasendcore


// Send message request interface.
type SendRequest struct {
	// The session ID.
	Session *string `field:"required" json:"session" yaml:"session"`
	// The recipient's phone number or group JID.
	//
	// Example:
	//   "+1234567890" or "1234567890-12345678@g.us"
	//
	To *string `field:"required" json:"to" yaml:"to"`
	// The audio URL.
	//
	// Example:
	//   "https://example.com/audio.mp3"
	//
	AudioUrl *string `field:"optional" json:"audioUrl" yaml:"audioUrl"`
	// The document URL.
	//
	// Example:
	//   "https://example.com/document.pdf"
	//
	DocumentUrl *string `field:"optional" json:"documentUrl" yaml:"documentUrl"`
	// The filename for the media.
	//
	// Example:
	//   "custom_image.jpg"
	//
	Filename *string `field:"optional" json:"filename" yaml:"filename"`
	// The image URL.
	//
	// Example:
	//   "https://example.com/image.jpg"
	//
	ImageUrl *string `field:"optional" json:"imageUrl" yaml:"imageUrl"`
	// The MIME type of the media.
	//
	// Example:
	//   "image/jpeg"
	//
	Mimetype *string `field:"optional" json:"mimetype" yaml:"mimetype"`
	// The link preview information.
	Preview *map[string]*string `field:"optional" json:"preview" yaml:"preview"`
	// The message text or caption.
	//
	// Example:
	//   "Hello from WASend!"
	//
	Text *string `field:"optional" json:"text" yaml:"text"`
	// The video URL.
	//
	// Example:
	//   "https://example.com/video.mp4"
	//
	VideoUrl *string `field:"optional" json:"videoUrl" yaml:"videoUrl"`
}

