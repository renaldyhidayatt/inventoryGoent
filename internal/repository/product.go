package repository

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/renaldyhidayatt/inventorygoent/internal/domain/request"
	"github.com/renaldyhidayatt/inventorygoent/pkg/database/postgres/ent"
	"github.com/renaldyhidayatt/inventorygoent/pkg/database/postgres/ent/category"
	"github.com/renaldyhidayatt/inventorygoent/pkg/database/postgres/ent/product"
)

type productRepository struct {
	db      *ent.Client
	context context.Context
}

func NewProductRepository(db *ent.Client, ctx context.Context) *productRepository {
	return &productRepository{
		db:      db,
		context: ctx,
	}
}

func (r *productRepository) GetAll() ([]*ent.Product, error) {
	res, err := r.db.Product.Query().WithCategory().All(r.context)

	if err != nil {
		return nil, errors.New("failed to get products")
	}

	return res, nil
}

func (r *productRepository) GetById(id int) (*ent.Product, error) {
	res, err := r.db.Product.Query().Where(product.IDEQ(id)).WithCategory().First(r.context)

	if err != nil {
		return nil, errors.New("failed to get product")
	}

	return res, nil
}

func (r *productRepository) GetByName(name string) (*ent.Product, error) {
	res, err := r.db.Product.Query().Where(product.NameEQ(name)).WithCategory().First(r.context)

	if err != nil {
		return nil, errors.New("failed to get product")
	}

	return res, nil
}

func (r *productRepository) Create(request request.CreateProductRequest) (*ent.Product, error) {
	_, err := r.db.Product.Query().Where(product.NameEQ(request.Name)).First(r.context)

	if err != nil {
		return nil, fmt.Errorf("failed query get product : %w", err)
	}

	category := r.db.Category.Query().Where(category.IDEQ(request.CategoryID)).FirstX(r.context)

	if category == nil {
		return nil, errors.New("category not found")
	}

	res, err := r.db.Product.Create().SetName(request.Name).SetImage(request.Image).SetQty(request.Qty).AddCategory(category).SetCreatedAt(time.Now()).Save(r.context)

	if err != nil {
		return nil, fmt.Errorf("failed query create product : %w", err)
	}

	return res, nil
}

func (r *productRepository) Update(request request.UpdateProductRequest) (*ent.Product, error) {
	_, err := r.db.Product.Query().Where(product.IDEQ(request.ID)).First(r.context)

	if err != nil {
		return nil, fmt.Errorf("failed query produvt by id :%w", err)
	}

	category, err := r.db.Category.Query().Where(category.IDEQ(request.CategoryID)).First(r.context)

	if err != nil {
		return nil, fmt.Errorf("failed query category by id: %w", err)
	}

	product, err := r.db.Product.UpdateOneID(request.ID).SetName(request.Name).SetImage(request.Image).SetQty(request.Qty).AddCategory(category).Save(r.context)

	if err != nil {
		return nil, fmt.Errorf("failed update product: %w", err)
	}

	return product, nil
}

func (r *productRepository) Delete(id int) error {
	_, err := r.db.Product.Delete().Where(product.IDEQ(id)).Exec(r.context)

	if err != nil {
		return errors.New("failed to delete product")
	}

	return nil
}
