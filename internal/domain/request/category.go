package request

import "github.com/go-playground/validator/v10"

type CreateCategoryRequest struct {
	Name string `json:"name" validate:"required, lowercase"`
}

type UpdateCategoryRequest struct {
	ID   int    `json:"id" validate:"required"`
	Name string `json:"name" validate:"required, lowercase"`
}

func (r *CreateCategoryRequest) Validate() error {
	validate := validator.New()

	err := validate.Struct(r)

	if err != nil {
		return err
	}

	return nil
}

func (r *UpdateCategoryRequest) Validate() error {
	validate := validator.New()

	err := validate.Struct(r)

	if err != nil {
		return err
	}

	return nil
}
