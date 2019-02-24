package registry

import "fmt"

// GetExchangeID returns the id for the given exchange name
func GetExchangeID(exchange string) (int, error) {
	if exchange == "binance" {
		return 1, nil
	} else if exchange == "bitfinex" {
		return 2, nil
	}
	return 0, fmt.Errorf("GetExchangeID: exchange: '%s' is not known", exchange)
}

// GetExchangeNameByID converts exchangeID into exchange name
func GetExchangeNameByID(exchangeID int) (string, error) {
	switch exchangeID {
	case 1:
		return "binance", nil
	case 2:
		return "bitfinex", nil
	default:
		return "", fmt.Errorf("GetExchangeNameByID does not know exchange with ID: %d", exchangeID)
	}
}
