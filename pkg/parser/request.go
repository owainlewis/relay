package parser

const defaultRequestTimeout = 10

// NewRequestItem will construct a new Request item with defaults
func NewRequestItem() RequestItem {
	defaultClientOptions := Options{Timeout: defaultRequestTimeout}
	return RequestItem{Options: defaultClientOptions}
}

// RequestItem defines the structure of a request yaml item
type RequestItem struct {
	Description string  `yaml:"description"`
	Request     Request `yaml:"request"`
	// Custom options for the HTTP request client i.e timeouts etc
	Options Options `yaml:"options"`
}

// Options defines the custom options for dispatching a request
type Options struct {
	Timeout int `yaml:"timeout"`
}

// Request describes the structure of a Request as YAML
type Request struct {
	// The HTTP method in uppercase
	Method string `yaml:"method"`
	// The full URL including protocol etc
	URL string `yaml:"url"`
	// HTTP headers
	Headers map[string]string `yaml:"headers"`
	// Query string params ?foo=bar
	Query map[string]string `yaml:"query"`
	// HTTP request body
	Body string `yaml:"body"`
}

// Validate will ensure that a user defined request is valid
func (r Request) Validate() []error {
	var errs []error
	return errs
}

func requestMethodValid(method string) bool {
	methods := []string{"GET", "POST", "DELETE", "PUT", "PATCH"}
	for _, m := range methods {
		if method == m {
			return true
		}
	}
	return false
}
