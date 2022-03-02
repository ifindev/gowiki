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

	loadHtmlFile("view.html", w, p)
}

func EditHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/edit/"):]
	p, err := LoadPage(title)
	if err != nil {
		p = &Page{Title: title}
	}

	loadHtmlFile("edit.html", w, p)
}

func loadHtmlFile(html string, w http.ResponseWriter, p *Page) {

	wd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	file := wd + "/view/" + html

	t, err := template.ParseFiles(file)
	if err != nil {
		log.Fatal(err)
	}

	t.Execute(w, p)
}
