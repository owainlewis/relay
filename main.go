package main

import (
	"flag"
	"fmt"
	"github.com/owainlewis/relay/dispatcher"
	"github.com/owainlewis/relay/template"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

// Runs a single request yaml file
func runFile(fileName string, params map[string]string) {
	if _, err := os.Stat(fileName); os.IsNotExist(err) {
		fmt.Printf("Could not find file %s\n", fileName)
		os.Exit(1)
	}

	response, err := dispatcher.FromFile(fileName, params)
	if err != nil {
		fmt.Println(err)
	}

	log.Println(response.Status)
	body, _ := ioutil.ReadAll(response.Body)
	log.Println(body)
}

// Split the key value pairs used to define custom params into a map
func parseParams(input string) (map[string]string, error) {
	params := make(map[string]string)
	parts := strings.Fields(input)
	for _, pair := range parts {
		kv := strings.Split(pair, "=")
		if len(kv) == 2 {
			k := kv[0]
			v := kv[1]
			params[k] = v
		}
	}
	return params, nil
}

func cli() {
	// Extract the first CLI argument as the file name
	args := os.Args[1:]
	if len(args) == 0 {
		fmt.Println("Missing file")
		fmt.Println("Use: relay request.yaml -params 'foo=bar'")
		os.Exit(1)
	}

	file := args[0]

	// Convert the query params if there are any into a map

	var parameters string
	flagSet := flag.NewFlagSet("Request params", flag.ExitOnError)
	flagSet.StringVar(&parameters, "params", "", "Request params")
	flagSet.Parse(os.Args[2:])

	requestParams, err := parseParams(parameters)

	if err != nil {
		fmt.Println("Failed to parse request params", err)
		os.Exit(1)
	}

	runFile(file, requestParams)
}

func main() {
	params := map[string]string{"id": "1"}
	result, err := template.Process("Foo is {{.id}}", params)

	fmt.Println(result)

	fmt.Println(err)
}
