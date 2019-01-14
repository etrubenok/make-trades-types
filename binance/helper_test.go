package binance

import (
	"io/ioutil"
	"testing"

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
