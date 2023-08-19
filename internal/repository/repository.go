package repository

import (
	"github.com/google/uuid"
	"github.com/nanmenkaimak/chat-app/internal/models"
)

type UserRepo interface {
	CreateUser(newUser models.User) (uuid.UUID, error)
	Authenticate(userLogin models.LoginRequest) (uuid.UUID, string, error)
}

type SessionRepo interface {
	CreateSession(newSession models.Session) (uuid.UUID, error)
	GetSession(renewAccessTokenRequest models.RenewAccessTokenRequest) (models.Session, error)
}

type MessageRepo interface {
}

type DatabaseRepo interface {
	UserRepo
	SessionRepo
	MessageRepo
}
