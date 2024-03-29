package main

import (
	"fmt"
	scrape "go-scrape"
	"io/ioutil"
)

func main() {
	data, err := ioutil.ReadFile("./res/simple_example.html")
	if err != nil {
		panic(err)
	}
	fmt.Println(resp.StatusCode)

	//fmt.Println(string(data))
	var sd ScrapeData
	//scrape.GenerateDocumentTree(data)
	scrape.Unmarshal(data, &sd)
	//scrape.Unmarshal(data, &sd)

	fmt.Println(sd)

}

type ScrapeData struct {
	Name   string `scrape:"., name"`
	FirstP string `scrape:"button, onclick"`
}
