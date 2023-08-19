package handlers

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/nanmenkaimak/chat-app/internal/JWT"
	"github.com/nanmenkaimak/chat-app/internal/models"
	"github.com/pkg/errors"
	"net/http"
	"os"
	"time"
)

type renewAccessTokenResponse struct {
	AccessToken          string    `json:"access_token"`
	AccessTokenExpiresAt time.Time `json:"access_token_expires_at"`
}

func (m *Repository) RenewAccessToken(c *fiber.Ctx) error {
	var req models.RenewAccessTokenRequest

	if err := c.BodyParser(&req); err != nil {
		return fiber.NewError(http.StatusBadRequest, errors.Wrap(err, "renew access token").Error())
	}

	claims, err := JWT.ParseToken(req.RefreshToken)
	if err != nil {
		return fiber.NewError(http.StatusBadRequest, errors.Wrap(err, "parse token").Error())
	}

	session, err := m.DB.GetSession(req)
	if err != nil {
		return fiber.NewError(http.StatusInternalServerError, errors.Wrap(err, "get session").Error())
	}

	if session.IsBlocked {
		return fiber.NewError(http.StatusUnauthorized, errors.New("blocked session").Error())
	}

	if session.UserID.String() != fmt.Sprintf("%v", claims["id"]) {
		return fiber.NewError(http.StatusUnauthorized, errors.New("incorrect session user").Error())
	}

	if time.Now().After(session.ExpiresAt) {
		return fiber.NewError(http.StatusUnauthorized, errors.New("expired session").Error())
	}

	accessDuration, _ := time.ParseDuration(os.Getenv("ACCESS_TOKEN_TIME"))

	accessToken, err := JWT.GenerateToken(fmt.Sprintf("%v", claims["id"]), accessDuration)
	if err != nil {
		return fiber.NewError(http.StatusInternalServerError, errors.Wrap(err, "generate access token").Error())
	}

	rsp := renewAccessTokenResponse{
		AccessToken:          accessToken,
		AccessTokenExpiresAt: time.Now().Add(accessDuration),
	}

	return c.Status(http.StatusOK).JSON(rsp)
}
