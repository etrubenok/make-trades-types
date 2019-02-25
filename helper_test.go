package types

import (
	"io/ioutil"
	"testing"

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

func TestRandStringBytes(t *testing.T) {
	str := RandStringBytes(10)
	assert.Equal(t, 10, len(str))
}
