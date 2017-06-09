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

func TestValidateRequestMethod(t *testing.T) {
	methods := []string{"GET", "POST", "PUT", "PATCH"}
	for _, m := range methods {
		if !RequestMethodValid(m) {
			t.Error(m + " should be a valid request method")
		}
	}
}
