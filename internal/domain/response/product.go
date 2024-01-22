package response

import "time"

type Product struct {
	Name      string    `json:"name"`
	Image     string    `json:"image"`
	Qty       string    `json:"qty"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Category  Category  `json:"category"`
}
