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

	//fmt.Println(string(data))
	var sd ScrapeData
	//scrape.GenerateDocumentTree(data)
	scrape.Unmarshal(data, &sd)
	//scrape.Unmarshal(data, &sd)

	fmt.Println(sd)

}

type ScrapeData struct {
	Name      string             `scrape:"., name"`
	FirstP    string             `scrape:"button, onclick"`
	Test      []string           `scrape:"., InnerHTML"`
	OneData   ScrapteDataChild   `scrape:"., awawd"`
	DataSlice []ScrapteDataChild `scrape:"wa, awda"`
}

type ScrapteDataChild struct {
	Data string `scrape:"., onclick"`
}
