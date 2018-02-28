package dispatcher

import (
	"bytes"
	"net/http"
	"time"

	"github.com/owainlewis/relay/pkg/parser"
)

func toHTTPRequest(request parser.Request) (*http.Request, error) {
	body := bytes.NewBuffer([]byte(request.Body))

	r, err := http.NewRequest(request.Method, request.URL, body)
	if err != nil {
		return nil, err
	}

	// Add HTTP headers to the request
	for k, v := range request.Headers {
		r.Header.Set(k, v)
	}

	// Add query params to the request
	if len(request.Query) > 0 {
		q := r.URL.Query()
		for k, v := range request.Query {
			q.Add(k, v)
		}
		r.URL.RawQuery = q.Encode()
	}

	return r, nil
}

// Run will execute a request item
func Run(request *parser.RequestItem) (*http.Response, error) {
	var client = &http.Client{
		Timeout: time.Second * 10,
	}

	httpRequest, err := toHTTPRequest(request.Request)
	if err != nil {
		return nil, err
	}

	return client.Do(httpRequest)
}
