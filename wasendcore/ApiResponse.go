package wasendcore


// Response from API calls.
type ApiResponse struct {
	// Whether the request was successful.
	Success *bool `field:"required" json:"success" yaml:"success"`
	// The response data (if successful).
	Data interface{} `field:"optional" json:"data" yaml:"data"`
	// Error message if the request failed.
	Error *string `field:"optional" json:"error" yaml:"error"`
}

