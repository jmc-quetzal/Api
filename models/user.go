package user

import (
	"github.com/google/uuid"
	"time"
)

type User struct {
	Username string `json:"username",omitempty`
	Email string
	ID uuid.UUID
	Birthdate time.Time
	Password string
}

type UserService interface {
	CreateUser(username, email, password string,birthDate time.Time) (error, *User)
	Authenticate(id, password string) (error)
}