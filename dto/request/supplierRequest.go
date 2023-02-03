package request

type SupplierRequest struct {
	Name    string `json:"name" validate:"required,lowercase" schema:"name"`
	Alamat  string `json:"alamat" validate:"required,max=1000"`
	Telepon string `json:"phone" validate:"required,gte=12"`
}
