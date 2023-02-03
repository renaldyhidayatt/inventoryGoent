package interfaces

import (
	"github.com/renaldyhidayatt/inventorygoent/dto/request"
	"github.com/renaldyhidayatt/inventorygoent/ent"
)

type IProductKeluarRepository interface {
	Create(input request.ProductKeluarRequest) (*ent.ProductKeluar, error)
	Result(id int) (*ent.ProductKeluar, error)
	Results() ([]*ent.ProductKeluar, error)
	DeleteById(id int) (bool, error)
	UpdateById(id int, input request.ProductKeluarRequest) (*ent.ProductKeluar, error)
}
type IProductKeluarService interface {
	Create(input request.ProductKeluarRequest) (*ent.ProductKeluar, error)
	Result(id int) (*ent.ProductKeluar, error)
	Results() ([]*ent.ProductKeluar, error)
	DeleteById(id int) (bool, error)
	UpdateById(id int, input request.ProductKeluarRequest) (*ent.ProductKeluar, error)
}
