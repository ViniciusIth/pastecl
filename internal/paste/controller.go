package paste

import (
	"github.com/go-chi/chi/v5"
)

func AddPasteRoutes() *chi.Mux {
	r := chi.NewRouter()

	r.Post("/new", newPasteHandler)
    r.Get("/{id}", getPasteHandler)
    r.Get("/{id}/file", getPasteFileHandler)

	return r
}
