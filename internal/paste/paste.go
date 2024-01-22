package paste

import (
	"strconv"
	"time"

	"github.com/google/uuid"
)

type Paste struct {
	UUID       string `json:"uuid"`
	Title      string `json:"title"`
	CreatedAt  string `json:"created_at"`
	ExpiresAt  string `json:"expires_at"`
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
		CreatedAt:  strconv.FormatInt(time.Now().Unix(), 10),
		ExpiresAt:  strconv.FormatInt(time.Now().AddDate(0, 6, 0).Unix(), 10),
		Visibility: visibility,
		ControlKey: controlKey.String(),
	}

	return &newPaste, nil
}
