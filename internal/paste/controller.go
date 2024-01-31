package paste

import (
	"pastecl/internal/jwt"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/jwtauth"
)

func AddPasteRoutes() *chi.Mux {
	r := chi.NewRouter()

    r.Use(jwtauth.Verifier(jwt.AuthGenerator))

	r.Post("/new", newPasteHandler)
    r.Get("/{id}", getPasteHandler)
    r.Get("/{id}/file", getPasteFileHandler)

	return r
}
