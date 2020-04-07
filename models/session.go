package models

import (
	"github.com/google/uuid"
)

type Session struct {
	SessionId uuid.UUID
	User      *User
}

type SessionService interface {
	Set(session *Session) error
	Get(sessionId uuid.UUID) (bool, error)
	Delete(sessionId uuid.UUID) error
}
