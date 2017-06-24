package parser

import "fmt"

type RequestItem struct {
	Description string  `yaml:"description"`
	Req         Request `yaml:"request"`
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
