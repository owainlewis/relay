package parser

import (
	"testing"
)

func TestParser(t *testing.T) {
	request, err := ParseFile("../examples/get.json")

	if err != nil {
		t.Error(err)
	}

	if request.Req.Method != "GET" {
		t.Error("Invalid status")
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
