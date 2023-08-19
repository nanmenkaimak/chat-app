package handlers

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/nanmenkaimak/chat-app/internal/forms"
	"github.com/nanmenkaimak/chat-app/internal/models"
	"github.com/pkg/errors"
	"net/http"
	"unicode"
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
	form.MinLength(newUser.Password, 8)

	if !form.Valid() {
		return fiber.NewError(http.StatusBadRequest, errors.New("your email or password are not valid").Error())
	}

	err := validPassword(newUser.Password)
	if err != nil {
		return fiber.NewError(http.StatusBadRequest, errors.Wrap(err, "sign up").Error())
	}

	id, err := m.DB.CreateUser(newUser)
	if err != nil {
		return fiber.NewError(http.StatusInternalServerError, errors.Wrap(err, "insert user").Error())
	}

	return c.Status(http.StatusCreated).JSON(signUpResponse{
		UserID: id,
	})
}

func validPassword(s string) error {
next:
	for name, classes := range map[string][]*unicode.RangeTable{
		"upper case": {unicode.Upper, unicode.Title},
		"lower case": {unicode.Lower},
		"numeric":    {unicode.Number, unicode.Digit},
		"special":    {unicode.Space, unicode.Symbol, unicode.Punct, unicode.Mark},
	} {
		for _, r := range s {
			if unicode.IsOneOf(classes, r) {
				continue next
			}
		}
		return fmt.Errorf("password must have at least one %s character", name)
	}
	return nil
}
