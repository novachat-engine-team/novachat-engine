package log

import (
	"encoding/json"
)

var (
	// empty
	emptyBytes = []byte("{}")
)

func JsonDebugData(message interface{}) []byte {
	if data, err := json.Marshal(message); err == nil {
		return data
	}
	return emptyBytes
}
