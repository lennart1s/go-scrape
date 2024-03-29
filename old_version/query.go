package scrape

import (
	"fmt"
	"regexp"
	"strconv"
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

func (p *HTMLElement) getElementsByClassName(className string) []*HTMLElement {
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

		found = append(found, c.getElementsByClassName(className)...)
	}

	return found
}
func (p *HTMLElement) GetElementsByClassName(classNames ...string) []*HTMLElement {
	occurrences := make(map[*HTMLElement]int)

	for _, className := range classNames {
		for _, e := range p.getElementsByClassName(className) {
			occurrences[e]++
		}
	}

	var found []*HTMLElement
	for k, v := range occurrences {
		if v == len(classNames) {
			found = append(found, k)
		}
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
	// Catch 'current-element' query
	if strings.TrimSpace(query) == "." {
		return []*HTMLElement{p}
	}

	var validChildren []*HTMLElement

	queryParts := strings.SplitN(query, ">", 2)
	myQuery := strings.TrimSpace(queryParts[0])

	for _, part := range strings.Split(myQuery, ",") {
		parts := strings.Split(strings.TrimSpace(part), ":")
		part = parts[0]

		var hits []*HTMLElement

		switch {
		case ClassSelectReggex.MatchString(part):
			//hits := p.GetElementsByClassName(strings.Split(strings.TrimLeft(part, "."), ".")...)
			hits = append(hits, p.GetElementsByClassName(strings.Split(strings.TrimLeft(part, "."), ".")...)...)
			//validChildren = append(validChildren, hits...)
			break

		case IDSelectReggex.MatchString(part):
			//validChildren = append(validChildren, p.GetElementByID(strings.TrimLeft(part, "#")))
			hits = append(hits, p.GetElementByID(strings.TrimLeft(part, "#")))
			break

		case AllSelectReggex.MatchString(part):
			//validChildren = append(validChildren, p.Children...)
			hits = append(hits, p.Children...)
			break

		case TagSelectReggex.MatchString(part):
			//validChildren = append(validChildren, p.GetElementsByTagName(part)...)
			hits = append(hits, p.GetElementsByTagName(part)...)
			break

		default:
			fmt.Printf("No valid reggex found: `%v`\n", part)
		}

		if len(parts) > 1 {
			cssTag := parts[1]
			switch {
			case CssNthTagReggex.MatchString(cssTag):
				//fmt.Println("nth tag")
				nStr := strings.SplitN(cssTag, "(", 2)[1]
				nStr = strings.SplitN(nStr, ")", 2)[0]
				n, _ := strconv.Atoi(nStr)
				for i, h := 0, hits[0]; i < len(hits); {
					if n-1 < 0 || n-1 > len(h.Parent.Children) || h.Parent.Children[n-1] != h {
						hits = append(hits[:i], hits[i+1:]...)
					} else {
						i++
					}
				}
				break
			default:
				fmt.Printf("No valid reggex found: `%v`\n", cssTag)
			}
		}

		validChildren = append(validChildren, hits...)
	}

	if len(queryParts) == 1 {
		return validChildren
	}

	var res []*HTMLElement
	for _, vc := range validChildren {
		res = append(res, vc.QuerySelector(queryParts[1])...)
	}

	return res

}

//case CLASS_DESCENDANT_REGGEX.MatchString(part):
//	break
//CLASS_DESCENDANT_REGGEX = regexp.MustCompile(`^\.[a-zA-Z0-9]+(\s\.[a-zA-Z0-9]+)*$`)

var (
	// Default Selectors
	ClassSelectReggex = regexp.MustCompile(`^(\.[\w]+)+$`) // accepts multiple classes
	IDSelectReggex    = regexp.MustCompile(`^#[\w]+$`)
	AllSelectReggex   = regexp.MustCompile(`^\*$`)
	TagSelectReggex   = regexp.MustCompile(`^[\w]+$`)

	// Advanced Selectors

	// Attribute selector

	// CSS Tag selectors
	CssNthTagReggex = regexp.MustCompile(`nth-child\([0-9]+\)`)

	// https://www.w3schools.com/cssref/css_selectors.asp for list of all selectors
)
