package scrape

import (
	"regexp"
	"strings"
)

func (p *HTMLElement) GetElementByID(id string) *HTMLElement {
	return p.getElementByID(id, -1)
}
func (p *HTMLElement) getElementByID(id string, depth int) *HTMLElement {
	if depth == 0 {
		return nil
	}

	for _, c := range p.Children {
		if v, ok := c.Attributes["id"]; ok && v == id {
			return c
		}
		if cq := c.getElementByID(id, depth-1); cq != nil {
			return cq
		}
	}

	return nil
}

func (p *HTMLElement) GetElementsByName(name string) []*HTMLElement {
	return p.getElementsByName(name, -1)
}
func (p *HTMLElement) getElementsByName(name string, depth int) []*HTMLElement {
	var found []*HTMLElement
	if depth == 0 {
		return found
	}

	for _, c := range p.Children {
		if v, ok := c.Attributes["name"]; ok && v == name {
			found = append(found, c)
		}

		found = append(found, c.getElementsByName(name, depth-1)...)
	}

	return found
}

func (p *HTMLElement) GetElementsByClassName(className string) []*HTMLElement {
	return p.getElementsByClassName(className, -1)
}
func (p *HTMLElement) getElementsByClassName(className string, depth int) []*HTMLElement {
	var found []*HTMLElement
	if depth == 0 {
		return found
	}

	for _, c := range p.Children {
		if v, ok := c.Attributes["class"]; ok {
			for _, cn := range strings.Split(v, " ") {
				if cn == className {
					found = append(found, c)
					break
				}
			}
		}

		found = append(found, c.getElementsByClassName(className, depth-1)...)
	}

	return found
}

func (p *HTMLElement) GetElementsByTagName(tagName string) []*HTMLElement {
	return p.getElementsByTagName(tagName, -1)
}
func (p *HTMLElement) getElementsByTagName(tagName string, depth int) []*HTMLElement {
	var found []*HTMLElement
	if depth == 0 {
		return found
	}

	for _, c := range p.Children {
		if c.TagName == tagName {
			found = append(found, c)
		}

		found = append(found, c.getElementsByTagName(tagName, depth-1)...)
	}

	return found
}

func (p *HTMLElement) QuerySelector(query string) []*HTMLElement {
	return nil
}

// Tag, id, class, attribute
func ParseLowLevelQuery(query string) *querySelector {
	qs := &querySelector{}
	qs.Attributes = make(map[string]string)

	inQuoteMarks := false
	splitLocations := []int{0}
	for i := 1; i < len(query); i++ {
		if !inQuoteMarks && query[i] == '"' {
			inQuoteMarks = true
			continue
		} else if inQuoteMarks && query[i] != '"' {
			continue
		} else if inQuoteMarks && query[i] == '"' {
			inQuoteMarks = false
			continue
		}
		if strings.ContainsRune("#.[", rune(query[i])) {
			splitLocations = append(splitLocations, i)
		} else if query[i-1] == ']' {
			splitLocations = append(splitLocations, i)
		}
	}
	splitLocations = append(splitLocations, len(query))

	for i := 0; i < len(splitLocations)-1; i++ {
		qs.addSelector(query[splitLocations[i]:splitLocations[i+1]])
	}

	return qs
}

type querySelector struct {
	TagName    string
	ID         string
	ClassNames []string
	Attributes map[string]string
}

func (qs *querySelector) addSelector(s string) {
	if classSelectReggex.MatchString(s) {
		qs.ClassNames = append(qs.ClassNames, strings.TrimLeft(s, "."))
	} else if idSelectReggex.MatchString(s) {
		qs.ID = strings.TrimLeft(s, "#")
	} else if tagSelectReggex.MatchString(s) {
		qs.TagName = s
	} else if attributeReggex.MatchString(s) {
		parts := strings.Split(s[1:len(s)-1], "=")
		if len(parts) == 1 {
			parts = append(parts, "")
		}
		qs.Attributes[parts[0]] = parts[1]
	} else if s != "*" {
		panic("No valid selector: " + s)
	}
}

var (
	// Default Selectors
	classSelectReggex = regexp.MustCompile(`^\.[\w]+$`)
	idSelectReggex    = regexp.MustCompile(`^#[\w]+$`)
	//allSelectReggex   = regexp.MustCompile(`^\*$`)
	tagSelectReggex = regexp.MustCompile(`^[\w]+$`)

	// Advanced Selectors

	// Attribute selector
	attributeReggex = regexp.MustCompile(`^\[\w`)

	// CSS Tag selectors
	cssNthTagReggex = regexp.MustCompile(`nth-child\([0-9]+\)`)

	// https://www.w3schools.com/cssref/css_selectors.asp for list of all selectors
)
