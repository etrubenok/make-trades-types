package types

import (
	"encoding/json"
	"io/ioutil"
	"testing"

	"github.com/etrubenok/make-trades-types/binance"
	"github.com/etrubenok/make-trades-types/types"
	"github.com/stretchr/testify/assert"
)

func TestGetStreamType(t *testing.T) {
	msgStr, err := ioutil.ReadFile("./binance/test-data/trade-msg.json")
	assert.NoError(t, err)

	msg, err := ConvertToJSON(string(msgStr))
	assert.NoError(t, err)

	tp, err := GetStreamType(&msg)
	assert.NoError(t, err)

	assert.Equal(t, "trade", tp)
}

func TestGetStreamTypeInvalid(t *testing.T) {
	msgStr, err := ioutil.ReadFile("./binance/test-data/trade-msg-invalid-type-1.json")
	assert.NoError(t, err)

	msg, err := ConvertToJSON(string(msgStr))
	assert.NoError(t, err)

	_, err = GetStreamType(&msg)
	if assert.Error(t, err) {
		assert.Equal(t, err.Error()[:43], "GetStreamType: cannot get 'stream' from msg")
	}
}

func TestGetExchange(t *testing.T) {
	msgStr, err := ioutil.ReadFile("./binance/test-data/trade-msg.json")
	assert.NoError(t, err)

	msg, err := ConvertToJSON(string(msgStr))
	assert.NoError(t, err)

	exchange, err := GetExchange(&msg)
	assert.NoError(t, err)

	assert.Equal(t, "binance", exchange)
}

func TestGetExchangeInvalid(t *testing.T) {
	msgStr, err := ioutil.ReadFile("./binance/test-data/trade-msg-exchange-invalid-format.json")
	assert.NoError(t, err)

	msg, err := ConvertToJSON(string(msgStr))
	assert.NoError(t, err)

	_, err = GetExchange(&msg)
	if assert.Error(t, err) {
		assert.Equal(t, err.Error()[:43], "GetExchange: cannot get 'exchange' from msg")
	}
}

func TestKafkaBinanceTradeToAPITrade(t *testing.T) {
	msgStr, err := ioutil.ReadFile("./binance/test-data/binance-trade-in-kafka.json")
	assert.NoError(t, err)

	binanceTradeInKafka := binance.TradeStreamMessageInKafka{}
	err = json.Unmarshal(msgStr, &binanceTradeInKafka)
	assert.NoError(t, err)

	expectedAPITrade := types.APITrade{
		Exchange:      "binance",
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
	msgStr, err := ioutil.ReadFile("./binance/test-data/depth-update-msg.json")
	assert.NoError(t, err)

	binanceDepthUpdateInKafka := binance.DepthStreamMessageInKafka{}
	err = json.Unmarshal(msgStr, &binanceDepthUpdateInKafka)
	assert.NoError(t, err)

	expectedAPIOrderBookUpdate := types.APIOrderBookUpdate{
		Exchange:      "binance",
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

func TestRandStringBytes(t *testing.T) {
	str := RandStringBytes(10)
	assert.Equal(t, 10, len(str))
}
