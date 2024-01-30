package jwt

import (
	"log"

	"github.com/go-chi/jwtauth"
)


var AuthGenerator *jwtauth.JWTAuth

func Init() {
    AuthGenerator = jwtauth.New("HS256", []byte("secret"), nil)

	_, tokenString, _ := AuthGenerator.Encode(map[string]interface{}{"user_id": 123})
	log.Printf("DEBUG: a sample jwt is %s\n", tokenString)
}

