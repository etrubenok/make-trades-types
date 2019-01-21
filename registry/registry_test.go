package registry

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetExchangeID(t *testing.T) {
	id, err := GetExchangeID("binance")
	if assert.NoError(t, err) {
		assert.Equal(t, 1, id)
	}
}

func TestGetExchangeIDError(t *testing.T) {
	name := "unknown"
	_, err := GetExchangeID(name)
	assert.EqualError(t, err, fmt.Sprintf("GetExchangeID: exchange: '%s' is not known", name))
}

func TestGetExchangeNameByID(t *testing.T) {
	id := 1
	ex, err := GetExchangeNameByID(id)
	if assert.NoError(t, err) {
		assert.Equal(t, "binance", ex)
	}
}

func TestGetExchangeNameByIDError(t *testing.T) {
	id := 999
	_, err := GetExchangeNameByID(id)
	assert.EqualError(t, err, fmt.Sprintf("GetExchangeNameByID does not know exchange with ID: %d", id))
}
