// press is a utility for applying simple HTML templates to content
package main

import (
	"flag"
	"html/template"
	"log"
	"os"
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
	t, err := template.ParseFiles(tf)
	if err != nil {
		log.Fatal(err)
	}
	data := pageData{"Hello World!", htmlContent}
	err = t.Execute(os.Stdout, data)
	if err != nil {
		log.Fatal(err)
	}
}

// TODO Implement function for retrieving title from content h1 tag
