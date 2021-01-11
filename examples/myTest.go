package main

type MyTest struct {
	Name string `scrape:"., name"`

	MeinTitel  string   `scrape:"h1, InnerHTML"`
	Paragraphs []string `scrape:"p, InnerHTML"`

	//Liste  LinkList   `scrape:"#my_div, ."`
	Listen []LinkList `scrape:"#my_div, ."`
}

type LinkList struct {
	Links []string `scrape:"a, href"`
	Link  string   `scrape:"a, href"`
}
