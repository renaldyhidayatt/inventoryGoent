package repository

import (
	"context"
	"fmt"

	"github.com/renaldyhidayatt/inventorygoent/dto/request"
	"github.com/renaldyhidayatt/inventorygoent/ent"
	"github.com/renaldyhidayatt/inventorygoent/ent/product"
	"github.com/renaldyhidayatt/inventorygoent/ent/productmasuk"
	"github.com/renaldyhidayatt/inventorygoent/ent/supplier"
	"github.com/renaldyhidayatt/inventorygoent/interfaces"
)

type ProductMasukRepository = interfaces.IProductMasukRepository

type productMasukRepository struct {
	db      *ent.Client
	context context.Context
}

func NewProductMasukRepository(db *ent.Client, context context.Context) *productMasukRepository {
	return &productMasukRepository{db: db, context: context}
}

func (r *productMasukRepository) Results() ([]*ent.ProductMasuk, error) {
	product, err := r.db.ProductMasuk.Query().WithProduct().WithSupplier().All(r.context)

	if err != nil {
		return nil, fmt.Errorf("failed query results productmasuk: %w", err)
	}

	return product, nil

}

func (r *productMasukRepository) Result(id int) (*ent.ProductMasuk, error) {
	product, err := r.db.ProductMasuk.Query().Where(productmasuk.IDEQ(id)).WithProduct().WithSupplier().First(r.context)

	if err != nil {
		return nil, fmt.Errorf("failed query result productmasuk: %w", err)
	}

	return product, nil
}

func (r *productMasukRepository) Create(input request.ProductMasukRequest) (*ent.ProductMasuk, error) {
	product, err := r.db.Product.Query().Where(product.IDEQ(input.ProductID)).First(r.context)

	if err != nil {
		return nil, fmt.Errorf("failed query product by id: %w", err)
	}

	supplier, err := r.db.Supplier.Query().Where(supplier.IDEQ(input.SupplierID)).First(r.context)

	if err != nil {
		return nil, fmt.Errorf("failed query supplier by id: %w", err)
	}

	res, err := r.db.ProductMasuk.Create().SetName(input.Name).SetQty(input.Qty).AddProduct(product).AddSupplier(supplier).Save(r.context)

	if err != nil {
		return nil, fmt.Errorf("failed query create productmasuk: %w", err)
	}

	return res, nil
}

func (r *productMasukRepository) UpdateById(id int, input request.ProductMasukRequest) (*ent.ProductMasuk, error) {
	_, err := r.db.ProductMasuk.Query().Where(productmasuk.IDEQ(id)).First(r.context)

	if err != nil {
		return nil, fmt.Errorf("failed query result productmasuk: %w", err)
	}

	product, err := r.db.Product.Query().Where(product.IDEQ(input.ProductID)).First(r.context)

	if err != nil {
		return nil, fmt.Errorf("failed query product by id: %w", err)
	}

	supplier, err := r.db.Supplier.Query().Where(supplier.IDEQ(input.SupplierID)).First(r.context)

	if err != nil {
		return nil, fmt.Errorf("failed query supplier by id: %w", err)
	}

	res, err := r.db.ProductMasuk.UpdateOneID(id).SetName(input.Name).SetQty(input.Qty).AddProduct(product).AddSupplier(supplier).Save(r.context)

	if err != nil {
		return nil, fmt.Errorf("failed query productmasuk by id: %w", err)
	}

	return res, nil
}

func (r *productMasukRepository) DeleteById(id int) (bool, error) {
	_, err := r.db.ProductMasuk.Query().Where(productmasuk.IDEQ(id)).First(r.context)

	if err != nil {
		return false, fmt.Errorf("failed query product masuk by id: %w", err)
	}

	err = r.db.Category.DeleteOneID(id).Exec(r.context)

	if err != nil {
		return false, fmt.Errorf("failed query delete; %w", err)
	}

	return true, nil
}
