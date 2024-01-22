package service

import (
	"github.com/renaldyhidayatt/inventorygoent/internal/domain/request"
	"github.com/renaldyhidayatt/inventorygoent/internal/domain/response"
	"github.com/renaldyhidayatt/inventorygoent/pkg/database/postgres/ent"
)

type AuthService interface {
	Register(requests request.RegisterRequest) (*ent.User, error)
	Login(request request.AuthRequest) (string, error)
}

type CategoryService interface {
	GetAll() (*[]response.Category, error)
	GetById(id int) (*response.Category, error)
	Create(request request.CreateCategoryRequest) (*response.Category, error)
	Update(request request.UpdateCategoryRequest) (*response.Category, error)
	Delete(id int) error
}

type CustomerService interface {
	GetAll() (*[]response.Customer, error)
	GetById(id int) (*response.Customer, error)
	FindByName(name string) (*response.Customer, error)
	Create(request request.CreateCustomerRequest) (*response.Customer, error)
	Update(request request.UpdateCustomerRequest) (*response.Customer, error)
	Delete(id int) error
}

type ProductService interface {
	GetAll() (*[]response.Product, error)
	GetById(id int) (*response.Product, error)
	Create(request request.CreateProductRequest) (*response.Product, error)
	Update(request request.UpdateProductRequest) (*response.Product, error)
	Delete(id int) error
}

type ProductKeluarService interface {
	GetAll() (*[]response.ProductKeluar, error)
	GetById(id int) (*response.ProductKeluar, error)
	Create(requests request.CreateProductKeluarRequest) (*response.ProductKeluar, error)
	Update(requests request.UpdateProductKeluarRequest) (*response.ProductKeluar, error)
	Delete(id int) error
}

type ProductMasukService interface {
	GetAll() (*[]response.ProductMasuk, error)
	GetById(id int) (*response.ProductMasuk, error)
	Create(requests request.CreateProductMasukRequest) (*response.ProductMasuk, error)
	Update(requests request.UpdateProductMasukRequest) (*response.ProductMasuk, error)
	Delete(id int) error
}

type SupplierService interface {
	GetAll() (*[]response.Supplier, error)
	GetByID(id int) (*response.Supplier, error)
	Create(requests request.CreateSupplierRequest) (*response.Supplier, error)
	Update(requests request.UpdateSupplierRequest) (*response.Supplier, error)
	Delete(id int) error
}
