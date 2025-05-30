package wasendcore


// Configuration information (without sensitive data).
type WasendConfigInfo struct {
	// The base URL for the API.
	BaseUrl *string `field:"optional" json:"baseUrl" yaml:"baseUrl"`
	// Request timeout in milliseconds.
	Timeout *float64 `field:"optional" json:"timeout" yaml:"timeout"`
}

