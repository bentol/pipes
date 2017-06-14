package event

import (
	"encoding/json"
	"log"
	"time"
	"strings"
	"strconv"
)

type KeyValueEvent struct {
	Raw string
	Timestamp time.Time
	Hostname string
}

func (event KeyValueEvent) GetRaw() string {
	return event.Raw
}

func (event KeyValueEvent) GetTimestamp() time.Time {
	return event.Timestamp
}

func (event KeyValueEvent) GetHostname() string {
	return event.Hostname
}

func (event KeyValueEvent) Json() string {
	type Template struct {
		Key string `json:"key"`
		Value string `json:"value"`
		Timestamp string `json:"@timestamp"`
		Hostname string `json:"hostname"`
	}

	components := strings.Split(event.Raw, "\t")
	tpl := Template{components[0], components[1], strconv.Itoa(int(event.Timestamp.Unix())), event.Hostname}
	result, err := json.MarshalIndent(tpl, "", "\t")
	if err != nil {
		log.Panic(err)
	}
	return string(result)
}
