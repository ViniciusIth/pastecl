package paste

import (
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"pastecl/internal/database"
	"time"

	"github.com/google/uuid"
)

type Paste struct {
	UUID       string `json:"uuid"`
	Title      string `json:"title"`
	CreatedAt  int64  `json:"created_at"`
	ExpiresAt  int64  `json:"expires_at"`
	Visibility bool   `json:"visibility"`
	ControlKey string `json:"control_key"`
	OwnerId    string `json:"ownser_id"`
	FileURL    string `json:"file_url"`
}

func CreateAnonPaste(title string, expires_at int64, visibility bool, file *multipart.FileHeader) (*Paste, error) {
	pasteID, err := uuid.NewV7()
	if err != nil {
		return nil, err
	}

	if expires_at == 0 {
		expires_at = time.Now().AddDate(0, 6, 0).Unix()
	}

	// Just for tests, use a better thing in the future
	controlKey, err := uuid.NewV7()

	newPaste := Paste{
		UUID:       pasteID.String(),
		Title:      title,
		CreatedAt:  time.Now().Unix(),
		ExpiresAt:  expires_at,
		Visibility: visibility,
		ControlKey: controlKey.String(),
	}

	return &newPaste, nil
}

func CreateUserPaste(title string, expires_at int64, visibility bool, userId string, file *multipart.FileHeader) (*Paste, error) {
	pasteID, err := uuid.NewV7()
	if err != nil {
		return nil, err
	}

	if expires_at == 0 {
		expires_at = time.Now().AddDate(0, 6, 0).Unix()
	}

	newPaste := Paste{
		UUID:       pasteID.String(),
		Title:      title,
		CreatedAt:  time.Now().Unix(),
		ExpiresAt:  expires_at,
		Visibility: visibility,
		OwnerId:    userId,
	}

	return &newPaste, nil
}

func getPasteDataById(id string) (*Paste, error) {
	var savedPaste Paste

	sqlStatement := `SELECT id, title, createdat, expiresat, visibility, ownerid FROM paste WHERE id = ?`
	row := database.Access.QueryRow(sqlStatement, id)

	err := row.Scan(
		&savedPaste.UUID,
		&savedPaste.Title,
		&savedPaste.CreatedAt,
		&savedPaste.ExpiresAt,
		&savedPaste.Visibility,
		&savedPaste.OwnerId,
	)
	if err != nil {
		return nil, err
	}

	savedPaste.FileURL = fmt.Sprintf("http://localhost:3000/paste/%s/file", savedPaste.UUID)
	return &savedPaste, nil
}

func (pst *Paste) SaveToDB() error {
	sqlStatement := `INSERT INTO paste 
    (id, title, createdat, expiresat, visibility, controlkey, ownerid)
    VALUES (?, ?, ?, ?, ?, ?, ?);`
	_, err := database.Access.Exec(sqlStatement,
		pst.UUID, pst.Title, pst.CreatedAt, pst.ExpiresAt, pst.Visibility, pst.ControlKey, pst.OwnerId)

	if err != nil {
		return err
	}

	return nil
}

func SavePasteToFile(paste *Paste, fileHeader *multipart.FileHeader) error {
	file, err := fileHeader.Open()
	if err != nil {
		return err
	}

	defer file.Close()

	saveFile, err := os.OpenFile(fmt.Sprintf("./data/%s.txt", paste.UUID), os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		return err
	}

	defer saveFile.Close()

	_, err = io.Copy(saveFile, file)
	if err != nil {
		return err
	}

	return nil
}
