package interfaces

import (
	"github.com/renaldyhidayatt/inventorygoent/dto/request"
	"github.com/renaldyhidayatt/inventorygoent/ent"
)

type ICategoryRepository interface {
	Create(input request.CategoryRequest) (*ent.Category, error)
	Result(id int) (*ent.Category, error)
	Results() ([]*ent.Category, error)
	UpdateById(id int, input request.CategoryRequest) (*ent.Category, error)
	DeleteById(id int) (bool, error)
}

type ICategoryService interface {
	Create(input request.CategoryRequest) (*ent.Category, error)
	Result(id int) (*ent.Category, error)
	Results() ([]*ent.Category, error)
	UpdateById(id int, input request.CategoryRequest) (*ent.Category, error)
	DeleteById(id int) (bool, error)
}
