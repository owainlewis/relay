package template

import (
	"bytes"
	"encoding/base64"
	"os"
	"text/template"
)

var funcMap = template.FuncMap{
	"env": func(envVar string) (string, error) {
		return os.Getenv(envVar), nil
	},
	"basic": func(username string, password string) (string, error) {
		auth := username + ":" + password
		return base64.StdEncoding.EncodeToString([]byte(auth)), nil
	},
	"base64encode": func(input string) (string, error) {
		return base64.StdEncoding.EncodeToString([]byte(input)), nil
	},
}

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
