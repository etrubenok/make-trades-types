package bitfinex

import (
	"encoding/json"
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTradeType(t *testing.T) {
	msgStr, err := ioutil.ReadFile("./test-data/trade-msg.json")
	assert.NoError(t, err)

	msg := TradeStreamMessageInKafka{}
	err = json.Unmarshal(msgStr, &msg)
	assert.NoError(t, err)
	assert.Equal(t, int64(1551071701829), msg.RawMessage.MTS)
}
