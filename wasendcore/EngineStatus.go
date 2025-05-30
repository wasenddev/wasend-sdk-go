package wasendcore


// Engine status information.
type EngineStatus struct {
	// GoWS status.
	Gows *GowsStatus `field:"required" json:"gows" yaml:"gows"`
	// gRPC status.
	Grpc *GrpcStatus `field:"required" json:"grpc" yaml:"grpc"`
}

