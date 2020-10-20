package models

import (
	"bytes"
	htmlTpl "html/template"
	textTpl "text/template"

	"github.com/BarTar213/notificator/utils"
)

type Template struct {
	ID      int    `json:"id"`
	Name    string `json:"name" binding:"required"`
	Title   string `json:"title"`
	Message string `json:"message" binding:"required"`
	HTML    bool   `json:"html"`
}

func (t *Template) Parse(data map[string]string) (string, error) {
	if t.HTML {
		return t.parseHTML(data)
	}
	return t.parseText(data)
}

func (t *Template) parseHTML(data map[string]string) (string, error) {
	tpl, err := htmlTpl.New(t.Name).Parse(t.Message)
	if err != nil {
		return utils.EmptyStr, err
	}

	var tplBytes bytes.Buffer
	err = tpl.Execute(&tplBytes, data)
	if err != nil {
		return utils.EmptyStr, err
	}

	return tplBytes.String(), nil
}

func (t *Template) parseText(data map[string]string) (string, error) {
	tpl, err := textTpl.New(t.Name).Parse(t.Message)
	if err != nil {
		return utils.EmptyStr, err
	}

	var tplBytes bytes.Buffer
	err = tpl.Execute(&tplBytes, data)
	if err != nil {
		return utils.EmptyStr, err
	}

	return tplBytes.String(), nil
}
