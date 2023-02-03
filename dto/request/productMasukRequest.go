package request

type ProductMasukRequest struct {
	Name       string `json:"name" validate:"required"`
	Qty        string `json:"qty" validate:"required"`
	ProductID  int    `json:"product_id" validate:"required"`
	SupplierID int    `json:"supplier_id" validate:"required"`
}
