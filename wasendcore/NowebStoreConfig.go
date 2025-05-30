package wasendcore


type NowebStoreConfig struct {
	Enabled *bool `field:"required" json:"enabled" yaml:"enabled"`
	FullSync *bool `field:"required" json:"fullSync" yaml:"fullSync"`
}

