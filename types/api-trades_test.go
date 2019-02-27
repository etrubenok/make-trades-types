package types

import (
	"encoding/json"
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAPITradeMarshaling(t *testing.T) {
	msgStr, err := ioutil.ReadFile("./test-data/api-trade.json")
	assert.NoError(t, err)

	trade := APITrade{}
	err = json.Unmarshal(msgStr, &trade)
	if assert.NoError(t, err) {
		assert.Equal(t, -1, trade.Side)
	}
}

func TestAPITradeMarshalingNoSide(t *testing.T) {
	msgStr, err := ioutil.ReadFile("./test-data/api-trade-no-side.json")
	assert.NoError(t, err)

	trade := APITrade{}
	err = json.Unmarshal(msgStr, &trade)
	if assert.NoError(t, err) {
		assert.Equal(t, 0, trade.Side)
	}
}
