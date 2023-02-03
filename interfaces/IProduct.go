package interfaces

import (
	"github.com/renaldyhidayatt/inventorygoent/dto/request"
	"github.com/renaldyhidayatt/inventorygoent/ent"
)

type IProductRepository interface {
	Create(input request.ProductRequest) (*ent.Product, error)
	Result(id int) (*ent.Product, error)
	Results() ([]*ent.Product, error)
	DeleteById(id int) (bool, error)
	UpdateById(id int, input request.ProductRequest) (*ent.Product, error)
}

type IProductService interface {
	Create(input request.ProductRequest) (*ent.Product, error)
	Result(id int) (*ent.Product, error)
	Results() ([]*ent.Product, error)
	DeleteById(id int) (bool, error)
	UpdateById(id int, input request.ProductRequest) (*ent.Product, error)
}
