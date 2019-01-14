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
