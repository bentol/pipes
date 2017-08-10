package event

import (
	//"encoding/json"
	//"log"
	//"time"
	//"strconv"
)
import (
	"time"
	"strings"
	"strconv"
	"encoding/json"
	"log"
)

type MapEvent struct {
	Raw string
	Timestamp time.Time
	Hostname string
}

func (event MapEvent) GetRaw() string {
	return event.Raw
}

func (event MapEvent) GetTimestamp() time.Time {
	return event.Timestamp
}

func (event MapEvent) GetHostname() string {
	return event.Hostname
}

func (event MapEvent) Json() string {
	var mapping = map[string]string{}
	mapping["@timestamp"] = strconv.Itoa(int(event.Timestamp.Unix()))
	mapping["hostname"] = event.Hostname

	lines := strings.Split(event.Raw, "\n")
	for _, line := range lines {
		kv := strings.Split(line, "\t")
		mapping[kv[0]] = kv[1]
	}

	result, err := json.MarshalIndent(mapping, "", "\t")
	if err != nil {
		log.Panic(err)
	}
	return string(result)
}