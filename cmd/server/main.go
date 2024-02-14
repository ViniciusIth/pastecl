package main

import (
	"log"
	"net/http"
	"pastecl/internal/database"
	"pastecl/internal/jwt"
	"pastecl/internal/paste"
	"pastecl/internal/user"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	err := database.ConnectDB()
	if err != nil {
		log.Fatal(err)
	}

	err = database.InitializeDB()
	if err != nil {
		log.Fatal(err)
	}

	jwt.Init()

	r := chi.NewRouter()

	r.Use(middleware.Logger)
    r.Use(middleware.Recoverer)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome"))
	})

	r.Mount("/paste", paste.AddPasteRoutes())
	r.Mount("/user", user.AddUserRoutes())

	chi.Walk(r, func(method string, route string, handler http.Handler, middlewares ...func(http.Handler) http.Handler) error {
		log.Printf("[%s]: '%s' has %d middlewares\n", method, route, len(middlewares))
		return nil
	})

	http.ListenAndServe(":3000", r)
}
