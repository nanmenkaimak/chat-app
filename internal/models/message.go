package models

import (
	"github.com/google/uuid"
	"time"
)

type Message struct {
	ID         uuid.UUID `json:"id" db:"id"`
	Text       string    `json:"text" db:"text"`
	SenderID   string    `json:"sender_id" db:"sender_id"`
	ReceiverID string    `json:"receiver_id" db:"receiver_id"`
	CreatedAt  time.Time `json:"created_at" db:"created_at"`
	UpdatedAt  time.Time `json:"updated_at" db:"updated_at"`
}
