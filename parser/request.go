package parser

import (
	"fmt"
	"strings"
)

const defaultRequestTimeout = 10

func NewRequestItem() RequestItem {
	defaultClientOptions := Options{Timeout: defaultRequestTimeout}
	return RequestItem{Options: defaultClientOptions}
}

type RequestItem struct {
	Description string  `yaml:"description"`
	Request     Request `yaml:"request"`
	// Custom options for the HTTP request client i.e timeouts etc
	Options Options `yaml:"options"`
}

type Options struct {
	Timeout int `yaml:"timeout"`
}

type Request struct {
	// The HTTP method in uppercase
	Method string `yaml:"method"`
	// The full URL including protocol etc
	Url string `yaml:"url"`
	// HTTP headers
	Headers map[string]string `yaml:"headers"`
	// Query string params ?foo=bar
	Query map[string]string `yaml:"query"`
	// HTTP request body
	Body string `yaml:"body"`
}

// Returns a request as a CURL command
func (request *Request) Curl() string {
	headerParts := []string{}
	for k, v := range request.Headers {
		header := fmt.Sprintf("-H '%s: %s'", k, v)
		headerParts = append(headerParts, header)
	}
	hdrs := strings.Join(headerParts, " ")
	return fmt.Sprintf("curl -X %s %s %s", request.Method, hdrs, request.Url)
}
