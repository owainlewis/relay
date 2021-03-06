package parser

import "testing"

func TestValidateRequestMethod(t *testing.T) {
	methods := []string{"GET", "POST", "PUT", "PATCH"}
	for _, m := range methods {
		if !requestMethodValid(m) {
			t.Error(m + " should be a valid request method")
		}
	}
}
