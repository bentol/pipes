package event

import (
	"log"
	"os"
	"time"
	"strings"
)

type Event interface {
	GetRaw() string
	GetTimestamp() time.Time
	GetHostname() string
	Json() string
}

func GetEvents(selectedMode, data string) []Event  {
	var events []Event
	hostname, _ := os.Hostname()
	timestamp := time.Now()
	if (selectedMode == "single_value") {
		events = append(events, SingleValueEvent{data, timestamp, hostname})
	} else if (selectedMode == "map") {
		events = append(events, MapEvent{data, timestamp, hostname})
	} else if (selectedMode == "key_value") {
		lines := strings.Split(data, "\n")
		for _, line := range lines {
			events = append(events, KeyValueEvent{line, timestamp, hostname})
		}
	} else {
		log.Fatal("Mode not found or not exists")
	}
	return events
}
