package repository

import (
	"context"

	"github.com/renaldyhidayatt/inventorygoent/pkg/database/postgres/ent"
)

type Repositories struct {
	Auth          AuthRepository
	Category      CategoryRepository
	Customer      CustomerRepository
	Product       ProductRepository
	ProductMasuk  ProductMasukRepository
	ProductKeluar ProductKeluarRepository
	Supplier      SupplierRepository
}

func NewRepositories(db *ent.Client, ctx context.Context) *Repositories {
	return &Repositories{
		Auth:          NewAuthRepository(db, ctx),
		Category:      NewCategoryRepository(db, ctx),
		Customer:      NewCustomerRepository(db, ctx),
		Product:       NewProductRepository(db, ctx),
		ProductMasuk:  NewProductMasukRepository(db, ctx),
		ProductKeluar: NewProductKeluarRepository(db, ctx),
		Supplier:      NewSupplierRepository(db, ctx),
	}
}
