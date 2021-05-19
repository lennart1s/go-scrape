package goscrape

import (
	"log"
	"regexp"
	"strings"
)

func (p *HTMLElement) QuerySelector(query string) []*HTMLElement {
	query = strings.TrimSpace(query)
	query = regexp.MustCompile(`\s+`).ReplaceAllString(query, " ")
	query = regexp.MustCompile(`\s*>\s*`).ReplaceAllString(query, ">")
	query = regexp.MustCompile(`\s*\+\s*`).ReplaceAllString(query, "+")
	query = regexp.MustCompile(`\s*~\s*`).ReplaceAllString(query, "~")

	if query == "." || query == "" {
		return []*HTMLElement{p}
	}

	var found []*HTMLElement

	queryParts := strings.Split(query, ",")
	if len(queryParts) > 1 {
		for _, qp := range queryParts {
			found = append(found, p.QuerySelector(qp)...)
		}
		return found
	}

	// handle concatinations

	if i := strings.IndexAny(query, " >+~"); i != -1 {
		if i != 0 {
			for _, first := range p.lowLevelSelector(query[:i]) {
				found = append(found, first.QuerySelector(query[i:])...)
			}
			return found
		}

		switch {
		case query[i] == '>':
			query = query[1:]
			newI := strings.IndexAny(query, " >+~")
			if newI == -1 {
				newI = len(query)
			}
			for _, d := range p.getDescendant(query[:newI]) {
				found = append(found, d.QuerySelector(query[newI:])...)
			}
			return found
		}

		/*for i := 0; i < len(query)-1; i++ {
			if strings.ContainsRune(" >+~", rune(query[i])) {
				switch {
				case query[i] == ' ' && !strings.ContainsRune(" >+~", rune(query[i+1])):
					for _, first := range p.QuerySelector(query[:i]) {
						found = append(found, first.QuerySelector(query[i+1:])...)
					}
					return found

				case query[i] == '>':
					for _, first := range p.QuerySelector(query[:i]) {
						found = append(found)
					}
					return found
				case query[i] == '+':
					break
				case query[i] == '~':
					break
				}
			}
		}*/
	} else {
		return p.lowLevelSelector(query)
	}

	return nil
}

func (p *HTMLElement) lowLevelSelector(query string) []*HTMLElement {
	switch {
	case IDSelectReggex.MatchString(query):
		return []*HTMLElement{p.GetElementByID(strings.TrimLeft(query, "#"))}

	case ClassSelectReggex.MatchString(query):
		return p.GetElementsByClassName(strings.Split(strings.TrimLeft(query, "."), ".")...)

	case TagSelectReggex.MatchString(query):
		return p.GetElementsByTagName(query)

	case AllSelectReggex.MatchString(query):
		return p.getAllDescendants()

	default:
		log.Println("[WARN] Unknown css query: " + query)
	}

	return nil
}

func (p *HTMLElement) getDescendant(query string) []*HTMLElement {
	childMap := make(map[*HTMLElement][]*HTMLElement)
	for _, c := range p.Children {
		childMap[c] = c.Children
		c.Children = []*HTMLElement{}
	}
	found := p.lowLevelSelector(query)
	for _, c := range p.Children {
		c.Children = childMap[c]
	}

	return found
}

func (p *HTMLElement) getAllDescendants() []*HTMLElement {
	var found []*HTMLElement
	for _, c := range p.Children {
		found = append(found, c)
		found = append(found, c.getAllDescendants()...)
	}
	return found
}

//*******************************
// LOW LEVEL SELECTORS
//*******************************

func (p *HTMLElement) GetElementByID(id string) *HTMLElement {
	for _, c := range p.Children {
		if v, ok := c.Attributes["id"]; ok && v == id {
			return c
		}

		if cc := c.GetElementByID(id); cc != nil {
			return cc
		}
	}

	return nil
}

/*func (p *HTMLElement) GetElementsByName(name string) []*HTMLElement {
	var found []*HTMLElement

	for _, c := range p.Children {
		if v, ok := c.Attributes["name"]; ok && v == name {
			found = append(found, c)
		}

		found = append(found, c.GetElementsByName(name)...)
	}

	return found
}*/

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

var (
	IDSelectReggex    = regexp.MustCompile(`^#[\w]+$`)
	ClassSelectReggex = regexp.MustCompile(`^(\.[\w]+)+$`) // accepts multiple classes
	TagSelectReggex   = regexp.MustCompile(`^[\w]+$`)
	AllSelectReggex   = regexp.MustCompile(`^\*$`)
)
