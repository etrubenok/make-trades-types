package bitfinex

import (
	"crypto/sha1"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	bitfinex "github.com/etrubenok/bitfinex-api-go/v2"
	"github.com/etrubenok/make-trades-types/types"
	"github.com/golang/glog"
)

// GetTradeStreamMessage converts a message in the string format into TradeStreamMessageInKafka
func GetTradeStreamMessage(message []byte) (*TradeStreamMessageInKafka, error) {
	msg := TradeStreamMessageInKafka{}
	err := json.Unmarshal(message, &msg)
	if err != nil {
		glog.Errorf("bitfinex.GetTradeStreamMessage: cannot convert message %s into TradeStreamMessageInKafka due to error %s", message, err)
		return nil, err
	}
	return &msg, nil
}

// GetFundingTradeStreamMessage converts a message in the string format into TradeStreamMessageInKafka
func GetFundingTradeStreamMessage(message []byte) (*FundingTradeStreamMessageInKafka, error) {
	msg := FundingTradeStreamMessageInKafka{}
	err := json.Unmarshal(message, &msg)
	if err != nil {
		glog.Errorf("bitfinex.GetFundingTradeStreamMessage: cannot convert message %s into FundingTradeStreamMessageInKafka due to error %s", message, err)
		return nil, err
	}
	return &msg, nil
}

// GetDepthStreamMessage converts a message in the string format into DepthStreamMessageInKafka
func GetDepthStreamMessage(message []byte) (*DepthStreamMessageInKafka, error) {
	msg := DepthStreamMessageInKafka{}
	err := json.Unmarshal(message, &msg)
	if err != nil {
		glog.Errorf("GetDepthStreamMessage: cannot convert message %s into DepthStreamMessageInKafka due to error %s", message, err)
		return nil, err
	}
	return &msg, nil
}

// GetFundingDepthStreamMessage converts a message in the string format into FundingDepthStreamMessageInKafka
func GetFundingDepthStreamMessage(message []byte) (*FundingDepthStreamMessageInKafka, error) {
	msg := FundingDepthStreamMessageInKafka{}
	err := json.Unmarshal(message, &msg)
	if err != nil {
		glog.Errorf("GetFundingDepthStreamMessage: cannot convert message %s into FundingDepthStreamMessageInKafka due to error %s", message, err)
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
	fields = append(fields, strconv.FormatInt(trade.RawMessage.ID, 10))
	return strings.Join(fields, "@")
}

// GetFundingTradeMessagePrimaryKey gets a concatenated string that forms from the message fields which identify the trade uniquely
func GetFundingTradeMessagePrimaryKey(trade *FundingTradeStreamMessageInKafka) string {
	fields := make([]string, 0)
	fields = append(fields, trade.Exchange)
	fields = append(fields, trade.Stream)
	fields = append(fields, strconv.FormatInt(trade.RawMessage.ID, 10))
	return strings.Join(fields, "@")
}

// GetDepthMessagePrimaryKey gets a concatenated string that forms from the message fields which identify the depth message uniquely
func GetDepthMessagePrimaryKey(depth *DepthStreamMessageInKafka) string {
	fields := make([]string, 0)
	fields = append(fields, depth.Exchange)
	fields = append(fields, depth.Stream)
	fields = append(fields, strconv.FormatInt(depth.RawMessage.ID, 10))
	fields = append(fields, fmt.Sprintf("%.8f", depth.RawMessage.Amount))
	fields = append(fields, strconv.FormatInt(depth.RawMessage.Count, 10))
	fields = append(fields, fmt.Sprintf("%.8f", depth.RawMessage.Price))
	fields = append(fields, strconv.FormatInt(int64(depth.RawMessage.Side), 10))
	return strings.Join(fields, "@")
}

// GetFundingDepthMessagePrimaryKey gets a concatenated string that forms from the message fields which identify the depth message uniquely
func GetFundingDepthMessagePrimaryKey(depth *FundingDepthStreamMessageInKafka) string {
	fields := make([]string, 0)
	fields = append(fields, depth.Exchange)
	fields = append(fields, depth.Stream)
	fields = append(fields, strconv.FormatInt(depth.RawMessage.ID, 10))
	fields = append(fields, fmt.Sprintf("%.8f", depth.RawMessage.Amount))
	fields = append(fields, strconv.FormatInt(depth.RawMessage.Count, 10))
	fields = append(fields, fmt.Sprintf("%.8f", depth.RawMessage.Rate))
	fields = append(fields, strconv.FormatInt(int64(depth.RawMessage.Side), 10))
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

// KafkaBitfinexOrderBookUpdateToAPIOrderBookUpdate converts Bitfinex depth stream message into APIOrderBookUpdate format
func KafkaBitfinexOrderBookUpdateToAPIOrderBookUpdate(kafkaDepth *DepthStreamMessageInKafka) (*types.APIOrderBookUpdate, error) {
	bids := make([]types.APIOrderBookPriceLevel, 0)
	asks := make([]types.APIOrderBookPriceLevel, 0)
	if kafkaDepth.RawMessage.Side == bitfinex.Bid {
		if kafkaDepth.RawMessage.Count == 0 {
			bids = append(bids, types.APIOrderBookPriceLevel{
				Price:    fmt.Sprintf("%.8f", kafkaDepth.RawMessage.Price),
				Quantity: fmt.Sprintf("%.8f", 0.0)})
		} else {
			bids = append(bids, types.APIOrderBookPriceLevel{
				Price:    fmt.Sprintf("%.8f", kafkaDepth.RawMessage.Price),
				Quantity: fmt.Sprintf("%.8f", kafkaDepth.RawMessage.Amount)})
		}
	} else if kafkaDepth.RawMessage.Side == bitfinex.Ask {
		if kafkaDepth.RawMessage.Count == 0 {
			asks = append(asks, types.APIOrderBookPriceLevel{
				Price:    fmt.Sprintf("%.8f", kafkaDepth.RawMessage.Price),
				Quantity: fmt.Sprintf("%.8f", 0.0)})
		} else {
			asks = append(asks, types.APIOrderBookPriceLevel{
				Price:    fmt.Sprintf("%.8f", kafkaDepth.RawMessage.Price),
				Quantity: fmt.Sprintf("%.8f", kafkaDepth.RawMessage.Amount)})
		}
	}

	apiOrderBookUpdate := types.APIOrderBookUpdate{
		Exchange:      kafkaDepth.Exchange,
		Type:          "orderbook",
		Symbol:        fmt.Sprintf("%s-%s", kafkaDepth.Exchange, kafkaDepth.Symbol),
		Received:      kafkaDepth.ReceivedTime,
		FirstUpdateID: kafkaDepth.RawMessage.ID,
		EventTime:     kafkaDepth.EventTime,
		LastUpdateID:  kafkaDepth.RawMessage.ID,
		Asks:          asks,
		Bids:          bids,
	}
	return &apiOrderBookUpdate, nil
}

// KafkaBitfinexTradeToAPITrade converts binance trade from Kafka message to APITrade format
func KafkaBitfinexTradeToAPITrade(kafkaTrade *TradeStreamMessageInKafka) (*types.APITrade, error) {

	side := 0
	if kafkaTrade.RawMessage.Side == bitfinex.Bid {
		side = 1
	} else if kafkaTrade.RawMessage.Side == bitfinex.Ask {
		side = -1
	}
	apiTrade := types.APITrade{
		Exchange:      kafkaTrade.Exchange,
		Type:          "trade",
		Symbol:        fmt.Sprintf("%s-%s", kafkaTrade.Exchange, kafkaTrade.Symbol),
		Received:      kafkaTrade.ReceivedTime,
		TradeID:       kafkaTrade.RawMessage.ID,
		EventTime:     kafkaTrade.RawMessage.MTS,
		TradeTime:     kafkaTrade.RawMessage.MTS,
		MarketMaker:   false,
		SellerOrderID: -1,
		BuyerOrderID:  -1,
		Price:         fmt.Sprintf("%.8f", kafkaTrade.RawMessage.Price),
		Quantity:      fmt.Sprintf("%.8f", kafkaTrade.RawMessage.Amount),
		Side:          side}
	return &apiTrade, nil
}
