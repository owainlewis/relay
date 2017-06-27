package commands

import (
	"fmt"
	"github.com/owainlewis/relay/dispatcher"
	"github.com/owainlewis/relay/parser"
	"io/ioutil"
	"os"
)

func ensureFileExists(filename string) {
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		fmt.Printf("Could not find file %s\n", filename)
		os.Exit(1)
	}
}

func getRequestFromFile(filename string, params map[string]string) (*parser.RequestItem, error) {
	req, err := parser.ParseFile(filename, params)
	if err != nil {
		return nil, err
	}
	return req, nil
}

// Runs a single request yaml file
func Run(filename string, params map[string]string) error {
	ensureFileExists(filename)

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

	fmt.Println(response.Status)
	fmt.Println(body)

	return nil
}

func AsCurl(filename string, params map[string]string) error {
	ensureFileExists(filename)
	requestItem, err := getRequestFromFile(filename, params)
	if err != nil {
		return err
	}

	fmt.Println(requestItem.Request.Curl())
	return nil
}
