package dispatcher

import (
	"bytes"
	"fmt"
	"github.com/owainlewis/relay/parser"
	"log"
	"net/http"
)

func ToHttpRequest(request parser.Request) (*http.Request, error) {
	payloadBytes := []byte(request.Body)
	body := bytes.NewBuffer(payloadBytes)

	r, err := http.NewRequest(request.Method, request.Url, body)
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

func showRequest(request *http.Request) string {
	return fmt.Sprintf("%s", request.URL.String())
}

func Run(request parser.Request) (*http.Response, error) {
	client := &http.Client{}

	httpRequest, err := ToHttpRequest(request)
	if err != nil {
		return nil, err
	}

	log.Println("Request:", showRequest(httpRequest))

	response, err := client.Do(httpRequest)
	defer response.Body.Close()
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Read a request from file and return either an error or HTTP response
func FromFile(file string, params map[string]string) (*http.Response, error) {
	req, err := parser.ParseFile(file, params)
	if err != nil {
		return nil, err
	}

	response, err := Run(req.Req)
	if err != nil {
		return nil, err
	}

	return response, nil
}
