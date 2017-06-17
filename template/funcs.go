package template

import (
	"bytes"
	"os"
	"text/template"
)

func Expand(input string) (string, error) {
	fm := template.FuncMap{
		"env": func(s string) (string, error) {
			return os.Getenv(s), nil
		},
	}

	t := template.New("").Funcs(fm)

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
