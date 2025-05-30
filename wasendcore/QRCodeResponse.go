package wasendcore


// QR code response.
type QRCodeResponse struct {
	// QR code data.
	Data *string `field:"required" json:"data" yaml:"data"`
}

