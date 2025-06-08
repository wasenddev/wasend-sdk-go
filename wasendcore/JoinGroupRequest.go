package wasendcore


// Join group request.
type JoinGroupRequest struct {
	// Group code or invite URL.
	//
	// Example:
	//   "https://chat.whatsapp.com/1234567890abcdef"
	//
	Code *string `field:"required" json:"code" yaml:"code"`
}

