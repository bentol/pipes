package event_test

import (
	. "github.com/onsi/gomega"
	"testing"
	. "github.com/bentol/pipes/event"
	"github.com/stretchr/testify/assert"
	"os"
)

func TestGetEvents_SingleValueMode(t *testing.T) {
	RegisterTestingT(t)
	hostname, _ := os.Hostname()

	events := GetEvents("single_value", "777")

	assert.Len(t, events, 1)
	assert.Equal(t, events[0].GetRaw(), "777")
	assert.Equal(t, events[0].GetHostname(), hostname)
}

func TestGetEvents_KeyValueMode(t *testing.T) {
	RegisterTestingT(t)
	hostname, _ := os.Hostname()

	events := GetEvents("key_value", "apple\t777\nmanggo\t999\ngrape\t55")

	assert.Len(t, events, 3)
	assert.Equal(t, events[0].GetRaw(), "apple\t777")
	assert.Equal(t, events[0].GetHostname(), hostname)
	assert.Equal(t, events[1].GetRaw(), "manggo\t999")
	assert.Equal(t, events[1].GetHostname(), hostname)
	assert.Equal(t, events[2].GetRaw(), "grape\t55")
	assert.Equal(t, events[2].GetHostname(), hostname)
}
