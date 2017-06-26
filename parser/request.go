package parser

import "fmt"

func NewRequestItem() RequestItem {
	defaultClientOptions := Options{Timeout: 10}
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
	return fmt.Sprintf("curl -X%s %s", request.Method, request.Url)
}
