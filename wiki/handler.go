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

	file := "/view/view.html"
	loadHtmlFile(file, w, p)
}

func EditHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/edit/"):]
	p, err := LoadPage(title)
	if err != nil {
		p = &Page{Title: title}
	}

	file := "/view/edit.html"
	loadHtmlFile(file, w, p)
}

func loadHtmlFile(relPath string, w http.ResponseWriter, p *Page) {

	wd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	file := wd + relPath

	t, err := template.ParseFiles(file)
	if err != nil {
		log.Fatal(err)
	}

	t.Execute(w, p)
}
