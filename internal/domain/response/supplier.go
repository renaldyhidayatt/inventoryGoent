package response

import "time"

type Supplier struct {
	Name      string    `json:"name"`
	Alamat    string    `json:"alamat"`
	Telepon   string    `json:"telepon"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
