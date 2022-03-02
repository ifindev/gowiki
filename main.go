package main

import (
	"gowiki/wiki"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/view/", wiki.ViewHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
