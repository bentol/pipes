package event

import (
	"encoding/json"
	"log"
	"time"
	"strconv"
)

type SingleValueEvent struct {
	Raw string
	Timestamp time.Time
	Hostname string
}

func (event SingleValueEvent) GetRaw() string {
	return event.Raw
}

func (event SingleValueEvent) GetTimestamp() time.Time {
	return event.Timestamp
}

func (event SingleValueEvent) GetHostname() string {
	return event.Hostname
}

func (event SingleValueEvent) Json() string {
	type Template struct {
		Value string `json:"value"`
		Timestamp string `json:"@timestamp"`
		Hostname string `json:"hostname"`
	}

	tpl := Template{event.Raw, strconv.Itoa(int(event.Timestamp.Unix())), event.Hostname}
	result, err := json.MarshalIndent(tpl, "", "\t")
	if err != nil {
		log.Panic(err)
	}
	return string(result)
}