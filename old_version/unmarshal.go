package scrape

import (
	"fmt"
	"reflect"
	"strings"
)

func Unmarshal(data []byte, o interface{}) {
	// Generate HTML Tree
	htmlBody := GenerateDocumentTree(data)

	UnmarshalHTMLTree(htmlBody, o)
}

func UnmarshalHTMLTree(root *HTMLElement, o interface{}) {
	// Get interface value and type
	val := reflect.ValueOf(o)
	if val.Kind() != reflect.Ptr {
		panic("Must have pointer")
	}
	val = val.Elem()
	t := val.Type()

	// Iterate fields
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)

		data := processField(field, root)

		if data != nil {
			v := reflect.ValueOf(data)
			if v.Kind() == reflect.Ptr {
				v = v.Elem()
			}
			//fmt.Println(v)
			val.Field(i).Set(v)
		}

	}
}

func processField(field reflect.StructField, root *HTMLElement) interface{} {
	var scrapeTag string
	var ok bool
	// Skip non tagged fields
	if scrapeTag, ok = field.Tag.Lookup("scrape"); !ok {
		return nil
	}

	parts := strings.Split(scrapeTag, ",")
	// Warn about invalid tags
	if len(parts) < 2 {
		fmt.Printf("[WARN] Skipping field '%v': No valid scrape tag. Need query and value\n", field.Name)
		return nil
	}
	query := strings.TrimSpace(parts[0])
	attrib := strings.TrimSpace(parts[1])

	fieldKind := field.Type.Kind().String()
	if fieldKind == "string" {
		// String element
		elements := root.QuerySelector(query)
		if attrib == "InnerHTML" {
			return elements[0].InnerHTML
		}
		return elements[0].Attributes[attrib]
	} else if fieldKind == "slice" && field.Type.Elem().Kind().String() == "string" {
		// String slice element
		elements := root.QuerySelector(query)
		var data []string
		for _, e := range elements {
			if attrib == "InnerHTML" {
				data = append(data, e.InnerHTML)
			} else {
				data = append(data, e.Attributes[attrib])
			}
		}
		return data
	} else if fieldKind == "struct" {
		elements := root.QuerySelector(query)

		obj := reflect.New(field.Type).Interface()

		UnmarshalHTMLTree(elements[0], obj)

		return obj
	} else if fieldKind == "slice" && field.Type.Elem().Kind().String() == "struct" {
		elements := root.QuerySelector(query)

		data := reflect.MakeSlice(field.Type, 0, 0)

		for _, e := range elements {
			obj := reflect.New(field.Type.Elem()).Interface()
			UnmarshalHTMLTree(e, obj)
			data = reflect.Append(data, reflect.ValueOf(obj).Elem())
		}

		return data.Interface()
	}
	return nil
}
