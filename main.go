package main

import (
	"gowiki/wiki"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/view/", wiki.ViewHandler)
	http.HandleFunc("/edit/", wiki.EditHandler)
	http.HandleFunc("/save/", wiki.SaveHandler)
	http.HandleFunc("/create/", wiki.CreateHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
