package wasendcore


// File message request.
type MessageFileRequest struct {
	// The content of the message.
	Content *string `field:"required" json:"content" yaml:"content"`
	// The recipient of the message.
	To *string `field:"required" json:"to" yaml:"to"`
	// The file name.
	FileName *string `field:"required" json:"fileName" yaml:"fileName"`
	// The file mime type.
	MimeType *string `field:"required" json:"mimeType" yaml:"mimeType"`
	// The file data in base64 format.
	Data *string `field:"optional" json:"data" yaml:"data"`
	// The file URL.
	Url *string `field:"optional" json:"url" yaml:"url"`
}

