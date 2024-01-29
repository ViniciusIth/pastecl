package user

import "github.com/go-chi/chi/v5"

func AddUserRoutes() *chi.Mux {
	r := chi.NewRouter()

	r.Post("/new", UserRegistrationHandler)
    // r.Get("/{id}", getPasteHandler)
    // r.Get("/{id}/file", getPasteFileHandler)

	return r
}
