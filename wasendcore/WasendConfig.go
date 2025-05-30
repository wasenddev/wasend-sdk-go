package wasendcore


// Configuration options for the Wasend SDK.
type WasendConfig struct {
	// The API key for authentication.
	ApiKey *string `field:"required" json:"apiKey" yaml:"apiKey"`
	// The base URL for the API (optional).
	// Default: https://api.wasend.dev
	//
	BaseUrl *string `field:"optional" json:"baseUrl" yaml:"baseUrl"`
	// Request timeout in milliseconds.
	// Default: 30000.
	//
	Timeout *float64 `field:"optional" json:"timeout" yaml:"timeout"`
}

