package main

import (
	"fmt"
	"go-scrape"
	"net/http"
)

func main() {
	//resp, err := http.Get("https://www.dogloversdigest.com/alabama-rescue-shelters-and-organizations/")
	resp, err := http.Get("https://www.stephansandbothe.de/")
	if err != nil {
		panic(err)
	}

	var data []byte
	buffer := make([]byte, 1024)
	for n, err := resp.Body.Read(buffer); n > 0 && err != nil; n, err = resp.Body.Read(buffer) {
		data = append(data, buffer[:n]...)
	}

	//fmt.Println(string(data))
	var sd ScrapeData
	scrape.Unmarshal(data, &sd)

	fmt.Println(sd)

}

type ScrapeData struct {
	SD ShelterData `scrape:"., ."`
}

type ShelterData struct {
	Name string `scrape:"*:nth-child(2), InnerHTML"`
}
