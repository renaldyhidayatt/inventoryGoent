package request

import "github.com/go-playground/validator/v10"

type CreateSupplierRequest struct {
	Name    string `json:"name" validate:"required,lowercase" schema:"name"`
	Alamat  string `json:"alamat" validate:"required,max=1000"`
	Telepon string `json:"phone" validate:"required,gte=12"`
}

type UpdateSupplierRequest struct {
	ID      int    `json:"id" validate:"required" schema:"id"`
	Name    string `json:"name" validate:"required,lowercase" schema:"name"`
	Alamat  string `json:"alamat" validate:"required,max=1000"`
	Telepon string `json:"phone" validate:"required,gte=12"`
}

func (c *CreateSupplierRequest) Validate() error {

	validate := validator.New()

	err := validate.Struct(c)

	if err != nil {
		return err
	}

	return nil
}

func (c *UpdateSupplierRequest) Validate() error {

	validate := validator.New()

	err := validate.Struct(c)

	if err != nil {
		return err
	}

	return nil
}
