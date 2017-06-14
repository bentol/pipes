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

func TestSingleValueEvent_Json(t *testing.T) {
	RegisterTestingT(t)


	timestamp := time.Now()
	eventObj := SingleValueEvent{
		Raw: "9",
		Timestamp: timestamp,
		Hostname: "my-host",
	}
	expectedJson, _ := json.Marshal(map[string]string{
		"value": "9",
		"@timestamp": strconv.Itoa(int(timestamp.Unix())),
		"hostname": "my-host",
	})
	assert.JSONEq(t, eventObj.Json(), string(expectedJson))
}
