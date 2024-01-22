package repository

import (
	"github.com/renaldyhidayatt/inventorygoent/internal/domain/request"
	"github.com/renaldyhidayatt/inventorygoent/pkg/database/postgres/ent"
)

type AuthRepository interface {
	Register(requests request.RegisterRequest) (*ent.User, error)
	Login(requests request.AuthRequest) (*ent.User, error)
}

type CategoryRepository interface {
	GetAll() ([]*ent.Category, error)
	GetByName(name string) (*ent.Category, error)
	GetById(id int) (*ent.Category, error)
	Create(request request.CreateCategoryRequest) (*ent.Category, error)
	Update(request request.UpdateCategoryRequest) (*ent.Category, error)
	Delete(id int) error
}

type CustomerRepository interface {
	GetAll() ([]*ent.Customer, error)
	GetById(id int) (*ent.Customer, error)
	GetByName(name string) (*ent.Customer, error)
	Create(request request.CreateCustomerRequest) (*ent.Customer, error)
	Update(request request.UpdateCustomerRequest) (*ent.Customer, error)
	Delete(id int) error
}

type ProductRepository interface {
	GetAll() ([]*ent.Product, error)
	GetById(id int) (*ent.Product, error)
	GetByName(name string) (*ent.Product, error)
	Create(request request.CreateProductRequest) (*ent.Product, error)
	Update(request request.UpdateProductRequest) (*ent.Product, error)
	Delete(id int) error
}

type ProductMasukRepository interface {
	GetAll() ([]*ent.ProductMasuk, error)
	GetById(id int) (*ent.ProductMasuk, error)
	Create(request request.CreateProductMasukRequest) (*ent.ProductMasuk, error)
	Update(request request.UpdateProductMasukRequest) (*ent.ProductMasuk, error)
	Delete(id int) error
}

type ProductKeluarRepository interface {
	GetAll() ([]*ent.ProductKeluar, error)
	GetById(id int) (*ent.ProductKeluar, error)
	Create(request request.CreateProductKeluarRequest) (*ent.ProductKeluar, error)
	Update(request request.UpdateProductKeluarRequest) (*ent.ProductKeluar, error)
	Delete(id int) error
}

type SupplierRepository interface {
	GetAll() ([]*ent.Supplier, error)
	GetById(id int) (*ent.Supplier, error)
	Create(request request.CreateSupplierRequest) (*ent.Supplier, error)
	Update(request request.UpdateSupplierRequest) (*ent.Supplier, error)
	Delete(id int) error
}
