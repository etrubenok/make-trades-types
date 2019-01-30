package types

// APISymbolInfo type contains information about one symbol
type APISymbolInfo struct {
	Symbol string `json:"symbol"`
	Status string `json:"status"`
	Asset  string `json:"asset"`
	Quote  string `json:"quote"`
}

// APIExchangeSymbols type contains information about symbols of an exchange
type APIExchangeSymbols struct {
	Exchange     string          `json:"exchange"`
	SnapshotTime int64           `json:"snapshot_time"`
	Symbols      []APISymbolInfo `json:"symbols"`
}

// APIExchangesSymbols type contains information about symbols of several exchanges
type APIExchangesSymbols struct {
	Exchanges []APIExchangeSymbols `json:"exchanges"`
}
