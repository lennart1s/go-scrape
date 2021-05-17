package goscrape

import "golang.org/x/net/html"

type HTMLElement struct {
	TagName string

	Attributes map[string]string

	isTextToken   bool
	isSelfClosing bool

	Parent   *HTMLElement
	Children []*HTMLElement
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

func (e *HTMLElement) InnerHTML() string {
	innerHTML := ""
	for _, c := range e.Children {
		if c.isSelfClosing {
			innerHTML += "</" + c.TagName + ">"
		} else if c.isTextToken {
			innerHTML += c.TagName
		} else {
			innerHTML += "<" + c.TagName
			for k, v := range c.Attributes {
				innerHTML += " " + k + `="` + v + `"`
			}
			innerHTML += ">"
			innerHTML += c.InnerHTML()
			innerHTML += "</" + c.TagName + ">"
		}
	}

	return innerHTML
}

func (e *HTMLElement) GetValue(name string) string {
	if name == "innerHTML" {
		return e.InnerHTML()
	}
	return e.Attributes[name]
}
