package request

import (
	"errors"

	"github.com/go-playground/validator/v10"
)

type AuthRequest struct {
	Email    string `json:"email" validate:"required,lowercase,email"`
	Password string `json:"password" validate:"required,lowercase"`
}

type RegisterRequest struct {
	Firstname       string `json:"firstname" validate:"required,lowercase"`
	LastName        string `json:"lastname" validate:"required,lowercase"`
	Email           string `json:"email" validate:"required,lowercase"`
	Password        string `json:"password" validate:"required,lowercase"`
	ConfirmPassword string `json:"confirm_password"`
}

func (c *AuthRequest) Validate() error {

	validate := validator.New()

	err := validate.Struct(c)

	if err != nil {
		return err
	}

	return nil
}

func (c *RegisterRequest) Validate() error {

	if c.Password != c.ConfirmPassword {
		return errors.New("password and confirm_password do not match")
	}

	validate := validator.New()

	err := validate.Struct(c)

	if err != nil {
		return err
	}

	return nil
}
