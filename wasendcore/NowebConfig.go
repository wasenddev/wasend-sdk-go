package wasendcore


type NowebConfig struct {
	MarkOnline *bool `field:"required" json:"markOnline" yaml:"markOnline"`
	Store *NowebStoreConfig `field:"required" json:"store" yaml:"store"`
}

