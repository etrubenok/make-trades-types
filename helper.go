package types

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/golang/glog"
)

// ConvertToJSON converts string to JSON struct
func ConvertToJSON(message string) (map[string]interface{}, error) {
	m := make(map[string]interface{})

	d := json.NewDecoder(bytes.NewBuffer([]byte(message)))
	d.UseNumber()
	if err := d.Decode(&m); err != nil {
		e := fmt.Errorf("ConvertToJSON: cannot unmarshal message %s to JSON due to error %s", message, err)
		glog.Error(e)
		return m, e
	}

	return m, nil
}

// GetStreamType returns type of the message
func GetStreamType(msg *map[string]interface{}) (string, error) {
	s, ok := (*msg)["stream"].(string)
	if !ok {
		return "", fmt.Errorf("GetStreamType: cannot get 'stream' from msg %v", msg)
	}
	tokens := strings.Split(s, "@")
	if len(tokens) != 2 {
		return "", fmt.Errorf("GetStreamType: the format of the 'stream': %s is not what we expect (it should have two tokens)", s)
	}
	return tokens[1], nil
}

// GetExchange return exchange name from the message
func GetExchange(msg *map[string]interface{}) (string, error) {
	e, ok := (*msg)["exchange"].(string)
	if !ok {
		return "", fmt.Errorf("GetExchange: cannot get 'exchange' from msg %v", msg)
	}
	return e, nil
}
