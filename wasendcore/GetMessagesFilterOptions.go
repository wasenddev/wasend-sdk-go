package wasendcore


type GetMessagesFilterOptions struct {
	Ack interface{} `field:"optional" json:"ack" yaml:"ack"`
	FromMe *bool `field:"optional" json:"fromMe" yaml:"fromMe"`
	TimestampGte *float64 `field:"optional" json:"timestampGte" yaml:"timestampGte"`
	TimestampLte *float64 `field:"optional" json:"timestampLte" yaml:"timestampLte"`
}

