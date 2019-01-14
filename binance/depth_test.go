package binance

import (
	"encoding/json"
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDepthType(t *testing.T) {
	msgStr, err := ioutil.ReadFile("./test-data/depth-msg.json")
	assert.NoError(t, err)

	msg := DepthStreamMessageInKafka{}
	err = json.Unmarshal(msgStr, &msg)
	assert.NoError(t, err)
	assert.Equal(t, int64(1547119839095), msg.RawMessage.Data.E)
	assert.Equal(t, 4, len(msg.RawMessage.Data.LowercaseB))
	assert.Equal(t, 0, len(msg.RawMessage.Data.LowercaseA))
	assert.Equal(t, "0.00999800", msg.RawMessage.Data.LowercaseB[2][0])
}
