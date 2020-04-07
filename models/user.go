package models

import (
	"github.com/google/uuid"
	"time"
)

type User struct {
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	ID        uuid.UUID `json:"id,omitempty"`
	BirthDate time.Time `json:"birthdate"`
}

//UserRegister type is utilized by register + login functions which
type UserRegister struct {
	*User
	Password string `json:"password,omitempty"`
}

type UserService interface {
	CreateUser(username, email, password string, birthDate time.Time) (error, uuid.UUID)
	Login(id, password string) (*User, error)
}
