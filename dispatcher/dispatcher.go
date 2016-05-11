package dispatcher

import (
	"bytes"
	"fmt"
	"github.com/owainlewis/relay/parser"
	"io/ioutil"
	"net/http"
	"os"
)

func ExtractResponse(resp *http.Response) *parser.Response {
	status := resp.StatusCode
	body, _ := ioutil.ReadAll(resp.Body)
	return &parser.Response{status, string(body)}
}

func ToHttpRequest(request parser.Request) *http.Request {
	payloadBytes := []byte(request.Body)

	method := request.Method
	url := request.Url
	body := bytes.NewBuffer(payloadBytes)

	r, err := http.NewRequest(method, url, body)

	if err != nil {
		panic(err)
	}

	for k, v := range request.Headers {
		r.Header.Set(k, v)
	}

	return r
}

func Run(request parser.Request) (*parser.Response, error) {

	client := &http.Client{}
	response, err := client.Do(ToHttpRequest(request))

	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	return ExtractResponse(response), nil
}

func FromFile(file string) *parser.Response {

	req, err := parser.ParseFile(file)

	if err != nil {
		fmt.Println("Error reading file")
		os.Exit(1)
	}

	response, err := Run(req.Req)

	if err != nil {
		fmt.Println("Error making HTTP request")
		return nil
	}

	return response
}
