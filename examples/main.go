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
	BodyName   string   `scrape:"body, name"`
	Paragraphs []string `scrape:"p, id"`
	Ps         []string `scrape:"body>*, id"`
	Test       string   `scrape:"#p2+button, onclick"`
	//DivPs []string `scrape:""`

	/*Name      string             `scrape:"., name"`
	FirstP    string             `scrape:"button, onclick"`
	Test      []string           `scrape:"., InnerHTML"`
	OneData   ScrapteDataChild   `scrape:"., awawd"`
	DataSlice []ScrapteDataChild `scrape:"wa, awda"`*/
}
