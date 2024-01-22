package main

import (
	"log"
	"pastecl/internal/paste"
)

func main() {
    anonPaste, err := paste.CreateAnonPaste("Test", "", true)
    if err != nil {
        log.Fatal(err)
    }

    log.Println(anonPaste)
}
