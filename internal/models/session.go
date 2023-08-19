package models

import (
	"github.com/google/uuid"
	"time"
)

type Session struct {
	ID           uuid.UUID `json:"id" db:"id"`
	UserID       uuid.UUID `json:"user_id" db:"user_id"`
	RefreshToken string    `json:"refresh_token" db:"refresh_token"`
	UserAgent    string    `json:"user_agent" db:"user_agent"`
	ClientIP     string    `json:"client_ip" db:"client_ip"`
	IsBlocked    bool      `json:"is_blocked" db:"is_blocked"`
	ExpiresAt    time.Time `json:"expires_at" db:"expires_at"`
	CreatedAt    time.Time `json:"created_at" db:"created_at"`
}

type RenewAccessTokenRequest struct {
	RefreshToken string `json:"refresh_token"`
}
