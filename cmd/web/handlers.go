package main

import (
	"fmt"
	"net/http"
	"strconv"
	"text/template"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Server", "Go")

	files := []string{
		"./ui/html/base.tmpl",
		"./ui/html/partials/nav.tmpl",
		"./ui/html/pages/home.tmpl",
	}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		app.serverError(w, r, err)
	}

	err = ts.ExecuteTemplate(w, "base", nil)
	if err != nil {
		app.serverError(w, r, err)
	}
}

func (app *application) snippetView(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))

	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}

	fmt.Fprint(w, "Display a specific id: %d", id)
}

func (app *application) snippetCreate(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Display a form for creating new snippets..."))
}

func (app *application) snippetCreatePost(w http.ResponseWriter, r *http.Request) {
  var (
    title = "A title"
    content = "Some content"
    expires = 7
  )

  id, err := app.snippets.Insert(title, content, expires)
  if err != nil {
    app.serverError(w, r, err)
  }

  http.Redirect(w, r, fmt.Sprintf("/view/%d", id), http.StatusSeeOther)
}
