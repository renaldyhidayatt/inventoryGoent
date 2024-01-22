package mapper

import (
	"github.com/renaldyhidayatt/inventorygoent/internal/domain/response"
	"github.com/renaldyhidayatt/inventorygoent/pkg/database/postgres/ent"
)

type categoryMapper struct {
}

func NewCategoryMapper() *categoryMapper {
	return &categoryMapper{}
}

func (m *categoryMapper) ToCategoryResponse(category *ent.Category) *response.Category {

	return &response.Category{
		Name:      category.Name,
		CreatedAt: category.CreatedAt,
		UpdatedAt: category.UpdatedAt,
	}
}

func (m *categoryMapper) ToCategoryResponses(category []*ent.Category) *[]response.Category {
	var categoryResponses []response.Category

	for _, category := range category {
		categoryResponse := response.Category{
			Name:      category.Name,
			CreatedAt: category.CreatedAt,
			UpdatedAt: category.UpdatedAt,
		}
		categoryResponses = append(categoryResponses, categoryResponse)
	}

	return &categoryResponses
}
