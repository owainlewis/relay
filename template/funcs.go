package template

import (
	"bytes"
	"os"
	"text/template"
)

var funcMap = template.FuncMap{
	"env": func(s string) (string, error) {
		return os.Getenv(s), nil
	},
}

func Process(input string, params map[string]string) ([]byte, error) {
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
