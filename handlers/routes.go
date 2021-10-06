package handlers

import (
	"fmt"
	"net/http"
	"text/template"

	"github.com/google/uuid"
	"github.com/rbixby/journal/entry"
)

const (
	INDEX_TMPL  = "templates/index.html"
	EDIT_TMPL   = "templates/edit.html"
	INDEX_ROUTE = "index.html"
	EDIT_ROUTE  = "edit.html"
)

var templates = template.Must(template.ParseFiles(INDEX_TMPL, EDIT_TMPL))

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	println("In IndexHandler!")
	entries, err := entry.LoadEntries()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	if len(*entries) == 0 {
		// should go to an edit page, but I think I'll post a message saying
		// no entries, with a link to the add page.
		fmt.Printf("No data!")
	} else {
		fmt.Println("Got Data!")
	}
	p := &entry.Page{Name: "Roger's Page", Entries: *entries}
	renderTemplate(w, INDEX_ROUTE, p)
}

func EditHandler(w http.ResponseWriter, r *http.Request) {
	entryID := r.URL.Path[len("/edit/"):]
	fmt.Printf("The Entry ID : %s", entryID)

	renderTemplate(w, EDIT_ROUTE, nil)
}

func SaveHandler(w http.ResponseWriter, r *http.Request) {
	println("In SaveHandler")
	fmt.Println(r)
	title := r.PostFormValue("title")
	date := r.PostFormValue("date")
	body := r.PostFormValue("entry-text")

	fmt.Printf("The title: %s \nThe date: %s \n The body: %s", title, date, body)
	id := uuid.New().String()

	e := entry.NewEntry(id, title, body, date, nil)

	err := e.SaveEntry()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/", http.StatusFound)
}

func renderTemplate(w http.ResponseWriter, tmpl string, p *entry.Page) {
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
