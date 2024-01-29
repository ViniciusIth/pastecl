package user

import (
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
