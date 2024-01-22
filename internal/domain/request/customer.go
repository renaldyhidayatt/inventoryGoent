package request

import "github.com/go-playground/validator/v10"

type CreateCustomerRequest struct {
	Name    string `json:"name" validate:"required,lowercase" schema:"name"`
	Alamat  string `json:"alamat" validate:"required,max=1000"`
	Telepon string `json:"phone" validate:"required,gte=12"`
}

type UpdateCustomerRequest struct {
	ID      int    `json:"id" validate:"required" schema:"id"`
	Name    string `json:"name" validate:"required,lowercase" schema:"name"`
	Alamat  string `json:"alamat" validate:"required,max=1000"`
	Telepon string `json:"phone" validate:"required,gte=12"`
}

func (c *CreateCustomerRequest) Validate() error {

	validate := validator.New()

	err := validate.Struct(c)

	if err != nil {
		return err
	}

	return nil
}

func (c *UpdateCustomerRequest) Validate() error {

	validate := validator.New()

	err := validate.Struct(c)

	if err != nil {
		return err
	}

	return nil
}
