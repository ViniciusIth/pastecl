package user

import (
	"encoding/json"
	"log"
	"net/http"
	"pastecl/internal/paste"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/jwtauth"
)

func UserRegistrationHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	username := r.Form.Get("username")
	email := r.Form.Get("email")
	password := r.Form.Get("password")

	user, err := CreateNewUser(username, email, password)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		log.Fatal(err)
	}

	err = user.SaveToDB()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		log.Fatal(err)
	}

	w.WriteHeader(http.StatusCreated)
}

func UserLoginHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	email := r.FormValue("email")
	password := r.FormValue("password")

	user, err := CheckUserCredentials(email, password)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte(err.Error()))
		return
	}

	token, err := user.generateJWT()
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte(err.Error()))
		return
	}

	loginResult := struct {
		Token    string `json:"token"`
		UserData User   `json:"user"`
	}{
		Token:    *token,
		UserData: *user,
	}

	res, err := json.Marshal(loginResult)
	if err != nil {
		log.Fatal(err)
	}
	w.Write(res)
	w.WriteHeader(http.StatusOK)
}

func getUserHandler(w http.ResponseWriter, r *http.Request) {
	userID := chi.URLParam(r, "id")

	// _, claims, err := jwtauth.FromContext(r.Context())
	// if err != nil {
	// 	log.Fatal(err)
	// }
	//
	// subject := claims["sub"].(string)

	getUser, err := getUserByID(userID)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(err.Error()))
		return
	}

	getUser.Password = ""

	jsonRes, err := json.Marshal(getUser)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.Write(jsonRes)
	w.WriteHeader(http.StatusOK)
}

func getUserPastesHandler(w http.ResponseWriter, r *http.Request) {
	pasteId := chi.URLParam(r, "id")
	var subject string

	_, claims, err := jwtauth.FromContext(r.Context())
	if err != nil {
	    log.Println("Anonymous request to user pastes")
        claims["sub"] = ""
	}

	subject = claims["sub"].(string)

	userPastes, err := paste.GetPastesByUser(pasteId, subject)
	if err != nil {
		log.Fatal(err)
	}

	jsonRes, err := json.Marshal(userPastes)
	if err != nil {
		log.Fatal(err)
	}

	w.Write(jsonRes)
	w.WriteHeader(http.StatusOK)
}
