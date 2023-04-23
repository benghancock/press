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
	pageTitle := getFirstElemValue(string(content), "h1")
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

// getFirstElemValue reads the HTML content provided and retrieves
// the value in the first tag matching the given name
func getFirstElemValue(content string, name string) string {
	var val string
	reader := strings.NewReader(content)
	tokenizer := html.NewTokenizer(reader)

	for {
		token := tokenizer.Next()

		switch token {

		case html.ErrorToken:
			// we've reached the end
			return val

		case html.StartTagToken, html.SelfClosingTagToken:
			tagName, _ := tokenizer.TagName()
			if string(tagName) == name {
				token = tokenizer.Next()
				val = tokenizer.Token().String()
				break
			}
		}
	}

	return val
}
