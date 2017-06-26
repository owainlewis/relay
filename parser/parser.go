package parser

import (
	"github.com/owainlewis/relay/template"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

func Parse(data []byte) (*RequestItem, error) {
	//	defaultClientOpts := Options{Timeout: 10}
	var req = &RequestItem{}

	err := yaml.Unmarshal(data, req)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// Parse a file. We also do template interpolation at this stage
//
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
