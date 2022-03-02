package wiki

import (
	"os"
)

type Page struct {
	Title string
	Body  []byte
}

func (p *Page) Save() error {
	path, err := os.Getwd()
	if err != nil {
		return err
	}

	filename := path + "/assets/" + p.Title + ".txt"
	return os.WriteFile(filename, p.Body, 0600)
}

func LoadPage(title string) (*Page, error) {
	path, err := os.Getwd()
	filename := path + "/assets/" + title + ".txt"

	if err != nil {
		return nil, err
	}

	body, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}
