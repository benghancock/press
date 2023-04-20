// press is a utility for applying simple HTML templates to content
package main

import (
	"html/template"
	"log"
	"os"
)

func main() {
	type pageData struct {
		Title   string
		Content template.HTML
	}
	data := pageData{"Hello World!", template.HTML("<p>This is some content</p>")}
	t, err := template.ParseFiles("sample/template1.html")
	if err != nil {
		log.Fatal(err)
	}
	err = t.Execute(os.Stdout, data)
	if err != nil {
		log.Fatal(err)
	}
}
