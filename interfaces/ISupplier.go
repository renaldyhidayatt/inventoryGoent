package interfaces

import (
	"github.com/renaldyhidayatt/inventorygoent/dto/request"
	"github.com/renaldyhidayatt/inventorygoent/ent"
)

type ISupplierRepository interface {
	Create(input request.SupplierRequest) (*ent.Supplier, error)
	Result(id int) (*ent.Supplier, error)
	Results() ([]*ent.Supplier, error)
	UpdateById(id int, input request.SupplierRequest) (*ent.Supplier, error)
	DeleteById(id int) (bool, error)
}
type ISupplierService interface {
	Create(input request.SupplierRequest) (*ent.Supplier, error)
	Result(id int) (*ent.Supplier, error)
	Results() ([]*ent.Supplier, error)
	UpdateById(id int, input request.SupplierRequest) (*ent.Supplier, error)
	DeleteById(id int) (bool, error)
}
