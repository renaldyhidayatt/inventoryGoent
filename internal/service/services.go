package service

import (
	"github.com/renaldyhidayatt/inventorygoent/internal/mapper"
	"github.com/renaldyhidayatt/inventorygoent/internal/repository"
	"github.com/renaldyhidayatt/inventorygoent/pkg/auth"
	"github.com/renaldyhidayatt/inventorygoent/pkg/hash"
	"github.com/renaldyhidayatt/inventorygoent/pkg/logger"
)

type Services struct {
	Auth          AuthService
	Category      CategoryService
	Customer      CustomerService
	Product       ProductService
	ProductKeluar ProductKeluarService
	ProductMasuk  ProductMasukService
	Supplier      SupplierService
}

type Deps struct {
	Repository *repository.Repositories
	Logger     *logger.Logger
	Hash       *hash.Hashing
	Token      auth.TokenManager
	Mapper     *mapper.Mapper
}

func NewServices(deps Deps) *Services {
	return &Services{
		Auth:          NewAuthService(*deps.Hash, deps.Token, deps.Repository.Auth, *deps.Logger),
		Category:      NewCategoryService(*deps.Logger, deps.Repository.Category, deps.Mapper.CategoryMapper),
		Customer:      NewCustomerService(*deps.Logger, deps.Repository.Customer, deps.Mapper.CustomerMapper),
		Product:       NewProductService(*deps.Logger, deps.Repository.Product, deps.Mapper.ProductMapper),
		ProductKeluar: NewProductKeluarService(*deps.Logger, deps.Repository.ProductKeluar, deps.Mapper.ProductKeluarMapper),
		ProductMasuk:  NewProductMasukService(*deps.Logger, deps.Repository.ProductMasuk, deps.Mapper.ProductMasukMapper),
		Supplier:      NewSupplierService(*deps.Logger, deps.Repository.Supplier, deps.Mapper.SupplierMapper),
	}
}
