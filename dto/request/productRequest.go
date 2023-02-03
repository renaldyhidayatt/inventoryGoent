package request

type ProductRequest struct {
	Name       string `json:"name" validate:"required,lowercase"`
	Image      string `json:"image" validate:"required"`
	Qty        string `json:"qty" validate:"required"`
	CategoryID int    `json:"category_id" validate:"required"`
}
