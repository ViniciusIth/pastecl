package paste

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

var newPasteHandler = func(w http.ResponseWriter, r *http.Request) {
    err := r.ParseMultipartForm(2 << 18) // 512 kb, I think
    if err != nil {
        log.Fatal(err)
    }

    _, header, err := r.FormFile("content")
    if err != nil {
        log.Fatal(err)
    }

    expires, _ := strconv.ParseInt(r.FormValue("expire_at"), 10, 64)
    visibility, _ := strconv.ParseBool(r.FormValue("visibility"))

	newPaste, err := CreateAnonPaste(r.FormValue("title"), expires, visibility, header)
	if err != nil {
		log.Fatal(err)
	}

    err = newPaste.SaveToFile()
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

var getPasteHandler = func (w http.ResponseWriter, r *http.Request) {
    pasteId := chi.URLParam(r, "id")
    
    savedPaste, err := getPasteById(pasteId)
    if err != nil {
        w.WriteHeader(http.StatusInternalServerError)
        log.Fatal(err)
    }

    jsonResult, err := json.Marshal(savedPaste)
    w.Write(jsonResult)
}
