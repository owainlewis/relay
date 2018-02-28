package parser

import (
	"io/ioutil"

	"github.com/owainlewis/relay/pkg/template"
	"gopkg.in/yaml.v2"
)

// Parse will turn a raw byte array into a RequestItem
func Parse(data []byte) (*RequestItem, error) {
	defaultClientOptions := Options{Timeout: 10}
	var req = &RequestItem{Options: defaultClientOptions}

	err := yaml.Unmarshal(data, req)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// ParseFile will a file returning a RequestItem. We also do template interpolation at this stage
func ParseFile(file string, params map[string]string) (*RequestItem, error) {
	contents, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}

	expanded, err := template.Expand(string(contents), params)
	if err != nil {
		return nil, err
	}

	return Parse(expanded)
}
