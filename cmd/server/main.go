package main

import (
	"log"
	"pastecl/internal/database"
	"pastecl/internal/paste"
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

    anonPaste, err := paste.CreateAnonPaste("Test", "", true)
    if err != nil {
        log.Fatal(err)
    }

    err = anonPaste.SaveToDB()
    if err != nil {
        log.Fatal(err)
    }
}
