package binance

import (
	"crypto/sha1"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"github.com/etrubenok/make-trades-types/types"
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

// GetTradeMessagePrimaryKey gets a concatenated string that forms from the message fields which identify the trade uniquely
func GetTradeMessagePrimaryKey(trade *TradeStreamMessageInKafka) string {
	fields := make([]string, 0)
	fields = append(fields, trade.Exchange)
	fields = append(fields, trade.Stream)
	fields = append(fields, strconv.FormatInt(trade.EventTime, 10))
	fields = append(fields, strconv.FormatInt(trade.RawMessage.Data.LowercaseT, 10))
	return strings.Join(fields, "@")
}

// GetDepthMessagePrimaryKey gets a concatenated string that forms from the message fields which identify the depth message uniquely
func GetDepthMessagePrimaryKey(trade *DepthStreamMessageInKafka) string {
	fields := make([]string, 0)
	fields = append(fields, trade.Exchange)
	fields = append(fields, trade.Stream)
	fields = append(fields, strconv.FormatInt(trade.EventTime, 10))
	fields = append(fields, strconv.FormatInt(trade.RawMessage.Data.U, 10))
	return strings.Join(fields, "@")
}

// GetKeyHash converts a plain string key into hash
func GetKeyHash(key string) (string, error) {
	h := sha1.New()
	_, err := h.Write([]byte(key))
	if err != nil {
		glog.Errorf("GetKeyHash: cannot make a hash for key %s due to error %s", key, err)
		return "", err
	}
	bs := h.Sum(nil)
	return fmt.Sprintf("%x", bs), nil
}

// KafkaBinanceTradeToAPITrade converts binance trade from Kafka message to APITrade format
func KafkaBinanceTradeToAPITrade(kafkaTrade *TradeStreamMessageInKafka) (*types.APITrade, error) {

	apiTrade := types.APITrade{
		Exchange:      kafkaTrade.Exchange,
		Symbol:        fmt.Sprintf("%s-%s", kafkaTrade.Exchange, kafkaTrade.Symbol),
		Received:      kafkaTrade.ReceivedTime,
		TradeID:       kafkaTrade.RawMessage.Data.LowercaseT,
		EventTime:     kafkaTrade.RawMessage.Data.E,
		TradeTime:     kafkaTrade.RawMessage.Data.T,
		MarketMaker:   kafkaTrade.RawMessage.Data.LowercaseM,
		SellerOrderID: kafkaTrade.RawMessage.Data.LowercaseA,
		BuyerOrderID:  kafkaTrade.RawMessage.Data.LowercaseB,
		Price:         kafkaTrade.RawMessage.Data.LowercaseP,
		Quantity:      kafkaTrade.RawMessage.Data.LowercaseQ}
	return &apiTrade, nil
}

// ConvertsBinancePriceLevelsIntoAPIOrderBookPriceLevels converts binance price level types into APIOrderBookPriceLevel
func ConvertsBinancePriceLevelsIntoAPIOrderBookPriceLevels(levels []PriceLevelType) ([]types.APIOrderBookPriceLevel, error) {
	apiLevels := make([]types.APIOrderBookPriceLevel, len(levels))
	for i, l := range levels {
		p, _ := l[0].(string)
		q, _ := l[1].(string)
		apiLevels[i] = types.APIOrderBookPriceLevel{
			Price:    p,
			Quantity: q}
	}
	return apiLevels, nil
}

// KafkaBinanceOrderBookUpdateToAPIOrderBookUpdate converts Binance depth stream message into APIOrderBookUpdate format
func KafkaBinanceOrderBookUpdateToAPIOrderBookUpdate(kafkaDepth *DepthStreamMessageInKafka) (*types.APIOrderBookUpdate, error) {

	asks, err := ConvertsBinancePriceLevelsIntoAPIOrderBookPriceLevels(kafkaDepth.RawMessage.Data.LowercaseA)
	if err != nil {
		glog.Errorf("KafkaBinanceOrderBookUpdateToAPIOrderBookUpdate: cannot get asks from msg %v due to error %s", kafkaDepth, err)
		return nil, err
	}
	bids, err := ConvertsBinancePriceLevelsIntoAPIOrderBookPriceLevels(kafkaDepth.RawMessage.Data.LowercaseB)
	if err != nil {
		glog.Errorf("KafkaBinanceOrderBookUpdateToAPIOrderBookUpdate: cannot get bids from msg %v due to error %s", kafkaDepth, err)
		return nil, err
	}

	apiOrderBookUpdate := types.APIOrderBookUpdate{
		Exchange:      kafkaDepth.Exchange,
		Symbol:        fmt.Sprintf("%s-%s", kafkaDepth.Exchange, kafkaDepth.Symbol),
		Received:      kafkaDepth.ReceivedTime,
		FirstUpdateID: kafkaDepth.RawMessage.Data.U,
		EventTime:     kafkaDepth.RawMessage.Data.E,
		LastUpdateID:  kafkaDepth.RawMessage.Data.LowercaseU,
		Asks:          asks,
		Bids:          bids,
	}
	return &apiOrderBookUpdate, nil
}
