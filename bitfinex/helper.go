package bitfinex

import (
	"crypto/sha1"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"github.com/golang/glog"
)

// GetTradeStreamMessage converts a message in the string format into TradeStreamMessageInKafka
func GetTradeStreamMessage(message []byte) (*TradeStreamMessageInKafka, error) {
	msg := TradeStreamMessageInKafka{}
	err := json.Unmarshal(message, &msg)
	if err != nil {
		glog.Errorf("bitfinex.GetTradeStreamMessage: cannot convert message %s into TradeStreamMessageInKafka due to error %s", message, err)
		return nil, err
	}
	return &msg, nil
}

// GetDepthStreamMessage converts a message in the string format into DepthStreamMessageInKafka
func GetDepthStreamMessage(message []byte) (*DepthStreamMessageInKafka, error) {
	msg := DepthStreamMessageInKafka{}
	err := json.Unmarshal(message, &msg)
	if err != nil {
		glog.Errorf("GetDepthStreamMessage: cannot convert message %s into DepthStreamMessageInKafka due to error %s", message, err)
		return nil, err
	}
	return &msg, nil
}

// GetTradeMessagePrimaryKey gets a concatenated string that forms from the message fields which identify the trade uniquely
func GetTradeMessagePrimaryKey(trade *TradeStreamMessageInKafka) string {
	fields := make([]string, 0)
	fields = append(fields, trade.Exchange)
	fields = append(fields, trade.Stream)
	fields = append(fields, strconv.FormatInt(trade.EventTime, 10))
	fields = append(fields, strconv.FormatInt(trade.RawMessage.ID, 10))
	return strings.Join(fields, "@")
}

// GetDepthMessagePrimaryKey gets a concatenated string that forms from the message fields which identify the depth message uniquely
func GetDepthMessagePrimaryKey(trade *DepthStreamMessageInKafka) string {
	fields := make([]string, 0)
	fields = append(fields, trade.Exchange)
	fields = append(fields, trade.Stream)
	fields = append(fields, strconv.FormatInt(trade.EventTime, 10))
	fields = append(fields, strconv.FormatInt(trade.RawMessage.ID, 10))
	return strings.Join(fields, "@")
}

// GetKeyHash converts a plain string key into hash
func GetKeyHash(key string) (string, error) {
	h := sha1.New()
	_, err := h.Write([]byte(key))
	if err != nil {
		glog.Errorf("GetKeyHash: cannot make a hash for key %s due to error %s", key, err)
		return "", err
	}
	bs := h.Sum(nil)
	return fmt.Sprintf("%x", bs), nil
}
