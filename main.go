package main

import (
	"fmt"
	"io/ioutil"
	"strings"

	"golang.org/x/net/html"
)

func main() {
	html_code := readFile("./res/simple_example.html")

	tk := html.NewTokenizer(strings.NewReader(html_code))
	for {
		token_type := tk.Next()

		if token_type == html.ErrorToken {
			break
		} else if token_type == html.TextToken {
			fmt.Println(tk.Token().Data)
		} else if token_type == html.StartTagToken || token_type == html.EndTagToken {
			fmt.Printf("<%v>\n", tk.Token().Data)
		}
		fmt.Println(tk.Token().Data)
	}

}

func readFile(path string) string {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}

	return string(data)
}
