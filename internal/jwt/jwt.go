package jwt

import (
	"log"

	"github.com/go-chi/jwtauth"
)

var AuthGenerator *jwtauth.JWTAuth

func Init() {
	AuthGenerator = jwtauth.New("HS256", []byte("secret"), nil)

	_, tokenString, _ := AuthGenerator.Encode(map[string]interface{}{"sub": "018d568b-0e0d-7cb9-ada6-d6c8ef3f94bf"})
	log.Printf("JWT: a sample jwt is %s\n", tokenString)
	_, err := AuthGenerator.Decode(tokenString)
	if err != nil {
		log.Panic(err)
	}
	log.Printf("JWT: the sample jwt is valid.")
}
