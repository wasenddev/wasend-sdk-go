package wasendcore


// Link preview request.
type LinkPreviewRequest struct {
	// The link title.
	Title *string `field:"required" json:"title" yaml:"title"`
	// The link description.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The link thumbnail URL.
	ThumbnailUrl *string `field:"optional" json:"thumbnailUrl" yaml:"thumbnailUrl"`
}

