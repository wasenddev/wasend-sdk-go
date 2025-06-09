package wasendcore


// Request to react to a message.
type MessageReactionRequest struct {
	// The ID of the message to react to.
	//
	// Example:
	//   "false_11111111111@c.us_AAAAAAAAAAAAAAAAAAAA"
	//
	MessageId *string `field:"required" json:"messageId" yaml:"messageId"`
	// Emoji to react with.
	//
	// Send an empty string to remove the reaction.
	//
	// Example:
	//   "👍"
	//
	Reaction *string `field:"required" json:"reaction" yaml:"reaction"`
	// The session ID.
	// Default: "default".
	//
	Session *string `field:"required" json:"session" yaml:"session"`
}

