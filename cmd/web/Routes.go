package web

import (
	"net/http"
)

func (app *application) Routes() http.Handler {
	mux := http.NewServeMux()
	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.HandleFunc("/", app.MainHandler)
	mux.HandleFunc("/artist", app.ArtistHandler)
	mux.Handle("/static/", http.StripPrefix("/static/", fileServer))
	return app.logRequest(mux)
}
