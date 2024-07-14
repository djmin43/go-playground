package main

import (
	"errors"
	"fmt"
	"github.com/djmin43/todo-list/internal/models"
	"html/template"
	"net/http"
	"strconv"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	files := []string{
		"./ui/html/pages/base.tmpl",
		"./ui/html/pages/home.tmpl",
	}

	ts, err := template.ParseFiles(files...)

	if err != nil {
		app.serverError(w, err)
		return
	}

	err = ts.ExecuteTemplate(w, "base", nil)

	if err != nil {
		app.serverError(w, err)
		return
	}

}

func (app *application) todoView(w http.ResponseWriter, r *http.Request) {

	id, err := strconv.Atoi(r.URL.Query().Get("id"))

	if err != nil || id < 1 {
		app.notFoundError(w)
		return
	}

	todo, err := app.todos.Get(id)

	if err != nil {
		if errors.Is(err, models.ErrNoRecord) {
			app.notFoundError(w)
		} else {
			app.serverError(w, err)
		}
		return
	}

	fmt.Fprintf(w, "%+v", todo)
}

func (app *application) todoCreate(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.Header().Set("Allow", "POST")
		app.methodNotAllowedError(w)
		return
	}

	title := "first todo"
	content := "do something cool"
	interval := 7

	id, err := app.todos.Insert(title, content, interval)
	if err != nil {
		app.serverError(w, err)
	}

	http.Redirect(w, r, fmt.Sprintf("/view?id=%d", id), http.StatusSeeOther)
}
