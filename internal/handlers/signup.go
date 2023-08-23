package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/nanmenkaimak/chat-app/internal/forms"
	"github.com/nanmenkaimak/chat-app/internal/models"
	"github.com/pkg/errors"
	"net/http"
)

type signUpResponse struct {
	UserID uuid.UUID `json:"user_id"`
}

func (m *Repository) SignUp(c *fiber.Ctx) error {
	var newUser models.User

	if err := c.BodyParser(&newUser); err != nil {
		return fiber.NewError(http.StatusBadRequest, errors.Wrap(err, "signup json parse").Error())
	}

	form := forms.New()

	form.IsEmail(newUser.Email)
	if form.Errors.Get(newUser.Email) != "" {
		return fiber.NewError(http.StatusBadRequest, form.Errors.Get(newUser.Email))
	}
	form.MinLength(newUser.Password, 8)
	if form.Errors.Get(newUser.Password) != "" {
		return fiber.NewError(http.StatusBadRequest, form.Errors.Get(newUser.Password))
	}
	form.ValidPassword(newUser.Password)
	if form.Errors.Get(newUser.Password) != "" {
		return fiber.NewError(http.StatusBadRequest, form.Errors.Get(newUser.Password))
	}

	id, err := m.DB.CreateUser(newUser)
	if err != nil {
		return fiber.NewError(http.StatusInternalServerError, errors.Wrap(err, "insert user").Error())
	}

	return c.Status(http.StatusCreated).JSON(signUpResponse{
		UserID: id,
	})
}
