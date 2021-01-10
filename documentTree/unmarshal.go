package documentTree

import (
	"fmt"
	"reflect"
	"strings"
)

// Convert html data to given object structure
func Unmarshal(data []byte, o interface{}) {
	tree := GenerateDocumentTree(data)

	val := reflect.ValueOf(o)
	if val.Kind() != reflect.Ptr {
		panic("Must have pointer")
	}
	val = val.Elem()
	t := val.Type()
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)

		//fmt.Printf("%v(%v) with tag: `%v`\n", field.Name, field.Type, field.Tag)
		tag := field.Tag
		fmt.Println(field.Tag)

		if scrapeTag, ok := tag.Lookup("scrape"); ok {
			// Relevant objects
			fmt.Println(field.Type.Kind())
			if field.Type.Kind().String() == "struct" {
				// Do recursion here

			} else if field.Type.Kind().String() == "string" {
				parts := strings.Split(scrapeTag, ",")
				query := strings.TrimSpace(parts[0])
				attrib := strings.TrimSpace(parts[1])

				elements := tree.QuerySelector(query)
				var data string
				if attrib == "InnerHTML" {
					data = elements[0].InnerHTML
				} else {
					data = elements[0].Attributes[attrib]
				}
				val.Field(i).SetString(data)
			} else if field.Type.Kind().String() == "slice" {
				//fmt.Println(field.Type.Elem())

				parts := strings.Split(scrapeTag, ",")
				query := strings.TrimSpace(parts[0])
				attrib := strings.TrimSpace(parts[1])

				elements := tree.QuerySelector(query)
				var data []string
				if attrib == "InnerHTML" {
					for _, e := range elements {
						data = append(data, e.InnerHTML)
					}
				} else {
					for _, e := range elements {
						data = append(data, e.Attributes[attrib])
					}
				}
				val.Field(i).Set(reflect.ValueOf(data))
			}

			fmt.Println(scrapeTag)
			htmlElem := tree.QuerySelector(scrapeTag)
			fmt.Printf("Foud elements for '%v': %v\n", field.Name, htmlElem)
		}

		/* if field.Type.String() == "string" {
			val.Field(i).SetString("tshoyo")
		} else if field.Type.String() == "int" {
			val.Field(i).SetInt(420)
		} */
	}

}
