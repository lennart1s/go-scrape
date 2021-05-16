package goscrape

import (
	"reflect"
)

func Unmarshal(data []byte, o interface{}) {
	roots := GenerateDocumentTree(data)

	for _, r := range roots {
		if r.TagName == "html" { // eventually do for all roots
			UnmarshalHTMLTree(r, o)
		}
	}
}

func UnmarshalHTMLTree(root *HTMLElement, o interface{}) {
	val := reflect.ValueOf(o)
	if val.Kind() != reflect.Ptr {
		panic("Must have pointer")
	}

	oElem := val.Elem()
	oType := val.Type()

	for i := 0; i < oType.NumField(); i++ {
		field := oType.Field(i)
	}
}
