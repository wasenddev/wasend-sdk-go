package wasendcore


// Profile picture request.
type ProfilePictureRequest struct {
	// Picture URL.
	Url *string `field:"required" json:"url" yaml:"url"`
	// Whether to crop the picture to a square.
	// Default: true.
	//
	CropToSquare *bool `field:"optional" json:"cropToSquare" yaml:"cropToSquare"`
	// Picture format (optional).
	// Default: "jpeg".
	//
	Format *string `field:"optional" json:"format" yaml:"format"`
}

