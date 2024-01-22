package request

import "github.com/go-playground/validator/v10"

type CreateProductMasukRequest struct {
	Name       string `json:"name" validate:"required"`
	Qty        string `json:"qty" validate:"required"`
	ProductID  int    `json:"product_id" validate:"required"`
	SupplierID int    `json:"supplier_id" validate:"required"`
}

type UpdateProductMasukRequest struct {
	ID         int    `json:"id" validate:"required"`
	Name       string `json:"name" validate:"required"`
	Qty        string `json:"qty" validate:"required"`
	ProductID  int    `json:"product_id" validate:"required"`
	SupplierID int    `json:"supplier_id" validate:"required"`
}

func (c *CreateProductMasukRequest) Validate() error {

	validate := validator.New()

	err := validate.Struct(c)

	if err != nil {
		return err
	}

	return nil
}

func (c *UpdateProductMasukRequest) Validate() error {

	validate := validator.New()

	err := validate.Struct(c)

	if err != nil {
		return err
	}

	return nil
}
