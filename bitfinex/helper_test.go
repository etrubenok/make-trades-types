package bitfinex

import (
	"encoding/json"
	"io/ioutil"
	"testing"

	bitfinexOrig "github.com/bitfinexcom/bitfinex-api-go/v2"
	"github.com/etrubenok/make-trades-types/types"
	"github.com/stretchr/testify/assert"
)

func TestGetTradeStreamMessage(t *testing.T) {
	msgStr, err := ioutil.ReadFile("./test-data/trade-msg.json")
	assert.NoError(t, err)

	msg, err := GetTradeStreamMessage(msgStr)
	if assert.NoError(t, err) {
		assert.Equal(t, "bitfinex", msg.Exchange)
		assert.Equal(t, 26.57, msg.RawMessage.Amount)
		assert.Equal(t, 3.6751, msg.RawMessage.Price)
		assert.Equal(t, bitfinexOrig.Ask, msg.RawMessage.Side)
	}
}

func TestGetDepthStreamMessage(t *testing.T) {
	msgStr, err := ioutil.ReadFile("./test-data/depth-msg.json")
	assert.NoError(t, err)
	msg, err := GetDepthStreamMessage(msgStr)
	if assert.NoError(t, err) {
		assert.Equal(t, "bitfinex", msg.Exchange)
		assert.Equal(t, 1.0, msg.RawMessage.Amount)
		assert.Equal(t, 3.7475, msg.RawMessage.Price)
		assert.Equal(t, json.Number("-1"), msg.RawMessage.AmountJsNum)
		assert.Equal(t, json.Number("3.7475"), msg.RawMessage.PriceJsNum)
		assert.Equal(t, bitfinexOrig.BookAction(0), msg.RawMessage.Action)
		assert.Equal(t, bitfinexOrig.OrderSide(2), msg.RawMessage.Side)
	}
}

func TestKafkaBitfinexOrderBookUpdateToAPIOrderBook(t *testing.T) {
	msgStr, err := ioutil.ReadFile("./test-data/depth-msg.json")
	assert.NoError(t, err)

	bitfinexDepthUpdateInKafka := DepthStreamMessageInKafka{}
	err = json.Unmarshal(msgStr, &bitfinexDepthUpdateInKafka)
	assert.NoError(t, err)

	expectedAPIOrderBookUpdate := types.APIOrderBookUpdate{
		Exchange:      "bitfinex",
		Type:          "orderbook",
		Symbol:        "bitfinex-tEOSUSD",
		Received:      1551070004436,
		FirstUpdateID: 0,
		EventTime:     1551070004436,
		LastUpdateID:  0,
		Asks:          []types.APIOrderBookPriceLevel{types.APIOrderBookPriceLevel{Price: "3.74750000", Quantity: "0.00000000"}},
		Bids:          []types.APIOrderBookPriceLevel{}}

	apiOrdeBookUpdate, err := KafkaBitfinexOrderBookUpdateToAPIOrderBookUpdate(&bitfinexDepthUpdateInKafka)
	if assert.NoError(t, err) {
		assert.Equal(t, &expectedAPIOrderBookUpdate, apiOrdeBookUpdate)
	}
}

func TestKafkaBitfinexOrderBookUpdateToAPIOrderBookUpdate(t *testing.T) {
	msgStr, err := ioutil.ReadFile("./test-data/depth-msg-update.json")
	assert.NoError(t, err)

	bitfinexDepthUpdateInKafka := DepthStreamMessageInKafka{}
	err = json.Unmarshal(msgStr, &bitfinexDepthUpdateInKafka)
	assert.NoError(t, err)

	expectedAPIOrderBookUpdate := types.APIOrderBookUpdate{
		Exchange:      "bitfinex",
		Type:          "orderbook",
		Symbol:        "bitfinex-tEOSUSD",
		Received:      1551096530408,
		FirstUpdateID: 0,
		EventTime:     1551096530408,
		LastUpdateID:  0,
		Asks:          []types.APIOrderBookPriceLevel{},
		Bids:          []types.APIOrderBookPriceLevel{types.APIOrderBookPriceLevel{Price: "3.45310000", Quantity: "1300.00000000"}}}

	apiOrdeBookUpdate, err := KafkaBitfinexOrderBookUpdateToAPIOrderBookUpdate(&bitfinexDepthUpdateInKafka)
	if assert.NoError(t, err) {
		assert.Equal(t, &expectedAPIOrderBookUpdate, apiOrdeBookUpdate)
	}
}

func TestKafkaBitfinexOrderBookUpdateToAPIOrderBookUpdateExp(t *testing.T) {
	msgStr, err := ioutil.ReadFile("./test-data/depth-msg-price-exp.json")
	assert.NoError(t, err)

	bitfinexDepthUpdateInKafka := DepthStreamMessageInKafka{}
	err = json.Unmarshal(msgStr, &bitfinexDepthUpdateInKafka)
	assert.NoError(t, err)

	expectedAPIOrderBookUpdate := types.APIOrderBookUpdate{
		Exchange:      "bitfinex",
		Type:          "orderbook",
		Symbol:        "bitfinex-tBATBTC",
		Received:      1551267404491,
		FirstUpdateID: 0,
		EventTime:     1551267404491,
		LastUpdateID:  0,
		Asks:          []types.APIOrderBookPriceLevel{},
		Bids:          []types.APIOrderBookPriceLevel{types.APIOrderBookPriceLevel{Price: "0.00000010", Quantity: "51200.00000000"}}}

	apiOrdeBookUpdate, err := KafkaBitfinexOrderBookUpdateToAPIOrderBookUpdate(&bitfinexDepthUpdateInKafka)
	if assert.NoError(t, err) {
		assert.Equal(t, &expectedAPIOrderBookUpdate, apiOrdeBookUpdate)
	}
}

func TestKafkaBitfinexOrderBookUpdateToAPIOrderBookUpdateAsk(t *testing.T) {
	msgStr, err := ioutil.ReadFile("./test-data/depth-msg-ask.json")
	assert.NoError(t, err)

	bitfinexDepthUpdateInKafka := DepthStreamMessageInKafka{}
	err = json.Unmarshal(msgStr, &bitfinexDepthUpdateInKafka)
	assert.NoError(t, err)

	expectedAPIOrderBookUpdate := types.APIOrderBookUpdate{
		Exchange:      "bitfinex",
		Type:          "orderbook",
		Symbol:        "bitfinex-tZECBTC",
		Received:      1551267598182,
		FirstUpdateID: 0,
		EventTime:     1551267598182,
		LastUpdateID:  0,
		Asks:          []types.APIOrderBookPriceLevel{types.APIOrderBookPriceLevel{Price: "0.01364400", Quantity: "3.99030282"}},
		Bids:          []types.APIOrderBookPriceLevel{}}

	apiOrdeBookUpdate, err := KafkaBitfinexOrderBookUpdateToAPIOrderBookUpdate(&bitfinexDepthUpdateInKafka)
	if assert.NoError(t, err) {
		assert.Equal(t, &expectedAPIOrderBookUpdate, apiOrdeBookUpdate)
	}
}
func TestKafkaBitfinexTradeToAPITrade(t *testing.T) {
	msgStr, err := ioutil.ReadFile("./test-data/trade-msg.json")
	assert.NoError(t, err)

	bitfinexTradeInKafka := TradeStreamMessageInKafka{}
	err = json.Unmarshal(msgStr, &bitfinexTradeInKafka)
	assert.NoError(t, err)

	expectedAPITrade := types.APITrade{
		Exchange:      "bitfinex",
		Type:          "trade",
		Symbol:        "bitfinex-tEOSUSD",
		Received:      1547119839859,
		TradeID:       340528434,
		EventTime:     1551071701829,
		TradeTime:     1551071701829,
		MarketMaker:   false,
		SellerOrderID: -1,
		BuyerOrderID:  -1,
		Price:         "3.67510000",
		Quantity:      "26.57000000",
		Side:          -1}

	apiTrade, err := KafkaBitfinexTradeToAPITrade(&bitfinexTradeInKafka)
	if assert.NoError(t, err) {
		assert.Equal(t, &expectedAPITrade, apiTrade)
	}
}

func TestKafkaBitfinexTradeToAPITrade2(t *testing.T) {
	msgStr, err := ioutil.ReadFile("./test-data/trade-msg-2.json")
	assert.NoError(t, err)

	bitfinexTradeInKafka := TradeStreamMessageInKafka{}
	err = json.Unmarshal(msgStr, &bitfinexTradeInKafka)
	assert.NoError(t, err)

	expectedAPITrade := types.APITrade{
		Exchange:      "bitfinex",
		Type:          "trade",
		Symbol:        "bitfinex-tLTCBTC",
		Received:      1551267619178,
		TradeID:       340960930,
		EventTime:     1551267619093,
		TradeTime:     1551267619093,
		MarketMaker:   false,
		SellerOrderID: -1,
		BuyerOrderID:  -1,
		Price:         "0.01179800",
		Quantity:      "7.89520029",
		Side:          1}

	apiTrade, err := KafkaBitfinexTradeToAPITrade(&bitfinexTradeInKafka)
	if assert.NoError(t, err) {
		assert.Equal(t, &expectedAPITrade, apiTrade)
	}
}

func TestGetDepthMessagePrimaryKey(t *testing.T) {
	msgStr, err := ioutil.ReadFile("./test-data/depth-msg-ask.json")
	assert.NoError(t, err)

	bitfinexDepthUpdateInKafka := DepthStreamMessageInKafka{}
	err = json.Unmarshal(msgStr, &bitfinexDepthUpdateInKafka)
	assert.NoError(t, err)

	assert.Equal(t, "bitfinex@tZECBTC@depth@1551267598182@0@3.99030282@4@0.01364400@2",
		GetDepthMessagePrimaryKey(&bitfinexDepthUpdateInKafka))
}

func TestGetDepthMessagePrimaryKeyPriceExp(t *testing.T) {
	msgStr, err := ioutil.ReadFile("./test-data/depth-msg-price-exp.json")
	assert.NoError(t, err)

	bitfinexDepthUpdateInKafka := DepthStreamMessageInKafka{}
	err = json.Unmarshal(msgStr, &bitfinexDepthUpdateInKafka)
	assert.NoError(t, err)

	assert.Equal(t, "bitfinex@tBATBTC@depth@1551267404491@0@51200.00000000@3@0.00000010@1",
		GetDepthMessagePrimaryKey(&bitfinexDepthUpdateInKafka))
}
