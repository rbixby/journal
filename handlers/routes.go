package handlers

import (
	"net/http"
	"text/template"

	"github.com/rbixby/journal/page"
)

const (
	INDEX_TMPL  = "templates/index.html"
	EDIT_TMPL   = "templates/edit.html"
	INDEX_ROUTE = "index.html"
	EDIT_ROUTE  = "edit.html"
)

var templates = template.Must(template.ParseFiles(INDEX_TMPL))

func IndexHandler(w http.ResponseWriter, r *http.Request) {

	renderTemplate(w, INDEX_ROUTE, nil)
}

func renderTemplate(w http.ResponseWriter, tmpl string, p *page.Page) {
	err := templates.ExecuteTemplate(w, tmpl, p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// func loadPage(title string) (*page.Page, error) {
// filename := title + ".txt"
// body, err := ioutil.ReadFile(filename)

// if err != nil {
// 	return nil, err
// }
// testPage = &page.Page{}
// return &page.Page{Title: title, Body: body}, nil
// }
