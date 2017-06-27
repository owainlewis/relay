package main

import (
	"flag"
	"fmt"
	"github.com/owainlewis/relay/commands"
	"os"
	"strings"
)

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

func handleAsCurl(args []string) {
	if len(args) == 0 {
		fmt.Println("Missing file for curl command")
		os.Exit(1)
	}

	file := args[0]
	requestParams, err := extractParamsFromArgs(args)
	if err != nil {
	} else {
		ok := commands.AsCurl(file, requestParams) == nil
		if !ok {
			fmt.Println("Failed to execute curl command", err)
		}
	}
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

	ok := commands.Run(file, requestParams) == nil
	if !ok {
		fmt.Println("Failed to execute run command", err)
	}
}

func main() {
	args := os.Args[1:]
	cmd := args[0]
	cmdArgs := args[1:]

	switch cmd {
	case "run":
		handleRun(cmdArgs)
	case "curl":
		handleAsCurl(cmdArgs)
	default:
		fmt.Println("Use: relay run request.yaml -params 'foo=bar'")
	}
}
