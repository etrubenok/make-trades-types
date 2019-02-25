package types

// examples:
// Binance Trade:
// {"exchange":"binance","stream":"trxbtc@trade","symbol":"trxbtc","event_time":1551052690833,"received_time":1551052690919,"raw_message":{"data":{"E":1551052690833,"M":true,"T":1551052690825,"a":95055990,"b":95053089,"e":"trade","m":true,"p":"0.00000625","q":"169.00000000","s":"TRXBTC","t":39604013},"stream":"trxbtc@trade"}}
//
// Binance Order Book Update:
// {"exchange":"binance","stream":"knceth@depth","symbol":"knceth","event_time":1551052690164,"received_time":1551052690343,"raw_message":{"data":{"E":1551052690164,"U":52434492,"a":[["0.00106550","880.00000000",[]],["0.00106560","1581.00000000",[]]],"b":[["0.00105760","0.00000000",[]],["0.00105740","0.00000000",[]],["0.00105700","0.00000000",[]],["0.00104360","3223.00000000",[]],["0.00104250","2441.00000000",[]]],"e":"depthUpdate","s":"KNCETH","u":52434500},"stream":"knceth@depth"}}
//

// ExchangeMessageMeta contains information extracted from the exchange message
type ExchangeMessageMeta struct {
	Exchange     string                 `json:"exchange"`
	Stream       string                 `json:"stream"`
	Symbol       string                 `json:"symbol"`
	EventTime    int64                  `json:"event_time"`
	ReceivedTime int64                  `json:"received_time"`
	RawMessage   map[string]interface{} `json:"raw_message"`
}
