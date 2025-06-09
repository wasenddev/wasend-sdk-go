package wasendcore


type ReadChatMessagesResponse struct {
	Success *bool `field:"required" json:"success" yaml:"success"`
	Message *string `field:"optional" json:"message" yaml:"message"`
}

