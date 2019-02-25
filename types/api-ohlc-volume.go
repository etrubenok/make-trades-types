package types

// {"start":1550834040000,"finish":1550834100000,"pair":"binance-eosusdt","count":69,"volume":566351999997,"open":381110000,"high":381110000,"low":380480000,"close":380480000}

// APIOHLCVolume type represents API messages for order book updates
type APIOHLCVolume struct {
	Start  int64  `json:"start"`
	Finish int64  `json:"finish"`
	Pair   string `json:"pair"`
	Count  int64  `json:"count"`
	Volume int64  `json:"volume"`
	Open   int64  `json:"open"`
	High   int64  `json:"high"`
	Low    int64  `json:"low"`
	Close  int64  `json:"close"`
}
