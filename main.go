package main

import (
	"gowiki/wiki"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/view/", wiki.ViewHandler)
	http.HandleFunc("/edit/", wiki.EditHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
