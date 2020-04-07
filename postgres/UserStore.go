package postgres

import (
	"errors"
	"github.com/jmc-quetzal/api/common"
	"github.com/google/uuid"
	"github.com/jmc-quetzal/api/models"
	"github.com/jmoiron/sqlx"
	 "github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"
	"time"
)

//UserStore implements UserService with postges store + session store
//Also
type UserStore struct {
	DB *sqlx.DB
}

//CreateUser on UserStore.
func (store UserStore) CreateUser(username, email, password string, birthDate time.Time) (error, uuid.UUID) {
	var pid uuid.UUID
	err := store.DB.Get(&pid, insertUser, username, email, password, birthDate)
	if err != nil {
		if e, ok := err.(*pq.Error); ok {
		 	switch e.Code {
			case "23505":
				if e.Constraint == "users_email_key" {
					return errors.New("Email already exists"), pid
				}
				if e.Constraint == "users_username_key" {
					return errors.New("Username already exists"), pid
				}
			}
		}
	}
	return err, pid
}

//Login method authenticates a user via email/username + password and creates a session
func (store UserStore) Login(identifier, password string) (*models.User, error) {
	user := models.UserRegister{}
	err := store.DB.Get(&user, authenticateViaEmail, identifier)
	if err != nil {
		return nil, &common.DataError{400,"User not found"}
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password),[]byte(password))
	if err != nil {
		return nil, err
	}
	return user.User, nil
}
