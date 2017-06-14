package main

import (
	flag "github.com/ogier/pflag"
	"bufio"
	"os"
	"github.com/bentol/pipes/event"
	"net/http"
	"log"
	"bytes"
	"fmt"
	"io/ioutil"
	"strings"
)

var (
	host string
	port string
	user string
	password string
	types string
	index string
	selectedMode string
	verboseMode bool
	dryRunMode bool
)

func main()  {
	flag.Parse()

	input := getInput()
	events := event.GetEvents(selectedMode, input)
	for _, eventObj := range events {
		payload := eventObj.Json()

		if dryRunMode || verboseMode {
			fmt.Println(payload)
		}

		if !dryRunMode {
			makeRequest(payload)
		}
	}
}

func init()  {
	flag.StringVarP(&host, "host", "h", "localhost", "Elasticsearch host")
	flag.StringVarP(&port, "port", "P", "9200", "Elasticsearch port")
	flag.StringVarP(&user, "user", "u", "", "Basic auth username")
	flag.StringVarP(&password, "password", "p", "", "Basic auth password")
	flag.StringVarP(&types, "type", "t", "log", "Index log type")
	flag.StringVarP(&index, "index", "i", "", "Index name")
	flag.StringVarP(&selectedMode, "mode", "m", "single_value", "Mode (single_value|key_value)")
	flag.BoolVarP(&verboseMode, "verbose", "v", false, "Verbose output")
	flag.BoolVarP(&dryRunMode, "dry-run", "", false, "Enable dry-run mode (just output json, without make request)")
}

func makeRequest(payload string) {
	client := &http.Client{}

	req, _ := http.NewRequest("POST", "http://" + host + ":" + port + "/" + index + "/" + types, bytes.NewReader([]byte(payload)))
	req.Header.Add("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	responseData, _ := ioutil.ReadAll(resp.Body)

	if verboseMode {
		fmt.Println("Response: ", string(responseData))
	}
}

func getInput() string {
	buffer := ""
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		buffer += scanner.Text() + "\n"
	}
	buffer = strings.Trim(buffer, "\n")
	return buffer
}
