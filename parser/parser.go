package parser

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type RequestItem struct {
	Description string  `yaml:"description"`
	Req         Request `yaml:"request"`
}

type Request struct {
	// The HTTP method in uppercase
	Method string `yaml:"method"`
	// The full URL
	Url string `yaml:"url"`
	// HTTP headers
	Headers map[string]string `yaml:"headers"`
	// Query string params ?foo=bar
	Query map[string]string `yaml:"query"`
	// An optional HTTP request body
	Body string `yaml:"body"`
}

func Parse(data []byte) (*RequestItem, error) {
	var req = &RequestItem{}

	err := yaml.Unmarshal(data, req)
	if err != nil {
		return nil, err
	}

	return req, nil
}

func ParseFile(file string) (*RequestItem, error) {
	contents, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}

	return Parse(contents)
}
