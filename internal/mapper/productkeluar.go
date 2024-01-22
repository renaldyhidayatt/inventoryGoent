package mapper

import (
	"github.com/renaldyhidayatt/inventorygoent/internal/domain/response"
	"github.com/renaldyhidayatt/inventorygoent/pkg/database/postgres/ent"
)

type productKeluarMapper struct{}

func NewProductKeluarMapper() *productKeluarMapper {
	return &productKeluarMapper{}
}

func (m *productKeluarMapper) ToProductKeluarResponse(productKeluar *ent.ProductKeluar) *response.ProductKeluar {

	var categoryResponses []response.Category

	for _, category := range productKeluar.Edges.Category {
		categoryResponse := response.Category{
			Name:      category.Name,
			CreatedAt: category.CreatedAt,
			UpdatedAt: category.UpdatedAt,
		}

		categoryResponses = append(categoryResponses, categoryResponse)
	}

	var productResponses []response.Product

	for _, product := range productKeluar.Edges.Products {
		productResponse := response.Product{
			Name:      product.Name,
			Image:     product.Image,
			Qty:       product.Qty,
			CreatedAt: product.CreatedAt,
			UpdatedAt: product.UpdatedAt,
		}
		productResponses = append(productResponses, productResponse)
	}

	return &response.ProductKeluar{
		Qty:       productKeluar.Qty,
		Products:  productResponses,
		Category:  categoryResponses[0],
		CreatedAt: productKeluar.CreatedAt,
		UpdatedAt: productKeluar.UpdatedAt,
	}
}

func (m *productKeluarMapper) ToProductKeluarResponses(productKeluars []*ent.ProductKeluar) *[]response.ProductKeluar {
	var productKeluarResponses []response.ProductKeluar

	for _, productKeluar := range productKeluars {
		productKeluarResponse := m.ToProductKeluarResponse(productKeluar)

		productKeluarResponses = append(productKeluarResponses, *productKeluarResponse)
	}

	return &productKeluarResponses
}
