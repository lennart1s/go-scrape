package goscrape

import (
	"log"
	"reflect"
	"strings"
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
	oType := oElem.Type()

	for i := 0; i < oType.NumField(); i++ {
		field := oType.Field(i)

		data := processField(field, root)

		if data != nil {
			v := reflect.ValueOf(data)
			if v.Kind() == reflect.Ptr {
				v = v.Elem()
			}
			oElem.Field(i).Set(v)
		}
	}
}

func processField(field reflect.StructField, root *HTMLElement) interface{} {
	var scrapeTag string
	var ok bool

	//Skip non tagged fields
	if scrapeTag, ok = field.Tag.Lookup("scrape"); !ok {
		return nil
	}

	parts := strings.Split(scrapeTag, ",")
	// Warn about invalid tags
	if len(parts) < 2 {
		log.Printf("[WARN] Skipping field '%v': No valid scrape tag. Need query and value\n", field.Name)
		return nil
	}
	query := strings.TrimSpace(parts[0])
	attrib := strings.TrimSpace(parts[1])

	fieldKind := field.Type.Kind()
	switch fieldKind.String() {
	case "string":
		elements := root.QuerySelector(query)
		if len(elements) == 0 {
			return nil
		}
		return elements[0].GetValue(attrib)
	case "struct":
		elements := root.QuerySelector(query)
		if len(elements) == 0 {
			return nil
		}
		obj := reflect.New(field.Type).Interface()
		UnmarshalHTMLTree(elements[0], obj)
		return obj
	case "slice":
		elements := root.QuerySelector(query)
		if len(elements) == 0 {
			return nil
		}
		if field.Type.Elem().Kind().String() != "string" && field.Type.Elem().Kind().String() != "struct" {
			log.Printf("Invalid field elem kind '%v' on field '%v'", field.Type.Elem().Kind().String(), field.Name)
			return nil
		}
		data := reflect.MakeSlice(field.Type, 0, 0)
		for _, e := range elements {
			var obj interface{}
			if field.Type.Elem().Kind().String() == "string" {
				obj = e.GetValue(attrib)
				data = reflect.Append(data, reflect.ValueOf(obj))
			} else if field.Type.Elem().Kind().String() == "struct" {
				obj = reflect.New(field.Type.Elem()).Interface()
				UnmarshalHTMLTree(e, obj)
				data = reflect.Append(data, reflect.ValueOf(obj).Elem())
			}
		}

		return data.Interface()
	default:
		log.Printf("Invalid field type '%v' on field '%v'", fieldKind.String(), field.Name)
	}

	return nil
}
