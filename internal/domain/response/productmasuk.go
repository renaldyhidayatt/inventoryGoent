package response

import "time"

type ProductMasuk struct {
	Name      string    `json:"name"`
	Qty       string    `json:"qty"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Product   Product   `json:"product"`
	Supplier  Supplier  `json:"supplier"`
}
