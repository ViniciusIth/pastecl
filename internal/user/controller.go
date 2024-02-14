package user

import (
	"pastecl/internal/jwt"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/jwtauth"
)

func AddUserRoutes() *chi.Mux {
	r := chi.NewRouter()

	r.Use(jwtauth.Verifier(jwt.AuthGenerator))

	r.Post("/new", UserRegistrationHandler)
	r.Post("/login", UserLoginHandler)
	r.Get("/{id}/pastes", getUserPastesHandler)

	return r
}
