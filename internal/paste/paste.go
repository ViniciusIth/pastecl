package paste

import (
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
}

func CreateAnonPaste(title string, expires_at string, visibility bool) (*Paste, error) {
	pasteID, err := uuid.NewV7()
	if err != nil {
		return nil, err
	}

	// Just for tests, use a better thing in the future
	controlKey, err := uuid.NewV7()

	newPaste := Paste{
		UUID:       pasteID.String(),
		Title:      title,
		CreatedAt:  time.Now().Unix(),
		ExpiresAt:  time.Now().AddDate(0, 6, 0).Unix(),
		Visibility: visibility,
		ControlKey: controlKey.String(),
	}

	return &newPaste, nil
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
