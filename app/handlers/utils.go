package handlers

import (
	"net/http"

	"github.com/a-h/templ"
	"github.com/corentings/kafejo-books/app/views/page"
)

const (
	ContentTypeHTML = "text/html"
)

func Render(w http.ResponseWriter, r *http.Request, _ int, t templ.Component) error {
	w.Header().Set("Content-Type", ContentTypeHTML)
	return t.Render(r.Context(), w)
}

func Redirect(w http.ResponseWriter, _ *http.Request, path string, statusCode int) error {
	w.Header().Set("HX-Redirect", path)
	w.WriteHeader(statusCode)
	return nil
}

func RedirectToErrorPage(w http.ResponseWriter, r *http.Request, errorCode int) error {
	pageToReturn := page.InternalServerError()
	switch errorCode {
	case http.StatusUnauthorized:
		pageToReturn = page.NotAuthorized()
	case http.StatusNotFound:
		pageToReturn = page.NotFound()
	case http.StatusInternalServerError:
		pageToReturn = page.InternalServerError()
	case http.StatusBadRequest:
		pageToReturn = page.BadRequest()
	}

	errorPage := page.ErrorPage("Kafejo Books", pageToReturn)

	return Render(w, r, errorCode, errorPage)
}
