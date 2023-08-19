package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/nanmenkaimak/chat-app/internal/JWT"
	"github.com/nanmenkaimak/chat-app/internal/forms"
	"github.com/nanmenkaimak/chat-app/internal/models"
	"github.com/pkg/errors"
	"net/http"
	"os"
	"time"
)

type loginResponse struct {
	SessionID       uuid.UUID `json:"session_id"`
	AccessToken     string    `json:"access_token"`
	AccessDuration  time.Time `json:"access_duration"`
	RefreshToken    string    `json:"refresh_token"`
	RefreshDuration time.Time `json:"refresh_duration"`
}

func (m *Repository) Login(c *fiber.Ctx) error {
	var userLogin models.LoginRequest

	if err := c.BodyParser(&userLogin); err != nil {
		return fiber.NewError(http.StatusBadRequest, errors.Wrap(err, "login json parse").Error())
	}

	form := forms.New()

	form.IsEmail(userLogin.Email)
	form.MinLength(userLogin.Password, 8)

	if !form.Valid() {
		return fiber.NewError(http.StatusBadRequest, errors.New("your email or password are not valid").Error())
	}

	err := validPassword(userLogin.Password)
	if err != nil {
		return fiber.NewError(http.StatusBadRequest, errors.Wrap(err, "login").Error())
	}

	id, _, err := m.DB.Authenticate(userLogin)
	if err != nil {
		return fiber.NewError(http.StatusInternalServerError, errors.Wrap(err, "login user").Error())
	}

	accessDuration, _ := time.ParseDuration(os.Getenv("ACCESS_TOKEN_TIME"))
	refreshDuration, _ := time.ParseDuration(os.Getenv("REFRESH_TOKEN_TIME"))

	accessToken, err := JWT.GenerateToken(id.String(), accessDuration)
	if err != nil {
		return fiber.NewError(http.StatusInternalServerError, errors.Wrap(err, "generate access token").Error())
	}

	refreshToken, err := JWT.GenerateToken(id.String(), refreshDuration)
	if err != nil {
		return fiber.NewError(http.StatusInternalServerError, errors.Wrap(err, "generate refresh token").Error())
	}

	newSession := models.Session{
		UserID:       id,
		RefreshToken: refreshToken,
		UserAgent:    string(c.Request().Header.UserAgent()),
		ClientIP:     c.IP(),
		ExpiresAt:    time.Now().Add(refreshDuration),
	}

	sessionID, err := m.DB.CreateSession(newSession)
	if err != nil {
		return fiber.NewError(http.StatusInternalServerError, errors.Wrap(err, "session").Error())
	}

	resp := loginResponse{
		SessionID:       sessionID,
		RefreshToken:    refreshToken,
		RefreshDuration: time.Now().Add(refreshDuration),
		AccessToken:     accessToken,
		AccessDuration:  time.Now().Add(accessDuration),
	}
	return c.Status(http.StatusOK).JSON(resp)
}
