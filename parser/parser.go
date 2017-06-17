package parser

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type RequestItem struct {
	Schema      string  `yaml:"schema"`
	Description string  `yaml:"description"`
	Req         Request `yaml:"request"`
}

type Request struct {
	Method  string            `yaml:"method"`
	Url     string            `yaml:"url"`
	Headers map[string]string `yaml:"headers"`
	Body    string            `yaml:"body"`
}

type Response struct {
	Status int
	Body   string
}

func Parse(data []byte) (*RequestItem, error) {
	var req = &RequestItem{}
	yaml.Unmarshal(data, req)
	return req, nil
}

func ParseFile(file string) (*RequestItem, error) {
	contents, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}

	return Parse(contents)
}
