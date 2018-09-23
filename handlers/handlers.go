package handlers

import (
	"net/http"

	"../models"
	"../utils"
)

// IndexHandler handles main page request
func IndexHandler(w http.ResponseWriter, r *http.Request, title string) {
	pages, err := utils.LoadAllPages()
	if err != nil {
		pages = make([]string, 0)
	}
	utils.RenderMainTemplate(w, "index", pages)
}

// ViewHandler for viewing view
func ViewHandler(w http.ResponseWriter, r *http.Request, title string) {
	p, err := utils.LoadPage(title)
	if err != nil {
		http.Redirect(w, r, "/edit/"+title, http.StatusFound)
		return
	}
	utils.RenderTemplate(w, "view", p)
}

// EditHandler for editing view
func EditHandler(w http.ResponseWriter, r *http.Request, title string) {
	p, err := utils.LoadPage(title)
	if err != nil {
		p = &models.Page{Title: title}
	}
	utils.RenderTemplate(w, "edit", p)
}

// SaveHandler for saving view
func SaveHandler(w http.ResponseWriter, r *http.Request, title string) {
	body := r.FormValue("body")
	p := &models.Page{Title: title, Body: []byte(body)}
	err := p.Save("data")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/view/"+title, http.StatusFound)
}
