package main

import (
	"fmt"
	"go-scrape"
	"io/ioutil"
)

func main() {
	html_code := readFile("./res/simple_example.html")

	//body := documentTree.GenerateDocumentTree(html_code)
	//fmt.Println(*body)
	//fmt.Println(body.GetElementsByClassName("class2"))

	//res := body.QuerySelector("#my_div > button, p")
	/* fmt.Println(len(res))
	for _, r := range res {
		fmt.Println(r)
	} */
	//fmt.Println(body.GetElementsByTagName("button")[0])

	var mt scrape.MyTest

	scrape.Unmarshal(html_code, &mt)
	fmt.Println("-----")
	fmt.Println(mt)
}

func readFile(path string) []byte {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}

	return data
}
