package template

import (
	"bytes"
	"os"
	"text/template"
)

// Container for variables passed by user
type TemplateVars struct {
	Values map[string]string
}

var funcMap = template.FuncMap{
	"env": func(s string) (string, error) {
		return os.Getenv(s), nil
	},
}

func Expand(input string) (string, error) {
	t := template.New("request-template").Funcs(funcMap)

	data, err := t.Parse(input)
	if err != nil {
		return "", err
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, data); err != nil {
		return "", err
	}

	return tpl.String(), nil
}
