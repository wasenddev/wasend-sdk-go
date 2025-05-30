package wasendcore


type GowsStatus struct {
	Connected *bool `field:"required" json:"connected" yaml:"connected"`
	Found *bool `field:"required" json:"found" yaml:"found"`
}

