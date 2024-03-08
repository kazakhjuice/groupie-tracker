package web

import (
	"net/http"
	"strconv"
	"text/template"
)

func (app *application) ArtistHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/artist" {
		app.ErrorHandler(w, r, http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		app.ErrorHandler(w, r, http.StatusMethodNotAllowed)
		return
	}
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		app.ErrorHandler(w, r, http.StatusBadRequest)
		return
	}
	artist, err := GetArtist(id)
	if err != nil {
		app.ErrorHandler(w, r, http.StatusBadRequest)
		return
	}
	tmpl, err := template.ParseFiles("./ui/html/artist.html")
	if err != nil {
		app.ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, artist)
}
