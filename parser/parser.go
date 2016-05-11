package parser

import (
	"encoding/json"
	"io/ioutil"
)

type RequestItem struct {
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Req         Request `json:"request"`
}

type Request struct {
	Method  string            `json:"method"`
	Url     string            `json:"url"`
	Headers map[string]string `json:"headers"`
	Body    json.RawMessage   `json:"body"`
}

type Response struct {
	Status int
	Body   string
}

func Parse(data []byte) (*RequestItem, error) {
	var req = &RequestItem{}
	json.Unmarshal(data, req)
	return req, nil
}

// Checks if a given request method is valid
func RequestMethodValid(method string) bool {
	methods := []string{"GET", "POST", "DELETE", "PUT", "PATCH"}
	for _, m := range methods {
		if method == m {
			return true
		}
	}
	return false
}

// Parse a single request from a .json file
func ParseFile(file string) (*RequestItem, error) {
	contents, err := ioutil.ReadFile(file)

	if err != nil {
		return nil, err
	}

	return Parse(contents)
}
