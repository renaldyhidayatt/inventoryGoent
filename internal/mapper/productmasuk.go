package mapper

import (
	"github.com/renaldyhidayatt/inventorygoent/internal/domain/response"
	"github.com/renaldyhidayatt/inventorygoent/pkg/database/postgres/ent"
)

type productMasukMapper struct{}

func NewProductMasukMapper() *productMasukMapper {
	return &productMasukMapper{}
}

func (m *productMasukMapper) ToProductMasukResponse(product *ent.ProductMasuk) *response.ProductMasuk {
	// for _, supplier := range product.Edges.Supplier {
	// 	supplierResponse := response.Supplier{
	// 		Name:      supplier.Name,
	// 		CreatedAt: supplier.CreatedAt,
	// 		UpdatedAt: supplier.UpdatedAt,
	// 	}
	// }

	// for _, product := range product.Edges.Product {
	// 	productResponse := response.Product{
	// 		Name:      product.Name,
	// 		Image:     product.Image,
	// 		Qty:       product.Qty,
	// 		CreatedAt: product.CreatedAt,
	// 		UpdatedAt: product.UpdatedAt,
	// 	}
	// }

	return &response.ProductMasuk{
		Name:      product.Name,
		Qty:       product.Qty,
		CreatedAt: product.CreatedAt,
		UpdatedAt: product.UpdatedAt,
	}
}

func (m *productMasukMapper) ToProductMasukResponses(products []*ent.ProductMasuk) *[]response.ProductMasuk {
	var productResponses []response.ProductMasuk

	for _, product := range products {
		productResponse := m.ToProductMasukResponse(product)

		productResponses = append(productResponses, *productResponse)
	}

	return &productResponses

}
