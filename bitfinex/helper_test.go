package bitfinex

import (
	"encoding/json"
	"io/ioutil"
	"testing"

	bitfinex "github.com/bitfinexcom/bitfinex-api-go/v2"
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
	}
}

func TestGetDepthStreamMessage(t *testing.T) {
	msgStr, err := ioutil.ReadFile("./test-data/depth-msg.json")
	assert.NoError(t, err)
	msg, err := GetDepthStreamMessage(msgStr)
	if assert.NoError(t, err) {
		assert.Equal(t, "bitfinex", msg.Exchange)
		assert.Equal(t, 218.84835729, msg.RawMessage.Amount)
		assert.Equal(t, 3.7475, msg.RawMessage.Price)
		assert.Equal(t, json.Number("-218.84835729"), msg.RawMessage.AmountJsNum)
		assert.Equal(t, json.Number("3.7475"), msg.RawMessage.PriceJsNum)
		assert.Equal(t, bitfinex.BookAction(0), msg.RawMessage.Action)
		assert.Equal(t, bitfinex.OrderSide(2), msg.RawMessage.Side)
	}
}
