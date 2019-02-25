package types

// APIOrderBookPriceLevel type repsents price level in order book in API
type APIOrderBookPriceLevel struct {
	Price    string `json:"price"`
	Quantity string `json:"quantity"`
}

// APIOrderBookUpdate type represents API messages for order book updates
type APIOrderBookUpdate struct {
	Exchange      string                   `json:"exchange"`
	Type          string                   `json:"type"`
	Symbol        string                   `json:"symbol"`
	Received      int64                    `json:"received_time"`
	FirstUpdateID int64                    `json:"first_update_id"`
	EventTime     int64                    `json:"event_time"`
	LastUpdateID  int64                    `json:"last_update_id"`
	Asks          []APIOrderBookPriceLevel `json:"asks"`
	Bids          []APIOrderBookPriceLevel `json:"bids"`
}
