package utils

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"

	"../models"
)

var validPath = regexp.MustCompile("^/(edit|save|view)/([a-zA-Z0-9]+)$")

// MakeHandler making handlers
func MakeHandler(fn func(http.ResponseWriter, *http.Request, string)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/" {
			fn(w, r, "")
			return
		}
		m := validPath.FindStringSubmatch(r.URL.Path)
		if m == nil {
			http.NotFound(w, r)
			return
		}
		fn(w, r, m[2])
	}
}

// LoadPage loads pages
func LoadPage(title string) (*models.Page, error) {
	filename := "data/" + title + ".txt"
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &models.Page{Title: title, Body: body}, nil
}

// LoadAllPages loading all pages
func LoadAllPages() ([]string, error) {
	var pages []string
	files, err := ioutil.ReadDir("data")
	if err != nil {
		return nil, err
	}
	for _, v := range files {
		if v != nil {
			name := strings.Split(v.Name(), ".")
			pages = append(pages, name[0])
		}
	}
	return pages, nil
}

// RenderMainTemplate rendering templates
func RenderMainTemplate(w http.ResponseWriter, tmpl string, p []string) {
	t, err := template.ParseFiles("views/index.html")
	if err != nil {
		fmt.Println(err)
	}
	t.Execute(w, p)
}

// RenderTemplate rendering templates
func RenderTemplate(w http.ResponseWriter, tmpl string, p *models.Page) {
	t, _ := template.ParseFiles("views/" + tmpl + ".html")
	t.Execute(w, p)
}
