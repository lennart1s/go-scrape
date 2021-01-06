package main

import (
	"fmt"
	"go-scrape/html_elements"
	"io/ioutil"
)

func main() {
	html_code := readFile("./res/simple_example.html")

	body := html_elements.GenerateDocumentTree(html_code)
	fmt.Println(*body)
}

func readFile(path string) []byte {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}

	return data
}
