package wasendcore


type SessionMetadata struct {
	UserEmail *string `field:"required" json:"userEmail" yaml:"userEmail"`
	UserId *string `field:"required" json:"userId" yaml:"userId"`
}

