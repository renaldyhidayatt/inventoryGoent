package repository

import (
	"context"
	"fmt"

	"github.com/renaldyhidayatt/inventorygoent/dto/request"
	"github.com/renaldyhidayatt/inventorygoent/ent"
	"github.com/renaldyhidayatt/inventorygoent/ent/category"
	"github.com/renaldyhidayatt/inventorygoent/ent/product"
	"github.com/renaldyhidayatt/inventorygoent/interfaces"
)

type ProductRepository = interfaces.IProductRepository

type productRepository struct {
	db      *ent.Client
	context context.Context
}

func NewProductRepository(db *ent.Client, context context.Context) *productRepository {
	return &productRepository{db: db, context: context}
}

func (r *productRepository) Results() ([]*ent.Product, error) {
	product, err := r.db.Product.Query().WithCategory().All(r.context)

	if err != nil {
		return nil, fmt.Errorf("failed result product: %w", err)

	}

	return product, nil
}

func (r *productRepository) Result(id int) (*ent.Product, error) {
	product, err := r.db.Product.Query().Where(product.IDEQ(id)).First(r.context)

	if err != nil {
		return nil, fmt.Errorf("failed result product: %w", err)
	}

	return product, nil
}

func (r *productRepository) Create(input request.ProductRequest) (*ent.Product, error) {
	product, err := r.db.Product.Query().Where(product.NameEQ(input.Name)).First(r.context)

	if err != nil {
		return nil, fmt.Errorf("failed query product by name: %w", err)
	}

	if product.ID != 0 {
		return nil, fmt.Errorf("name alread exits")
	}

	category, err := r.db.Category.Query().Where(category.IDEQ(input.CategoryID)).First(r.context)

	if err != nil {
		return nil, fmt.Errorf("failed query category by id: %w", err)
	}

	res, err := r.db.Product.Create().SetName(input.Name).SetImage(input.Image).SetQty(input.Qty).AddCategory(category).Save(r.context)

	if err != nil {
		return nil, fmt.Errorf("failed query product create: %w", err)
	}

	return res, nil
}

func (r *productRepository) UpdateById(id int, input request.ProductRequest) (*ent.Product, error) {
	_, err := r.db.Product.Query().Where(product.IDEQ(id)).First(r.context)

	if err != nil {
		return nil, fmt.Errorf("failed query produvt by id :%w", err)
	}

	category, err := r.db.Category.Query().Where(category.IDEQ(input.CategoryID)).First(r.context)

	if err != nil {
		return nil, fmt.Errorf("failed query category by id: %w", err)
	}

	product, err := r.db.Product.UpdateOneID(id).SetName(input.Name).SetImage(input.Image).SetQty(input.Qty).AddCategory(category).Save(r.context)

	if err != nil {
		return nil, fmt.Errorf("failed update product: %w", err)
	}

	return product, nil

}

func (r *productRepository) DeleteById(id int) (bool, error) {
	_, err := r.db.Product.Query().Where(product.IDEQ(id)).First(r.context)

	if err != nil {
		return false, fmt.Errorf("failed query product by id: %w", err)
	}

	err = r.db.Product.DeleteOneID(id).Exec(r.context)

	if err != nil {
		return false, fmt.Errorf("failed query delete; %w", err)
	}

	return true, nil
}
