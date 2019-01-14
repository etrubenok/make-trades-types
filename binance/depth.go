package binance

// PriceLevelType is a type for Binance order book price level
type PriceLevelType []interface{}

// RawDepthDataType is a type for Binance depth message data
type RawDepthDataType struct {
	E          int64            `json:"E"`
	LowercaseE string           `json:"e"`
	U          int64            `json:"U"`
	LowercaseU int64            `json:"u"`
	LowercaseS string           `json:"s"`
	LowercaseA []PriceLevelType `json:"a"`
	LowercaseB []PriceLevelType `json:"b"`
}

// RawMessageDepthType is a type for Binance depth message
type RawMessageDepthType struct {
	Data   RawDepthDataType `json:"data"`
	Stream string           `json:"stream"`
}

// DepthStreamMessageInKafka is a type of Binance stream depth messages
type DepthStreamMessageInKafka struct {
	Exchange     string              `json:"exchange"`
	Stream       string              `json:"stream"`
	Symbol       string              `json:"symbol"`
	EventTime    int64               `json:"event_time"`
	ReceivedTime int64               `json:"received_time"`
	RawMessage   RawMessageDepthType `json:"raw_message"`
}
