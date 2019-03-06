package bitfinex

import (
	"encoding/json"
	"io/ioutil"
	"testing"

	bitfinexOrig "github.com/etrubenok/bitfinex-api-go/v2"
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

func TestKafkaBitfinexFundingOrderBookUpdateToAPIFundingOrderBookUpdate(t *testing.T) {
	msgStr, err := ioutil.ReadFile("./test-data/fund-depth-msg.json")
	assert.NoError(t, err)

	bitfinexDepthUpdateInKafka := FundingDepthStreamMessageInKafka{}
	err = json.Unmarshal(msgStr, &bitfinexDepthUpdateInKafka)
	assert.NoError(t, err)

	expectedAPIOrderBookUpdate := types.APIFundingOrderBookUpdate{
		Exchange:      "bitfinex",
		Type:          "forderbook",
		Symbol:        "bitfinex-fXRP",
		Received:      1551831768843,
		FirstUpdateID: 0,
		EventTime:     1551831768843,
		LastUpdateID:  0,
		Asks:          []types.APIFundingOrderBookPriceLevel{types.APIFundingOrderBookPriceLevel{Rate: "0.00010000", Period: 2, Quantity: "2452.78422300"}},
		Bids:          []types.APIFundingOrderBookPriceLevel{}}

	apiOrdeBookUpdate, err := KafkaBitfinexFundingOrderBookUpdateToAPIFundingOrderBookUpdate(&bitfinexDepthUpdateInKafka)
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

	assert.Equal(t, "bitfinex@tZECBTC@depth@0@3.99030282@4@0.01364400@2",
		GetDepthMessagePrimaryKey(&bitfinexDepthUpdateInKafka))
}

func TestGetDepthMessagePrimaryKeyPriceExp(t *testing.T) {
	msgStr, err := ioutil.ReadFile("./test-data/depth-msg-price-exp.json")
	assert.NoError(t, err)

	bitfinexDepthUpdateInKafka := DepthStreamMessageInKafka{}
	err = json.Unmarshal(msgStr, &bitfinexDepthUpdateInKafka)
	assert.NoError(t, err)

	assert.Equal(t, "bitfinex@tBATBTC@depth@0@51200.00000000@3@0.00000010@1",
		GetDepthMessagePrimaryKey(&bitfinexDepthUpdateInKafka))
}

func TestGetFundingDepthMessagePrimaryKey(t *testing.T) {
	msgStr, err := ioutil.ReadFile("./test-data/fund-depth-msg.json")
	assert.NoError(t, err)

	bitfinexDepthUpdateInKafka := FundingDepthStreamMessageInKafka{}
	err = json.Unmarshal(msgStr, &bitfinexDepthUpdateInKafka)
	assert.NoError(t, err)

	assert.Equal(t, "bitfinex@fXRP@fdepth@0@2452.78422300@5@0.00010000@2",
		GetFundingDepthMessagePrimaryKey(&bitfinexDepthUpdateInKafka))
}

func TestGetFundingDepthMessagePrimaryKeyTheSame(t *testing.T) {
	msgStr, err := ioutil.ReadFile("./test-data/fund-depth-msg.json")
	assert.NoError(t, err)

	bitfinexDepthUpdateInKafka := FundingDepthStreamMessageInKafka{}
	err = json.Unmarshal(msgStr, &bitfinexDepthUpdateInKafka)
	assert.NoError(t, err)

	msgStr2, err := ioutil.ReadFile("./test-data/fund-depth-msg-dup.json")
	assert.NoError(t, err)

	bitfinexDepthUpdateInKafka2 := FundingDepthStreamMessageInKafka{}
	err = json.Unmarshal(msgStr2, &bitfinexDepthUpdateInKafka2)
	assert.NoError(t, err)

	assert.Equal(t, GetFundingDepthMessagePrimaryKey(&bitfinexDepthUpdateInKafka),
		GetFundingDepthMessagePrimaryKey(&bitfinexDepthUpdateInKafka2))

	hash, err := GetKeyHash(GetFundingDepthMessagePrimaryKey(&bitfinexDepthUpdateInKafka))
	assert.NoError(t, err)

	hash2, err := GetKeyHash(GetFundingDepthMessagePrimaryKey(&bitfinexDepthUpdateInKafka2))
	assert.NoError(t, err)

	assert.Equal(t, hash,
		hash2)
}

func TestGetKeyHash(t *testing.T) {
	msgStr, err := ioutil.ReadFile("./test-data/fund-depth-msg-2.json")
	assert.NoError(t, err)

	bitfinexDepthUpdateInKafka := FundingDepthStreamMessageInKafka{}
	err = json.Unmarshal(msgStr, &bitfinexDepthUpdateInKafka)
	assert.NoError(t, err)

	assert.Equal(t, "bitfinex@fDSH@fdepth@0@1.00000000@0@0.00120000@2",
		GetFundingDepthMessagePrimaryKey(&bitfinexDepthUpdateInKafka))

	hash, err := GetKeyHash(GetFundingDepthMessagePrimaryKey(&bitfinexDepthUpdateInKafka) + "@v5")
	assert.NoError(t, err)
	assert.Equal(t, "4c735f899c853a87b0d30fb14411e07220f04fa9", hash)
}

func TestKafkaBitfinexFundingTradeToAPIFundingTrade(t *testing.T) {

	msgStr, err := ioutil.ReadFile("./test-data/fund-trade-msg.json")
	assert.NoError(t, err)

	bitfinexTradeUpdateInKafka := FundingTradeStreamMessageInKafka{}
	err = json.Unmarshal(msgStr, &bitfinexTradeUpdateInKafka)

	expectedAPIFundingTrade := types.APIFundingTrade{
		Exchange:      "bitfinex",
		Type:          "ftrade",
		Symbol:        "bitfinex-fBTC",
		Received:      1551840569673,
		TradeID:       104413077,
		EventTime:     1551840570000,
		TradeTime:     1551840570000,
		MarketMaker:   false,
		SellerOrderID: -1,
		BuyerOrderID:  -1,
		Rate:          "0.00000410",
		Period:        10,
		Quantity:      "0.23744000",
		Side:          -1}

	apiFundingTrade, err := KafkaBitfinexFundingTradeToAPIFundingTrade(&bitfinexTradeUpdateInKafka)
	if assert.NoError(t, err) {
		assert.Equal(t, &expectedAPIFundingTrade, apiFundingTrade)
	}
}
