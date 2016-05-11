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

func Run(request parser.Request) (*parser.Response, error) {
	payloadBytes := []byte(request.Body)

	method := request.Method
	url := request.Url
	body := bytes.NewBuffer(payloadBytes)

	r, _ := http.NewRequest(method, url, body)
	for k, v := range request.Headers {
		r.Header.Set(k, v)
	}

	client := &http.Client{}
	response, err := client.Do(r)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	return ExtractResponse(response), nil
}

func FromFile(file string) {

	req, err := parser.ParseFile(file)

	if err != nil {
		fmt.Println("Error reading file")
		os.Exit(1)
	}

	response, err := Run(req.Req)

	if err != nil {
		fmt.Println("Error making HTTP request")
		return
	}

	fmt.Println(response.Body)
}
