package parser

import (
	"testing"
)

func TestParserGet(t *testing.T) {
	_, err := ParseFile("../examples/get.json")

	if err != nil {
		t.Error("Cannot parse file get.json")
	}
}

func TestParserPost(t *testing.T) {
	_, err := ParseFile("../examples/post.json")

	if err != nil {
		t.Error("Can't parse file post.json")
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
