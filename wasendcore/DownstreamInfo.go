package wasendcore


// Downstream connection information.
type DownstreamInfo struct {
	// Configuration for the downstream connection.
	Config *SessionConfig `field:"required" json:"config" yaml:"config"`
	// Engine status information.
	Engine *EngineStatus `field:"required" json:"engine" yaml:"engine"`
	// Name of the downstream connection.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Status of the downstream connection.
	Status *string `field:"required" json:"status" yaml:"status"`
}

