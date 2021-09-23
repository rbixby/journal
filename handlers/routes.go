package handlers

import (
	"net/http"
	"text/template"

	"github.com/rbixby/journal/page"
)

const (
	INDEX_TMPL  = "templates/index.html"
	EDIT_TMPL   = "templates/edit.html"
	INDEX_ROUTE = "index"
	EDIT_ROUTE  = "edit"
)

var templates = template.Must(template.ParseFiles(INDEX_TMPL, EDIT_TMPL))

func IndexHandler(w http.ResponseWriter, r *http.Request) {

}

func renderTemplate(w http.ResponseWriter, tmpl string, p *page.Page) {

}
