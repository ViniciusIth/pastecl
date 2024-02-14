package user

import (
	"errors"
	"pastecl/internal/database"
	"pastecl/internal/jwt"
	"time"

	"github.com/google/uuid"
)

type User struct {
	UUID      string `json:"uuid"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	Password  string `json:"password,omitempty"`
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

	return &newUser, nil
}

func CheckUserCredentials(email string, password string) (*User, error) {
	var user User

	sqlStatement := `SELECT * FROM users WHERE email = ?;`
	err := database.Access.QueryRow(sqlStatement, email).Scan(
		&user.UUID, &user.Username, &user.Email, &user.Password, &user.CreatedAt,
	)

	if err != nil {
		// Handle the error (user not found, database error, etc.)
		// You might want to check if err == sql.ErrNoRows to handle the case when no user is found.
		return nil, err
	}

	if password != user.Password {
		return nil, errors.New("Wrong password")
	}

    user.Password = ""

	return &user, nil
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

func (user *User) generateJWT() (jwt_token *string, err error) {
	_, tokenString, err := jwt.AuthGenerator.Encode(map[string]interface{}{"sub": user.UUID})
	if err != nil {
		return nil, err
	}

	return &tokenString, nil
}
