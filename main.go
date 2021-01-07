package main

import (
	"fmt"
	"go-scrape/documentTree"
	"io/ioutil"
)

func main() {
	html_code := readFile("./res/simple_example.html")

	body := documentTree.GenerateDocumentTree(html_code)
	//fmt.Println(*body)
	//fmt.Println(body.GetElementsByClassName("class2"))

	res := body.QuerySelector("#my_div > button")
	fmt.Println(len(res))
	for _, r := range res {
		fmt.Println(r)
	}
	//fmt.Println(body.GetElementsByTagName("button")[0])
}

func readFile(path string) []byte {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}

	return data
}
