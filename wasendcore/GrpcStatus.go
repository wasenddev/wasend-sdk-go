package wasendcore


type GrpcStatus struct {
	Client *string `field:"required" json:"client" yaml:"client"`
	Stream *string `field:"required" json:"stream" yaml:"stream"`
}

