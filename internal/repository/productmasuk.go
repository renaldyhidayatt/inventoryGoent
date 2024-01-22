package repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/renaldyhidayatt/inventorygoent/internal/domain/request"
	"github.com/renaldyhidayatt/inventorygoent/pkg/database/postgres/ent"
	"github.com/renaldyhidayatt/inventorygoent/pkg/database/postgres/ent/product"
	"github.com/renaldyhidayatt/inventorygoent/pkg/database/postgres/ent/productmasuk"
	"github.com/renaldyhidayatt/inventorygoent/pkg/database/postgres/ent/supplier"
)

type productMasukRepository struct {
	db      *ent.Client
	context context.Context
}

func NewProductMasukRepository(db *ent.Client, ctx context.Context) *productMasukRepository {
	return &productMasukRepository{
		db:      db,
		context: ctx,
	}
}

func (r *productMasukRepository) GetAll() ([]*ent.ProductMasuk, error) {
	res, err := r.db.ProductMasuk.Query().WithProduct().WithSupplier().All(r.context)

	if err != nil {
		return nil, errors.New("failed to get product masuks")
	}

	return res, nil
}

func (r *productMasukRepository) GetById(id int) (*ent.ProductMasuk, error) {
	res, err := r.db.ProductMasuk.Query().Where(productmasuk.IDEQ(id)).WithProduct().WithSupplier().First(r.context)

	if err != nil {
		return nil, errors.New("failed to get product masuk")
	}

	return res, nil
}

func (r *productMasukRepository) Create(request request.CreateProductMasukRequest) (*ent.ProductMasuk, error) {

	product, err := r.db.Product.Query().Where(product.IDEQ(request.ProductID)).First(r.context)

	if err != nil {
		return nil, fmt.Errorf("failed query get product : %w", err)
	}

	supplier, err := r.db.Supplier.Query().Where(supplier.IDEQ(request.SupplierID)).First(r.context)

	if err != nil {
		return nil, fmt.Errorf("failed query get supplier : %w", err)
	}

	res, err := r.db.ProductMasuk.Create().SetQty(request.Qty).AddProduct(product).AddSupplier(supplier).Save(r.context)

	if err != nil {
		return nil, fmt.Errorf("failed query create productmasuk : %w", err)
	}

	return res, nil
}

func (r *productMasukRepository) Update(request request.UpdateProductMasukRequest) (*ent.ProductMasuk, error) {

	_, err := r.db.ProductMasuk.Query().Where(productmasuk.IDEQ(request.ID)).First(r.context)

	if err != nil {
		return nil, errors.New("failed to get product masuk")
	}

	product, err := r.db.Product.Query().Where(product.IDEQ(request.ProductID)).First(r.context)

	if err != nil {
		return nil, fmt.Errorf("failed query get product : %w", err)
	}

	supplier, err := r.db.Supplier.Query().Where(supplier.IDEQ(request.SupplierID)).First(r.context)

	if err != nil {
		return nil, fmt.Errorf("failed query get supplier : %w", err)
	}

	res, err := r.db.ProductMasuk.UpdateOneID(request.ID).SetQty(request.Qty).AddProduct(product).AddSupplier(supplier).Save(r.context)

	if err != nil {
		return nil, fmt.Errorf("failed query update productmasuk : %w", err)
	}

	return res, nil
}

func (r *productMasukRepository) Delete(id int) error {
	_, err := r.db.ProductMasuk.Delete().Where(productmasuk.IDEQ(id)).Exec(r.context)

	if err != nil {
		return errors.New("failed to delete product masuk")
	}

	return nil
}
