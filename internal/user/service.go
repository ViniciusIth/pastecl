package user

import (
	"encoding/json"
	"log"
	"net/http"
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
}

func UserLoginHandler(w http.ResponseWriter, r *http.Request) {
    r.ParseForm()

    email := r.FormValue("email")
    password := r.FormValue("password")

    user, err := CheckUserCredentials(email, password)
    if err != nil {
        w.WriteHeader(http.StatusUnauthorized)
        w.Write([]byte(err.Error()))
    }

    token, err := user.generateJWT()
    if err != nil {
        log.Fatal(err)
    }

    loginResult := struct {
        Token string `json:"token"`
        UserData User `json:"user"`
    }{
        Token:  *token,
        UserData: *user,
    }

    res, err := json.Marshal(loginResult)
    if err != nil {
        log.Fatal(err)
    }
    w.Write(res)
    w.WriteHeader(http.StatusOK)
}
