package bitfinex

import bitfinex "github.com/bitfinexcom/bitfinex-api-go/v2"

// RawMessageTradeType is a type for Bitfinex raw mnessage
type RawMessageTradeType struct {
	ID     string  `json:"ID"`
	MTS    int64   `json:"MTS"`
	Amount float64 `json:"Amount"`
	Price  float64 `json:"Price"`
	Side   int     `json:"Side"`
	Pair   string  `json:"Pair"`
}

// TradeStreamMessageInKafka is a type of Binance stream trade messages
type TradeStreamMessageInKafka struct {
	Exchange     string         `json:"exchange"`
	Stream       string         `json:"stream"`
	Symbol       string         `json:"symbol"`
	EventTime    int64          `json:"event_time"`
	ReceivedTime int64          `json:"received_time"`
	RawMessage   bitfinex.Trade `json:"raw_message"`
}
