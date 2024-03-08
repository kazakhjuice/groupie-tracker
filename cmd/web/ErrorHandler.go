package web

import (
	"net/http"
	"text/template"
)

func (app *application) ErrorHandler(w http.ResponseWriter, r *http.Request, statusCode int) {
	w.WriteHeader(statusCode)
	t, err := template.ParseFiles("./ui/html/error.html")
	if err != nil {
		w.WriteHeader(statusCode)
		http.Error(w, http.StatusText(statusCode), statusCode)
		return
	}

	errors := struct {
		StatusCode int
		StatusMsg  string
	}{
		StatusCode: statusCode,
		StatusMsg:  http.StatusText(statusCode),
	}
	err = t.Execute(w, errors)
	if err != nil {
		w.WriteHeader(statusCode)
		http.Error(w, http.StatusText(statusCode), statusCode)
		return
	}
}
