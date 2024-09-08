package handlers

import (
	"log/slog"
	"net/http"

	"github.com/a-h/templ"
	"github.com/corentings/kafejo-books/app/views/components"
	"github.com/corentings/kafejo-books/app/views/page"
	"github.com/go-chi/chi/v5"
)

type BookHandler struct{}

func NewBookHandler() *BookHandler {
	return &BookHandler{}
}

func (p *BookHandler) HandleGetIndex() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		hero := page.BookIndex()

		index := page.BookPage("kafejo-books", hero)

		if err := Render(w, r, http.StatusOK, index); err != nil {
			slog.ErrorContext(r.Context(), "error rendering index page", slog.String("error", err.Error()))
		}
	}
}

func (p *BookHandler) HandleGetBooks() http.HandlerFunc {
	books := map[string]string{
		"1": "Welcome to Kafejo Books",
		"2": "Corentin GS's Manual",
	}
	return func(w http.ResponseWriter, r *http.Request) {
		bookList := components.BooksList(books)

		if err := Render(w, r, http.StatusOK, bookList); err != nil {
			slog.ErrorContext(r.Context(), "error rendering books list", slog.String("error", err.Error()))
		}
	}
}

func (p *BookHandler) HandleGetLoremIpsum() http.HandlerFunc {
	pages := map[string]templ.Component{
		"1": page.LoremIpsumChapter1(),
		"2": page.LoremIpsumChapter2(),
	}
	return func(w http.ResponseWriter, r *http.Request) {
		// Get the page from request url
		page := chi.URLParam(r, "page")

		// Check if the page exists
		if _, ok := pages[page]; !ok {
			http.NotFound(w, r)
			return
		}

		// Render the page
		if err := Render(w, r, http.StatusOK, pages[page]); err != nil {
			slog.ErrorContext(r.Context(), "error rendering lorem ipsum page", slog.String("error", err.Error()))
		}
	}
}

func (p *BookHandler) HandleGetBook() http.HandlerFunc {
	books := map[string]templ.Component{
		"1": page.BookIndex(),
		"2": page.Manual(),
	}

	return func(w http.ResponseWriter, r *http.Request) {
		// Get the book from request url
		book := chi.URLParam(r, "book")

		// Check if the book exists
		if _, ok := books[book]; !ok {
			http.NotFound(w, r)
			return
		}

		// Render the book
		bookPage := page.BookPage(book, books[book])

		if err := Render(w, r, http.StatusOK, bookPage); err != nil {
			slog.ErrorContext(r.Context(), "error rendering book page", slog.String("error", err.Error()))
		}
	}

}
