package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/owainlewis/relay/pkg/dispatcher"
	"github.com/owainlewis/relay/pkg/parser"
)

var version = ""

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

// Extra a map of strings to use as params from CLI args.
// Note the abstraction is a bit vague here (which args to pass).
// Tighten this up when abstractions are clearer
func extractParamsFromArgs(args []string) (map[string]string, error) {
	var parameters string
	flagSet := flag.NewFlagSet("Request params", flag.ExitOnError)
	flagSet.StringVar(&parameters, "params", "", "Request params")
	flagSet.Parse(args[1:])

	requestParams, err := parseParams(parameters)
	if err != nil {
		return nil, err
	}

	return requestParams, nil
}

func getRequestFromFile(filename string, params map[string]string) (*parser.RequestItem, error) {
	req, err := parser.ParseFile(filename, params)
	if err != nil {
		return nil, err
	}
	return req, nil
}

// Runs a single request yaml file
func run(filename string, params map[string]string) error {
	requestItem, err := getRequestFromFile(filename, params)
	if err != nil {
		return err
	}

	response, err := dispatcher.Run(requestItem)
	if err != nil {
		return err
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return err
	}

	fmt.Printf("Response: %#v\n", response)
	fmt.Println(string(body))

	return nil
}

func handleRun(args []string) {
	if len(args) == 0 {
		fmt.Println("Missing file for run command")
		os.Exit(1)
	}

	file := args[0]

	requestParams, err := extractParamsFromArgs(args)
	if err != nil {
		fmt.Println("Failed to parse request params", err)
		os.Exit(1)
	}

	ok := run(file, requestParams) == nil
	if !ok {
		fmt.Println("Failed to execute run command", err)
	}
}

func main() {
	handleRun(os.Args[1:])
}
