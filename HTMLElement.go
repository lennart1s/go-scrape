package goscrape

import (
	"golang.org/x/net/html"
)

type HTMLElement struct {
	TagName string

	Attributes map[string]string

	InnerHTML string
	Parent    *HTMLElement
	Children  []*HTMLElement
}

func NewHTMLElement() *HTMLElement {
	e := HTMLElement{}

	e.Attributes = make(map[string]string)

	return &e
}

func NewHTMLELementFromToken(token html.Token) *HTMLElement {
	e := NewHTMLElement()

	e.TagName = token.Data

	for _, attr := range token.Attr {
		e.Attributes[attr.Key] = attr.Val
	}

	return e
}

func (e *HTMLElement) AppendChild(c ...*HTMLElement) {
	e.Children = append(e.Children, c...)
}

func (e *HTMLElement) GetValue(name string) string {
	if name == "InnerHTML" {
		return e.InnerHTML
	} else {
		return e.Attributes[name]
	}
}
