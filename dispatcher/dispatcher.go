package dispatcher

import (
	"bytes"
	"github.com/owainlewis/relay/parser"
	"github.com/owainlewis/relay/template"
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

	for k, v := range request.Headers {
		expanded, err := template.Expand(v)
		if err != nil {
			return nil, err
		}
		r.Header.Set(k, expanded)
	}
	return r, nil
}

func Run(request parser.Request) (*http.Response, error) {

	client := &http.Client{}

	httpRequest, err := ToHttpRequest(request)
	if err != nil {
		return nil, err
	}

	log.Println("Running request", httpRequest)
	response, err := client.Do(httpRequest)
	defer response.Body.Close()
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Read a request from file and return either an error or HTTP response
func FromFile(file string) (*http.Response, error) {
	req, err := parser.ParseFile(file)
	if err != nil {
		return nil, err
	}

	response, err := Run(req.Req)
	if err != nil {
		return nil, err
	}

	return response, nil
}
