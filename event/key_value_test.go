package event_test

import (
	. "github.com/onsi/gomega"
	"time"
	"testing"
	. "github.com/bentol/pipes/event"
	"github.com/stretchr/testify/assert"
	"encoding/json"
	"strconv"
)

func TestKeyValueEvent_Json(t *testing.T) {
	RegisterTestingT(t)

	rawData := "apple\t2"

	timestamp := time.Now()
	eventObj := KeyValueEvent{
		Raw: rawData,
		Timestamp: timestamp,
		Hostname: "my-host",
	}
	expectedJson, _ := json.Marshal(map[string]string{
		"key": "apple",
		"value": "2",
		"@timestamp": strconv.Itoa(int(timestamp.Unix())),
		"hostname": "my-host",
	})
	assert.JSONEq(t, eventObj.Json(), string(expectedJson))
}
