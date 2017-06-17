package parser

import (
	"testing"
)

func TestParserGet(t *testing.T) {
	_, err := ParseFile("../examples/get.yaml")

	if err != nil {
		t.Error("Cannot parse file get.yaml")
	}
}
