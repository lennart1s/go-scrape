package html_elements

import (
	"fmt"
	"strings"

	"golang.org/x/net/html"
)

func GenerateDocumentTree(data []byte) *HTMLElement {
	tk := html.NewTokenizer(strings.NewReader(string(data)))

	var currentParent *HTMLElement
	for {
		tk.Next()
		token := tk.Token()
		if token.Type == html.ErrorToken {
			break
		} else if len(strings.TrimSpace(token.Data)) == 0 {
			continue
		} else if currentParent == nil && !(token.Type == html.StartTagToken && token.Data == "body") {
			continue
		} else if currentParent != nil && token.Type == html.EndTagToken && token.Data == "body" {
			break
		}

		// Process valid tokens
		if token.Type == html.StartTagToken {
			newElement := newHTMLElementFromSTToken(&token)
			if currentParent != nil {
				currentParent.AppendChild(newElement)
				newElement.Parent = currentParent
			}
			currentParent = newElement
		} else if token.Type == html.EndTagToken {
			currentParent = currentParent.Parent
		} else if token.Type == html.SelfClosingTagToken {
			newElement := newHTMLElementFromSCTToken(&token)
			currentParent.AppendChild(newElement)
			newElement.Parent = currentParent
		} else if token.Type == html.TextToken {
			currentParent.InnerHTML += token.Data
		} else if token.Type == html.ErrorToken {
			fmt.Printf("[WARN] Error during tokenization. %v\n", token)
		}

	}

	return currentParent //currentParent is not body element
}
