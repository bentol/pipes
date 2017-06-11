package rule

import (
	"encoding/json"
	"log"
)

type SingleValueRule struct {
	raw string
	timestamp int32
	hostname string
}

func (rule SingleValueRule) Json() string {
	type Template struct {
		Value string `json:"value"`
		Timestamp int32 `json:"@timestamp"`
		Hostname string `json:"hostname"`
	}

	tpl := Template{rule.raw, rule.timestamp, rule.hostname}
	result, err := json.MarshalIndent(tpl, "", " ")
	if err != nil {
		log.Panic(err)
	}
	return string(result)
}