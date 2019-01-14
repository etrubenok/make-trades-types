package binance

import (
	"encoding/json"

	"github.com/golang/glog"
)

// GetBinanceTradeStreamMessage converts a message in the string format into TradeStreamMessageInKafka
func GetBinanceTradeStreamMessage(message []byte) (*TradeStreamMessageInKafka, error) {
	msg := TradeStreamMessageInKafka{}
	err := json.Unmarshal(message, &msg)
	if err != nil {
		glog.Errorf("GetBinanceTradeStreamMessage: cannot convert message %s into TradeStreamMessageInKafka due to error %s", message, err)
		return nil, err
	}
	return &msg, nil
}

// GetBinanceDepthStreamMessage converts a message in the string format into DepthStreamMessageInKafka
func GetBinanceDepthStreamMessage(message []byte) (*DepthStreamMessageInKafka, error) {
	msg := DepthStreamMessageInKafka{}
	err := json.Unmarshal(message, &msg)
	if err != nil {
		glog.Errorf("GetBinanceDepthStreamMessage: cannot convert message %s into DepthStreamMessageInKafka due to error %s", message, err)
		return nil, err
	}
	return &msg, nil
}
