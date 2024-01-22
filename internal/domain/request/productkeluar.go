package request

import "github.com/go-playground/validator/v10"

type CreateProductKeluarRequest struct {
	Qty        string `json:"qty" validate:"required"`
	ProductID  int    `json:"product_id" validate:"required"`
	CategoryID int    `json:"category_id" validate:"required"`
}

type UpdateProductKeluarRequest struct {
	ID         int    `json:"id" validate:"required"`
	Qty        string `json:"qty" validate:"required"`
	ProductID  int    `json:"product_id" validate:"required"`
	CategoryID int    `json:"category_id" validate:"required"`
}

func (c *CreateProductKeluarRequest) Validate() error {

	validate := validator.New()

	err := validate.Struct(c)

	if err != nil {
		return err
	}

	return nil
}

func (c *UpdateProductKeluarRequest) Validate() error {

	validate := validator.New()

	err := validate.Struct(c)

	if err != nil {
		return err
	}

	return nil
}
