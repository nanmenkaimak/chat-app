package dbrepo

import (
	"github.com/google/uuid"
	"github.com/nanmenkaimak/chat-app/internal/models"
	"github.com/pkg/errors"
)

func (m *postgresDBRepo) CreateSession(newSession models.Session) (uuid.UUID, error) {
	var sessionID uuid.UUID

	err := m.DB.Get(&sessionID,
		`insert into sessions (user_id, refresh_token, user_agent, client_ip, expires_at)
				values ($1, $2, $3, $4, $5)
				returning id`,
		newSession.UserID, newSession.RefreshToken, newSession.UserAgent, newSession.ClientIP, newSession.ExpiresAt)
	if err != nil {
		return uuid.Nil, errors.Wrap(err, "insert session")
	}

	return sessionID, nil
}

func (m *postgresDBRepo) GetSession(renewAccessTokenRequest models.RenewAccessTokenRequest) (models.Session, error) {
	var session models.Session

	err := m.DB.Get(&session,
		`select id, user_id, refresh_token, user_agent, client_ip, is_blocked, expires_at, created_at
				from sessions where refresh_token = $1`, renewAccessTokenRequest.RefreshToken)

	return session, err
}
