package main

import (
	flag "github.com/ogier/pflag"
	//"github.com/bentol/pipes/rules"
	"bufio"
	"os"
	"github.com/bentol/pipes/rules"
	"net/http"
	"log"
	"bytes"
	"fmt"
)

var (
	host string
	port string
	user string
	password string
	types string
	index string
	selectedRule string
)

func main()  {
	flag.Parse()

	input := getInput()
	payload := buildPayload(input)
	fmt.Println(payload)
	makeRequest(payload)
}

func init()  {
	flag.StringVarP(&host, "host", "h", "localhost", "Elasticsearch host")
	flag.StringVarP(&port, "port", "P", "9200", "Elasticsearch port")
	flag.StringVarP(&user, "user", "u", "", "Basic auth username")
	flag.StringVarP(&password, "password", "p", "", "Basic auth password")
	flag.StringVarP(&types, "type", "t", "log", "Index log type")
	flag.StringVarP(&index, "index", "i", "", "Index name")
	flag.StringVarP(&selectedRule, "rule", "r", "single_value", "Rule")
}

func makeRequest(payload string) {
	client := &http.Client{}

	req, _ := http.NewRequest("POST", "http://" + host + ":" + port + "/" + index + "/" + types, bytes.NewReader([]byte(payload)))
	req.Header.Add("Content-Type", "application/json")
	_, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
}

func getInput() string {
	buffer := ""
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		buffer += scanner.Text()
	}
	return buffer
}

func buildPayload(input string) string {
	ruleObj := rule.GetRule(selectedRule, input)
	payload := ruleObj.Json()
	return payload
}
