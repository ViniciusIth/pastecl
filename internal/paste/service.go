package paste

import (
	"encoding/json"
	"log"
	"net/http"
)

var newPasteHandler = func(w http.ResponseWriter, r *http.Request) {
	var pasteStructure Paste
	json.NewDecoder(r.Body).Decode(&pasteStructure)

	newPaste, err := CreateAnonPaste(pasteStructure.Title, pasteStructure.ExpiresAt, pasteStructure.Visibility)
	if err != nil {
		log.Fatal(err)
	}

	newPaste.SaveToDB()

    jsonResult, err := json.Marshal(newPaste)
    if err != nil {
        log.Fatal(err)
    }

    w.Write(jsonResult)
    w.WriteHeader(http.StatusCreated)
}
