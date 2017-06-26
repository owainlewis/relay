package parser

import "testing"

func TestParserGet(t *testing.T) {

	params := map[string]string{"method": "get"}
	_, err := ParseFile("../examples/get.yml", params)

	if err != nil {
		t.Error("Cannot parse file get.yaml")
	}
}
