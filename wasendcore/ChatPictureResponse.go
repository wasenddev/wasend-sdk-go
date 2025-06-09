package wasendcore


type ChatPictureResponse struct {
	File *string `field:"optional" json:"file" yaml:"file"`
	Mimetype *string `field:"optional" json:"mimetype" yaml:"mimetype"`
	Url *string `field:"optional" json:"url" yaml:"url"`
}

