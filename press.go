// press is a utility for applying simple HTML templates to content
package main

import (
	"flag"
	"golang.org/x/net/html"
	"html/template"
	"log"
	"os"
	"strings"
)

var tf string
var cf string

// FIXME This doesn't actually handle more than one template file
func init() {
	flag.StringVar(&tf, "t", "", "path to template file(s)")
	flag.StringVar(&cf, "c", "", "path to content file")
}

func main() {
	flag.Parse()
	if tf == "" || cf == "" {
		log.Fatal("Template or content file not provided")
	}
	type pageData struct {
		Title   string
		Content template.HTML
	}
	content, err := os.ReadFile(cf)
	if err != nil {
		panic(err)
	}
	// Because the user is providing this content, we assume it is safe
	htmlContent := template.HTML(string(content))
	pageTitle := parseTitle(string(content))
	t, err := template.ParseFiles(tf)
	if err != nil {
		log.Fatal(err)
	}
	data := pageData{pageTitle, htmlContent}
	err = t.Execute(os.Stdout, data)
	if err != nil {
		log.Fatal(err)
	}
}

// parseTitle retrieves the string in the first h1 tag
func parseTitle(content string) string {
	var title string
	reader := strings.NewReader(content)
	tokenizer := html.NewTokenizer(reader)

	for {
		token := tokenizer.Next()

		switch token {

		case html.ErrorToken:
			// we've reached the end
			return title

		case html.StartTagToken, html.SelfClosingTagToken:
			tagName, _ := tokenizer.TagName()
			if string(tagName) == "h1" {
				token = tokenizer.Next()
				title = tokenizer.Token().String()
				break
			}
		}
	}

	return title
}
