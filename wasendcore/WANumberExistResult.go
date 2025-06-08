package wasendcore


// WhatsApp number existence check result.
type WANumberExistResult struct {
	// Whether the number exists on WhatsApp.
	Exists *bool `field:"required" json:"exists" yaml:"exists"`
	// The phone number that was checked.
	Phone *string `field:"required" json:"phone" yaml:"phone"`
	// The JID of the contact if it exists.
	Jid *string `field:"optional" json:"jid" yaml:"jid"`
}

