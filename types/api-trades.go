package types

// APITrade type represents trade we expose via maketrades API
type APITrade struct {
	Exchange      string `json:"exchange"`
	Type          string `json:"type"`
	Symbol        string `json:"symbol"`
	Received      int64  `json:"received_time"`
	TradeID       int64  `json:"trade_id"`
	EventTime     int64  `json:"event_time"`
	TradeTime     int64  `json:"trade_time"`
	MarketMaker   bool   `json:"market_maker"`
	SellerOrderID int64  `json:"seller_order_id"`
	BuyerOrderID  int64  `json:"buyer_order_id"`
	Price         string `json:"price"`
	Quantity      string `json:"quantity"`
	Side          int    `json:"side,omitempty"`
}

// APIFundingTrade type represents funding trade we expose via maketrades API
type APIFundingTrade struct {
	Exchange      string `json:"exchange"`
	Type          string `json:"type"`
	Symbol        string `json:"symbol"`
	Received      int64  `json:"received_time"`
	TradeID       int64  `json:"trade_id"`
	EventTime     int64  `json:"event_time"`
	TradeTime     int64  `json:"trade_time"`
	MarketMaker   bool   `json:"market_maker"`
	SellerOrderID int64  `json:"seller_order_id"`
	BuyerOrderID  int64  `json:"buyer_order_id"`
	Rate          string `json:"rate"`
	Period        int64  `json:"period"`
	Quantity      string `json:"quantity"`
	Side          int    `json:"side,omitempty"`
}
