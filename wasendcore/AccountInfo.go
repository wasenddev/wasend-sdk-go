package wasendcore


// Account information structure.
type AccountInfo struct {
	// Account ID.
	Id *string `field:"required" json:"id" yaml:"id"`
	// Account name.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Account plan.
	Plan *string `field:"required" json:"plan" yaml:"plan"`
}

