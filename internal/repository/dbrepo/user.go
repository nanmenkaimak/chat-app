package dbrepo

import (
	"github.com/google/uuid"
	"github.com/nanmenkaimak/chat-app/internal/models"
	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
)

func (m *postgresDBRepo) CreateUser(newUser models.User) (uuid.UUID, error) {
	var userID uuid.UUID
	password, err := hashPassword(newUser.Password)
	if err != nil {
		return uuid.Nil, err
	}
	err = m.DB.Get(&userID,
		`insert into users (username, first_name, last_name, email, password) 
    			values ($1, $2, $3, $4, $5)
    			returning id`,
		newUser.Username, newUser.FirstName, newUser.LastName, newUser.Email, password)
	if err != nil {
		return uuid.Nil, errors.Wrap(err, "insert user")
	}

	return userID, nil
}

func (m *postgresDBRepo) Authenticate(userLogin models.LoginRequest) (uuid.UUID, string, error) {
	var user models.User

	err := m.DB.Get(&user,
		`select id, username, password from users where email = $1`, userLogin.Email)
	if err != nil {
		return uuid.Nil, "", errors.Wrap(err, "select auth")
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userLogin.Password))
	if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
		return uuid.Nil, "", errors.Wrap(err, "incorrect password")
	} else if err != nil {
		return uuid.Nil, "", errors.Wrap(err, "password auth")
	}

	return user.ID, user.Username, nil
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}
