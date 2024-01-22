package mapper

import (
	"github.com/renaldyhidayatt/inventorygoent/internal/domain/response"
	"github.com/renaldyhidayatt/inventorygoent/pkg/database/postgres/ent"
)

type productMapper struct{}

func NewProductMapper() *productMapper {
	return &productMapper{}
}

func (m *productMapper) ToProductResponse(product *ent.Product) *response.Product {
	return &response.Product{
		Name:      product.Name,
		Image:     product.Image,
		Qty:       product.Qty,
		CreatedAt: product.CreatedAt,
		UpdatedAt: product.UpdatedAt,
	}
}

func (m *productMapper) ToProductResponses(products []*ent.Product) *[]response.Product {
	var productResponses []response.Product

	for _, product := range products {
		productResponse := response.Product{
			Name:      product.Name,
			Image:     product.Image,
			Qty:       product.Qty,
			CreatedAt: product.CreatedAt,
			UpdatedAt: product.UpdatedAt,
		}

		productResponses = append(productResponses, productResponse)
	}

	return &productResponses
}
