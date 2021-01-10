package scrape

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

func newHTMLElementFromSTToken(stToken *html.Token) *HTMLElement {
	e := NewHTMLElement()

	e.TagName = stToken.Data

	e.Attributes = make(map[string]string)
	for _, attr := range stToken.Attr {
		e.Attributes[attr.Key] = attr.Val
	}

	return e
}

func newHTMLElementFromSCTToken(sctToken *html.Token) *HTMLElement {
	e := NewHTMLElement()

	e.TagName = sctToken.Data

	return e
}

func (e *HTMLElement) AppendChild(c ...*HTMLElement) {
	e.Children = append(e.Children, c...)
}
