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

func TestMapEvent_Json (t *testing.T) {
	RegisterTestingT(t)

	timestamp := time.Now()
	eventObj := MapEvent{
		Raw: "apple\t2\nmanggo\t10\ngrape\t99999",
		Timestamp: timestamp,
		Hostname: "my-host",
	}
	expectedJson, _ := json.Marshal(map[string]string{
		"apple": "2",
		"manggo": "10",
		"grape": "99999",
		"@timestamp": strconv.Itoa(int(timestamp.Unix())),
		"hostname": "my-host",
	})
	assert.JSONEq(t, eventObj.Json(), string(expectedJson))
}
