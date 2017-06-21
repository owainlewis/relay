package main

import (
	"flag"
	"fmt"
	"github.com/owainlewis/relay/dispatcher"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func runFile(fileName string, verbose bool) {
	if _, err := os.Stat(fileName); os.IsNotExist(err) {
		fmt.Printf("Could not find file %s\n", fileName)
		os.Exit(1)
	}

	response, err := dispatcher.FromFile(fileName, verbose)
	if err != nil {
		fmt.Println(err)
	}

	log.Println(response.Status)
	body, _ := ioutil.ReadAll(response.Body)
	log.Println(body)
}

func cli() {
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

func main() {
	// Extract the first CLI argument as the file name
	args := os.Args
	if len(args) == 1 {
		fmt.Println("Missing file")
		fmt.Println("Use: relay request.yaml -params 'foo=bar'")
		os.Exit(1)
	}

	file := args[1]

	fmt.Println("File is", file)

	// Convert the query params if there are any into a map

	var parameters string
	flagSet := flag.NewFlagSet("Request params", flag.ExitOnError)
	flagSet.StringVar(&parameters, "params", "", "Request params")
	flagSet.Parse(os.Args[2:])

	requestParams, _ := parseParams(parameters)

	runFile(file, requestParams)
}
