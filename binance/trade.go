package binance

// RawTradeDataType is a type of data in the Binance stream trade messages
type RawTradeDataType struct {
	E          int64  `json:"E"`
	M          bool   `json:"M"`
	T          int64  `json:"T"`
	LowercaseA int64  `json:"a"`
	LowercaseB int64  `json:"b"`
	LowercaseE string `json:"e"`
	LowercaseM bool   `json:"m"`
	LowercaseP string `json:"p"`
	LowercaseQ string `json:"q"`
	LowercaseS string `json:"s"`
	LowercaseT int64  `json:"t"`
}

// RawMessageTradeType is a type of Binance stream trade messages
type RawMessageTradeType struct {
	Data   RawTradeDataType `json:"data"`
	Stream string           `json:"stream"`
}

// TradeStreamMessageInKafka is a type of Binance stream trade messages
type TradeStreamMessageInKafka struct {
	Exchange     string              `json:"exchange"`
	Stream       string              `json:"stream"`
	Symbol       string              `json:"symbol"`
	EventTime    int64               `json:"event_time"`
	ReceivedTime int64               `json:"received_time"`
	RawMessage   RawMessageTradeType `json:"raw_message"`
}
