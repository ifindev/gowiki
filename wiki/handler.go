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
		http.Redirect(w, r, "/create/", http.StatusFound)
	}

	renderHtmlFile("view.html", w, p)
}

func EditHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/edit/"):]
	p, err := LoadPage(title)
	if err != nil {
		p = &Page{Title: title}
	}

	renderHtmlFile("edit.html", w, p)
}

func CreateHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/create/"):]
	p, err := LoadPage(title)
	if err != nil {
		p = &Page{Title: title}
	}

	renderHtmlFile("create.html", w, p)
}

func SaveHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/save/"):]

	// Creating new wiki
	if title == "" {
		title = r.FormValue("title")
	}

	body := r.FormValue("body")
	p := &Page{Title: title, Body: []byte(body)}
	p.Save()
	http.Redirect(w, r, "/view/"+title, http.StatusFound)
}

//TODO : show All wiki as a list with each item is a link to the wiki detail

func renderHtmlFile(html string, w http.ResponseWriter, p *Page) {

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
