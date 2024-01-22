package request

import "github.com/go-playground/validator/v10"

type CreateProductRequest struct {
	Name       string `json:"name" validate:"required,lowercase"`
	Image      string `json:"image" validate:"required"`
	Qty        string `json:"qty" validate:"required"`
	CategoryID int    `json:"category_id" validate:"required"`
}

type UpdateProductRequest struct {
	ID         int    `json:"id" validate:"required"`
	Name       string `json:"name" validate:"required,lowercase"`
	Image      string `json:"image" validate:"required"`
	Qty        string `json:"qty" validate:"required"`
	CategoryID int    `json:"category_id" validate:"required"`
}

func (c *CreateProductRequest) Validate() error {

	validate := validator.New()

	err := validate.Struct(c)

	if err != nil {
		return err
	}

	return nil
}

func (c *UpdateProductRequest) Validate() error {

	validate := validator.New()

	err := validate.Struct(c)

	if err != nil {
		return err
	}

	return nil
}
