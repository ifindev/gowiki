package wiki

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"text/template"
)

func ViewHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL.Path)
	title := r.URL.Path[len("/view/"):]
	p, err := LoadPage(title)
	if err != nil {
		p = &Page{Title: title}
	}

	file, err := getHtmlFile("/view/view.html")
	if err != nil {
		log.Fatal(err)
	}

	loadHtmlFile(file, w, p)
}

func EditHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/edit/"):]
	p, err := LoadPage(title)
	if err != nil {
		p = &Page{Title: title}
	}

	file, err := getHtmlFile("/view/edit.html")
	if err != nil {
		log.Fatal(err)
	}

	loadHtmlFile(file, w, p)
}

func getHtmlFile(relPath string) (string, error) {
	var file string
	wd, err := os.Getwd()
	if err != nil {
		return file, err
	}

	file = wd + relPath
	return file, nil
}

func loadHtmlFile(file string, w http.ResponseWriter, p *Page) {
	t, err := template.ParseFiles(file)
	if err != nil {
		log.Fatal(err)
	}
	t.Execute(w, p)
}
