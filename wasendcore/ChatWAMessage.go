package wasendcore


type ChatWAMessage struct {
	ChatId *string `field:"required" json:"chatId" yaml:"chatId"`
	From *string `field:"required" json:"from" yaml:"from"`
	FromMe *bool `field:"required" json:"fromMe" yaml:"fromMe"`
	Id *string `field:"required" json:"id" yaml:"id"`
	Timestamp *float64 `field:"required" json:"timestamp" yaml:"timestamp"`
	To *string `field:"required" json:"to" yaml:"to"`
	Type *string `field:"required" json:"type" yaml:"type"`
	Ack *float64 `field:"optional" json:"ack" yaml:"ack"`
	Body interface{} `field:"optional" json:"body" yaml:"body"`
	Caption *string `field:"optional" json:"caption" yaml:"caption"`
	Filename *string `field:"optional" json:"filename" yaml:"filename"`
	HasMedia *bool `field:"optional" json:"hasMedia" yaml:"hasMedia"`
	MediaUrl *string `field:"optional" json:"mediaUrl" yaml:"mediaUrl"`
	Mimetype *string `field:"optional" json:"mimetype" yaml:"mimetype"`
	QuotedMsgId *string `field:"optional" json:"quotedMsgId" yaml:"quotedMsgId"`
}

