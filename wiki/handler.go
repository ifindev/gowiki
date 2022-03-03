package wiki

import (
	"net/http"
	"os"
	"text/template"
)

func ViewHandler(w http.ResponseWriter, r *http.Request) {
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

	// If title still not exists, probably user tried to
	// access /save/ route directly without being redirected
	// redirect user to create new wiki
	if title == "" {
		http.Redirect(w, r, "/create/", http.StatusFound)
	} else {
		body := r.FormValue("body")
		p := &Page{Title: title, Body: []byte(body)}
		err := p.Save()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		http.Redirect(w, r, "/view/"+title, http.StatusFound)
	}

}

//TODO : show All wiki as a list with each item is a link to the wiki detail

func renderHtmlFile(html string, w http.ResponseWriter, p *Page) {

	wd, err := os.Getwd()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	file := wd + "/view/" + html

	t, err := template.ParseFiles(file)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	err = t.Execute(w, p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
