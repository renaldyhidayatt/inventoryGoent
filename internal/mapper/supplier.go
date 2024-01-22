package mapper

import (
	"github.com/renaldyhidayatt/inventorygoent/internal/domain/response"
	"github.com/renaldyhidayatt/inventorygoent/pkg/database/postgres/ent"
)

type supplierMapper struct{}

func NewSupplierMapper() *supplierMapper {
	return &supplierMapper{}
}

func (m *supplierMapper) ToSupplierResponse(supplier *ent.Supplier) *response.Supplier {

	return &response.Supplier{
		Name:      supplier.Name,
		Alamat:    supplier.Alamat,
		Telepon:   supplier.Telepon,
		CreatedAt: supplier.CreatedAt,
		UpdatedAt: supplier.UpdatedAt,
	}
}

func (m *supplierMapper) ToSupplierResponses(suppliers []*ent.Supplier) *[]response.Supplier {
	var supplierResponses []response.Supplier

	for _, supplier := range suppliers {
		supplierResponse := response.Supplier{
			Name:      supplier.Name,
			Alamat:    supplier.Alamat,
			Telepon:   supplier.Telepon,
			CreatedAt: supplier.CreatedAt,
			UpdatedAt: supplier.UpdatedAt,
		}
		supplierResponses = append(supplierResponses, supplierResponse)
	}

	return &supplierResponses
}
