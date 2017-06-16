package dispatcher

import (
	"bytes"
	"github.com/owainlewis/relay/parser"
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
		r.Header.Set(k, v)
	}
	return r, nil
}

func Run(request parser.Request) (*http.Response, error) {
	client := &http.Client{}
	
	hRequest, err := ToHttpRequest(request)
	if err != nil {
		return nil, err
	}
	
	response, err := client.Do(hRequest)
	defer response.Body.Close()
	if err != nil {
		return nil, err
	}

	return response, nil
}

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
