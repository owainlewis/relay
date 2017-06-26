package parser

import (
	"fmt"
	"testing"
)

func TestParserGet(t *testing.T) {

	params := map[string]string{"method": "get"}
	_, err := ParseFile("../examples/get.yml", params)

	if err != nil {
		t.Error("Cannot parse file get.yaml")
	}
}

func TestDefaultOptions(t *testing.T) {
	params := map[string]string{"method": "get"}
	req, _ := ParseFile("../examples/client-opts.yml", params)

	fmt.Println(req.Options)

	if req.Options.Timeout != 20 {
		t.Error("Expecting timeout value of 20 seconds for HTTP client options")
	}

}
