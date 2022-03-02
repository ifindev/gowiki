package main

import (
	"fmt"
	"gowiki/wiki"
)

func main() {
	p1 := &wiki.Page{Title: "TestPage", Body: []byte("This is a sample page.")}
	p1.Save()
	p2, err := wiki.LoadPage("TestPage")
	if err != nil {
		fmt.Println("Error loading the file")
	}
	fmt.Println(string(p2.Body))

}
