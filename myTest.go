package scrape

type MyTest struct {
	MeinTitel  string   `scrape:"h1, InnerHTML"`
	Paragraphs []string `scrape:"p, InnderHTML"`

	Liste LinkList `scrape:"#my_div`
}

type LinkList struct {
	Links []string `scrape:"a, href"`
}
