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
	fmt.Fprintf(w, "<h1>%s</h1><div>%s</div>", p.Title, p.Body)
}

func EditHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/edit/"):]
	p, err := LoadPage(title)
	if err != nil {
		p = &Page{Title: title}
	}

	wd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	file := wd + "/view/edit.html"
	t, err := template.ParseFiles(file)
	if err != nil {
		log.Fatal(err)
	}
	t.Execute(w, p)
}
