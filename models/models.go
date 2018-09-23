package models

import "io/ioutil"

// Page represents wiki pages
type Page struct {
	Title string
	Body  []byte
}

// Save method saving wiki pages in file system
func (p *Page) Save(path string) error {
	filename := path + "/" + p.Title + ".txt"
	return ioutil.WriteFile(filename, p.Body, 0600)
}
