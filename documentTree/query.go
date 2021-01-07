package documentTree

import (
	"strings"
)

func (p *HTMLElement) GetElementByID(id string) *HTMLElement {
	for _, c := range p.Children {
		if v, ok := c.Attributes["id"]; ok && v == id {
			return c
		}

		if cf := c.GetElementByID(id); cf != nil {
			return cf
		}
	}

	return nil
}

func (p *HTMLElement) GetElementsByName(name string) []*HTMLElement {
	var found []*HTMLElement
	for _, c := range p.Children {
		if v, ok := c.Attributes["name"]; ok && v == name {
			found = append(found, c)
		}

		found = append(found, c.GetElementsByName(name)...)
	}

	return found
}

func (p *HTMLElement) GetElementsByClassName(className string) []*HTMLElement {
	var found []*HTMLElement
	for _, c := range p.Children {
		if v, ok := c.Attributes["class"]; ok {
			for _, cn := range strings.Split(v, " ") {
				if cn == className {
					found = append(found, c)
					break
				}
			}
		}

		found = append(found, c.GetElementsByClassName(className)...)
	}

	return found
}

func (p *HTMLElement) GetElementsByTagName(tagName string) []*HTMLElement {
	var found []*HTMLElement
	for _, c := range p.Children {
		if c.TagName == tagName {
			found = append(found, c)
		}

		found = append(found, c.GetElementsByTagName(tagName)...)
	}

	return found
}

func (p *HTMLElement) QuerySelector(query string) []*HTMLElement {
	//TODO: Implement

	return nil
}
