package jwt

import (
	"log"

	"github.com/go-chi/jwtauth"
)


var tokenAuth *jwtauth.JWTAuth

func Init() {
    tokenAuth = jwtauth.New("HS256", []byte("secret"), nil)

	_, tokenString, _ := tokenAuth.Encode(map[string]interface{}{"user_id": 123})
	log.Printf("DEBUG: a sample jwt is %s\n", tokenString)
}
