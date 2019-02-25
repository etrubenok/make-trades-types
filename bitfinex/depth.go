package bitfinex

import bitfinex "github.com/bitfinexcom/bitfinex-api-go/v2"

// DepthStreamMessageInKafka is a type of Binance stream depth messages
type DepthStreamMessageInKafka struct {
	Exchange     string              `json:"exchange"`
	Stream       string              `json:"stream"`
	Symbol       string              `json:"symbol"`
	EventTime    int64               `json:"event_time"`
	ReceivedTime int64               `json:"received_time"`
	RawMessage   bitfinex.BookUpdate `json:"raw_message"`
}
