package dispatcher

import (
	"bytes"
	"fmt"
	"github.com/owainlewis/relay/parser"
	"log"
	"net/http"
	"time"
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
	return fmt.Sprintf("%s %s", request.Method, request.URL.String())
}

func Run(request *parser.RequestItem) (*http.Response, error) {
	var client = &http.Client{
		Timeout: time.Second * 10,
	}

	httpRequest, err := ToHttpRequest(request.Request)
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
