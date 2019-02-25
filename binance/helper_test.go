package binance

import (
	"encoding/json"
	"io/ioutil"
	"testing"

	"github.com/etrubenok/make-trades-types/types"
	"github.com/stretchr/testify/assert"
)

func TestGetBinanceTradeStreamMessage(t *testing.T) {
	msgStr, err := ioutil.ReadFile("./test-data/trade-msg.json")
	assert.NoError(t, err)

	msg, err := GetBinanceTradeStreamMessage(msgStr)
	if assert.NoError(t, err) {
		assert.Equal(t, "binance", msg.Exchange)
		assert.Equal(t, "3738.35000000", msg.RawMessage.Data.LowercaseP)
	}
}

func TestGetBinanceDepthStreamMessage(t *testing.T) {
	msgStr, err := ioutil.ReadFile("./test-data/depth-update-msg.json")
	assert.NoError(t, err)
	msg, err := GetBinanceDepthStreamMessage(msgStr)
	if assert.NoError(t, err) {
		assert.Equal(t, "binance", msg.Exchange)
		assert.Equal(t, 5, len(msg.RawMessage.Data.LowercaseA))
		assert.Equal(t, "124.45000000", msg.RawMessage.Data.LowercaseA[0][0].(string))
		assert.Equal(t, "11.38000000", msg.RawMessage.Data.LowercaseA[0][1].(string))
		assert.Equal(t, 3, len(msg.RawMessage.Data.LowercaseB))
	}
}

func TestGetTradeMessagePrimaryKey(t *testing.T) {
	msgStr, err := ioutil.ReadFile("./test-data/trade-msg.json")
	assert.NoError(t, err)

	msg, err := GetBinanceTradeStreamMessage(msgStr)
	assert.NoError(t, err)

	key := GetTradeMessagePrimaryKey(msg)
	assert.Equal(t, "binance@btcusdt@trade@1547119839773@93684015", key)
}

func TestGetDepthMessagePrimaryKey(t *testing.T) {
	msgStr, err := ioutil.ReadFile("./test-data/depth-msg.json")
	assert.NoError(t, err)

	msg, err := GetBinanceDepthStreamMessage(msgStr)
	assert.NoError(t, err)

	key := GetDepthMessagePrimaryKey(msg)
	assert.Equal(t, "binance@ppteth@depth@1547119839095@49518600", key)
}

func TestGetHashKey(t *testing.T) {
	h, err := GetKeyHash("binance@btcusdt@trade@1547119839773@93684015")
	assert.NoError(t, err)

	assert.Equal(t, "9f875ee4585152c34a56c23a29df253104b60898", h)
}

func TestKafkaBinanceTradeToAPITrade(t *testing.T) {
	msgStr, err := ioutil.ReadFile("./test-data/binance-trade-in-kafka.json")
	assert.NoError(t, err)

	binanceTradeInKafka := TradeStreamMessageInKafka{}
	err = json.Unmarshal(msgStr, &binanceTradeInKafka)
	assert.NoError(t, err)

	expectedAPITrade := types.APITrade{
		Exchange:      "binance",
		Type:          "trade",
		Symbol:        "binance-bchsvusdt",
		Received:      1548817264969,
		TradeID:       4168042,
		EventTime:     1548817264887,
		TradeTime:     1548817264883,
		MarketMaker:   true,
		SellerOrderID: 18779843,
		BuyerOrderID:  18779842,
		Price:         "63.70000000",
		Quantity:      "0.16400000"}

	apiTrade, err := KafkaBinanceTradeToAPITrade(&binanceTradeInKafka)
	if assert.NoError(t, err) {
		assert.Equal(t, &expectedAPITrade, apiTrade)
	}
}

func TestKafkaBinanceOrderBookUpdateToAPIOrderBookUpdate(t *testing.T) {
	msgStr, err := ioutil.ReadFile("./test-data/depth-update-msg.json")
	assert.NoError(t, err)

	binanceDepthUpdateInKafka := DepthStreamMessageInKafka{}
	err = json.Unmarshal(msgStr, &binanceDepthUpdateInKafka)
	assert.NoError(t, err)

	expectedAPIOrderBookUpdate := types.APIOrderBookUpdate{
		Exchange:      "binance",
		Type:          "orderbook",
		Symbol:        "binance-ethusdt",
		Received:      1547291078868,
		FirstUpdateID: 268315354,
		EventTime:     1547291078782,
		LastUpdateID:  268315363,
		Asks:          []types.APIOrderBookPriceLevel{types.APIOrderBookPriceLevel{Price: "124.45000000", Quantity: "11.38000000"}, types.APIOrderBookPriceLevel{Price: "125.03000000", Quantity: "0.00000000"}, types.APIOrderBookPriceLevel{Price: "125.48000000", Quantity: "4.01880000"}, types.APIOrderBookPriceLevel{Price: "125.49000000", Quantity: "163.42146000"}, types.APIOrderBookPriceLevel{Price: "126.47000000", Quantity: "0.57238000"}},
		Bids:          []types.APIOrderBookPriceLevel{types.APIOrderBookPriceLevel{Price: "124.38000000", Quantity: "0.00000000"}, types.APIOrderBookPriceLevel{Price: "124.37000000", Quantity: "100.00000000"}, types.APIOrderBookPriceLevel{Price: "124.31000000", Quantity: "5.83937000"}}}

	apiOrdeBookUpdate, err := KafkaBinanceOrderBookUpdateToAPIOrderBookUpdate(&binanceDepthUpdateInKafka)
	if assert.NoError(t, err) {
		assert.Equal(t, &expectedAPIOrderBookUpdate, apiOrdeBookUpdate)
	}
}
