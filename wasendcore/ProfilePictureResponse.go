package wasendcore


// Profile picture response.
type ProfilePictureResponse struct {
	// The URL of the profile picture.
	Url *string `field:"required" json:"url" yaml:"url"`
	// The dimensions of the picture.
	//
	// Example:
	//   {
	//     "width": "800",
	//     "height": "600"
	//   }
	//
	Dimensions *map[string]*string `field:"optional" json:"dimensions" yaml:"dimensions"`
	// The format of the picture (e.g., 'jpeg', 'png').
	Format *string `field:"optional" json:"format" yaml:"format"`
	// The size of the picture in bytes.
	Size *float64 `field:"optional" json:"size" yaml:"size"`
}

