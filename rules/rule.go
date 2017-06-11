package rule

import (
	"log"
	"os"
	"time"
)

type Rule interface {
	Json() string
}

func GetRule(selectedRule, data string) Rule  {
	var ruleObj Rule
	hostname, _ := os.Hostname()
	timestamp := getTimestamp()
	if (selectedRule == "single_value") {
		ruleObj = SingleValueRule{data, timestamp, hostname}
	} else {
		log.Panic("Rule not found or not exists")
	}
	return ruleObj
}

func getTimestamp() int32 {
	return int32(time.Now().Unix())
}