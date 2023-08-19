package models

import (
	"github.com/google/uuid"
	"time"
)

type User struct {
	ID        uuid.UUID `json:"id" db:"id"`
	Username  string    `json:"username" db:"username"`
	FirstName string    `json:"first_name" db:"first_name"`
	LastName  string    `json:"last_name" db:"last_name"`
	Email     string    `json:"email" db:"email"`
	Password  string    `json:"password" db:"password"`
	Role      int       `json:"role_id" db:"role_id"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

type Role struct {
	ID   int    `json:"id" db:"id"`
	Name string `json:"name" db:"name"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
