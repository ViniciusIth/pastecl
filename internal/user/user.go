package user

import (
	"log"
	"pastecl/internal/database"
	"time"

	"github.com/google/uuid"
)

type User struct {
	UUID      string `json:"uuid"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	CreatedAt int64  `json:"created_at"`
}

func CreateNewUser(username string, email string, password string) (*User, error) {
	userID, err := uuid.NewV7()
	if err != nil {
		return nil, err
	}

	newUser := User{
		UUID:      userID.String(),
		Username:  username,
		Email:     email,
		Password:  password,
		CreatedAt: time.Now().Unix(),
	}

	log.Println(newUser.Username)

	return &newUser, nil
}

func (user *User) SaveToDB() error {
	sqlStatement := `INSERT INTO users 
    (uuid, username, email, password, created_at)
    VALUES (?, ?, ?, ?, ?);`

	_, err := database.Access.Exec(sqlStatement,
		user.UUID, user.Username, user.Email, user.Password, user.CreatedAt)

	if err != nil {
		return err
	}

	return nil
}
