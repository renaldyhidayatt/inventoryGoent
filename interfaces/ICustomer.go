package interfaces

import (
	"github.com/renaldyhidayatt/inventorygoent/dto/request"
	"github.com/renaldyhidayatt/inventorygoent/ent"
)

type ICustomerRepository interface {
	Result(id int) (*ent.Customer, error)
	Results() ([]*ent.Customer, error)
	Create(input request.CustomerRequest) (*ent.Customer, error)
	DeleteById(id int) (bool, error)
	UpdateById(id int, input request.CustomerRequest) (*ent.Customer, error)
}

type ICustomerService interface {
	Result(id int) (*ent.Customer, error)
	Results() ([]*ent.Customer, error)
	Create(input request.CustomerRequest) (*ent.Customer, error)
	DeleteById(id int) (bool, error)
	UpdateById(id int, input request.CustomerRequest) (*ent.Customer, error)
}
