package handlers

import (
	"log/slog"
	"net/http"

	"github.com/corentings/kafejo-books/app/views/page"
)

type BookHandler struct{}

func NewBookHandler() *BookHandler {
	return &BookHandler{}
}

func (p *BookHandler) HandleGetIndex() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		hero := page.Index()

		index := page.IndexPage("kafejo-books", hero)

		if err := Render(w, r, http.StatusOK, index); err != nil {
			slog.ErrorContext(r.Context(), "error rendering index page", slog.String("error", err.Error()))
		}
	}
}
