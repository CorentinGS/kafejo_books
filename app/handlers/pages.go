package handlers

import (
	"log/slog"
	"net/http"

	"github.com/corentings/kafejo-books/app/views/page"
)

type PageHandler struct{}

func NewPageHandler() *PageHandler {
	return &PageHandler{}
}

func (p *PageHandler) HandleGetIndex() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		hero := page.Index()

		index := page.IndexPage("Kafejo Books", hero)

		if err := Render(w, r, http.StatusOK, index); err != nil {
			slog.ErrorContext(r.Context(), "error rendering index page", slog.String("error", err.Error()))
		}
	}
}
