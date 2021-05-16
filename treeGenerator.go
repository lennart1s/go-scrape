package goscrape

import (
	"log"
	"strings"

	"golang.org/x/net/html"
)

func GenerateDocumentTree(data []byte) []*HTMLElement {
	tk := html.NewTokenizer(strings.NewReader((string(data))))

	var roots []*HTMLElement
	var currentParent *HTMLElement
	for {
		tk.Next()
		token := tk.Token()
		if token.Type == html.ErrorToken {
			log.Println("[WARN] Found error token: " + token.String())
			break
		} else if len(strings.TrimSpace(token.Data)) == 0 {
			continue
		}

		if token.Type == html.StartTagToken {
			newElement := NewHTMLELementFromToken(token)
			if currentParent == nil {
				currentParent = newElement
				roots = append(roots, newElement)
			} else {
				currentParent.AppendChild(newElement)
				newElement.Parent = currentParent
				currentParent = newElement
			}
		} else if token.Type == html.EndTagToken {
			if currentParent == nil {
				log.Println("[WARN] Unexpected EndTagToken: " + token.String())
			} else {
				currentParent = currentParent.Parent
			}
		} else if token.Type == html.SelfClosingTagToken {
			newElement := NewHTMLELementFromToken(token)
			if currentParent == nil {
				roots = append(roots, newElement)
			} else {
				currentParent.AppendChild(newElement)
			}
		} else if token.Type == html.TextToken {
			//log.Println("[INFO] Found text token. Text tokens not yet implemented: " + token.String())
		}

	}

	return roots
}
