package template

import (
	"bytes"
	"encoding/base64"
	"os"
	"text/template"
)

func envFn(envVar string) (string, error) {
	return os.Getenv(envVar), nil
}

func basicAuthFn(username string, password string) (string, error) {
	auth := username + ":" + password
	return base64.StdEncoding.EncodeToString([]byte(auth)), nil
}

func base64Fn(input string) (string, error) {
	return base64.StdEncoding.EncodeToString([]byte(input)), nil
}

var funcMap = template.FuncMap{
	"env":          envFn,
	"basic":        basicAuthFn,
	"base64encode": base64Fn,
}

// Expand will expand out any user defined functions in a template
func Expand(input string, params map[string]string) ([]byte, error) {
	tmpl := template.New("request-template").Funcs(funcMap)

	t, err := tmpl.Parse(input)
	if err != nil {
		return []byte{}, err
	}

	var buffer bytes.Buffer
	if err := t.Execute(&buffer, params); err != nil {
		return []byte{}, err
	}

	return buffer.Bytes(), nil
}
