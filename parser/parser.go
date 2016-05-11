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

func ParseFile(file string) (*RequestItem, error) {
	contents, err := ioutil.ReadFile(file)

	if err != nil {
		return nil, err
	}

	var req = &RequestItem{}
	json.Unmarshal(contents, req)
	return req, nil
}
