package types

import (
	"bytes"
	"encoding/json"
	"fmt"
	"math/rand"
	"strings"

	"github.com/etrubenok/make-trades-types/binance"
	"github.com/etrubenok/make-trades-types/types"
	"github.com/golang/glog"
)

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

// RandStringBytes generates a random string with a given length
func RandStringBytes(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

// ConvertToJSON converts string to JSON struct
func ConvertToJSON(message string) (map[string]interface{}, error) {
	m := make(map[string]interface{})

	d := json.NewDecoder(bytes.NewBuffer([]byte(message)))
	d.UseNumber()
	if err := d.Decode(&m); err != nil {
		e := fmt.Errorf("ConvertToJSON: cannot unmarshal message %s to JSON due to error %s", message, err)
		glog.Error(e)
		return m, e
	}

	return m, nil
}

// GetStreamType returns type of the message
func GetStreamType(msg *map[string]interface{}) (string, error) {
	s, ok := (*msg)["stream"].(string)
	if !ok {
		return "", fmt.Errorf("GetStreamType: cannot get 'stream' from msg %v", msg)
	}
	tokens := strings.Split(s, "@")
	if len(tokens) != 2 {
		return "", fmt.Errorf("GetStreamType: the format of the 'stream': %s is not what we expect (it should have two tokens)", s)
	}
	return tokens[1], nil
}

// GetExchange return exchange name from the message
func GetExchange(msg *map[string]interface{}) (string, error) {
	e, ok := (*msg)["exchange"].(string)
	if !ok {
		return "", fmt.Errorf("GetExchange: cannot get 'exchange' from msg %v", msg)
	}
	return e, nil
}

// KafkaBinanceTradeToAPITrade converts binance trade from Kafka message to APITrade format
func KafkaBinanceTradeToAPITrade(kafkaTrade *binance.TradeStreamMessageInKafka) (*types.APITrade, error) {

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
func ConvertsBinancePriceLevelsIntoAPIOrderBookPriceLevels(levels []binance.PriceLevelType) ([]types.APIOrderBookPriceLevel, error) {
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
func KafkaBinanceOrderBookUpdateToAPIOrderBookUpdate(kafkaDepth *binance.DepthStreamMessageInKafka) (*types.APIOrderBookUpdate, error) {

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
