package main

import (
	"fmt"
	"go-scrape/documentTree"
	"io/ioutil"
)

func main() {
	html_code := readFile("./res/simple_example.html")

	body := documentTree.GenerateDocumentTree(html_code)
	fmt.Println(*body)

	body.QuerySelector("#my_div > .megaCoolerLink .mcl2")
	//fmt.Println(body.GetElementsByTagName("button")[0])
}

func readFile(path string) []byte {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}

	return data
}
