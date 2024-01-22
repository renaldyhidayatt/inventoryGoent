package response

import "time"

type ProductKeluar struct {
	Qty       string    `json:"qty"`
	Products  []Product `json:"products"`
	Category  Category  `json:"category"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
