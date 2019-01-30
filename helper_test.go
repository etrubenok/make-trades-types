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
