package request

type ProductKeluarRequest struct {
	Qty        string `json:"qty" validate:"required"`
	ProductID  int    `json:"product_id" validate:"required"`
	CategoryID int    `json:"category_id" validate:"required"`
}
