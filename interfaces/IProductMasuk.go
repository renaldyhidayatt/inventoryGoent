package interfaces

import (
	"github.com/renaldyhidayatt/inventorygoent/dto/request"
	"github.com/renaldyhidayatt/inventorygoent/ent"
)

type IProductMasukRepository interface {
	Create(input request.ProductMasukRequest) (*ent.ProductMasuk, error)
	Result(id int) (*ent.ProductMasuk, error)
	Results() ([]*ent.ProductMasuk, error)
	DeleteById(id int) (bool, error)
	UpdateById(id int, input request.ProductMasukRequest) (*ent.ProductMasuk, error)
}
type IProductMasukService interface {
	Create(input request.ProductMasukRequest) (*ent.ProductMasuk, error)
	Result(id int) (*ent.ProductMasuk, error)
	Results() ([]*ent.ProductMasuk, error)
	DeleteById(id int) (bool, error)
	UpdateById(id int, input request.ProductMasukRequest) (*ent.ProductMasuk, error)
}
