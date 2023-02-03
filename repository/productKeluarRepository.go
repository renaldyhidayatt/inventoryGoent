package repository

import (
	"context"
	"fmt"

	"github.com/renaldyhidayatt/inventorygoent/dto/request"
	"github.com/renaldyhidayatt/inventorygoent/ent"
	"github.com/renaldyhidayatt/inventorygoent/ent/category"
	"github.com/renaldyhidayatt/inventorygoent/ent/product"
	"github.com/renaldyhidayatt/inventorygoent/ent/productkeluar"
	"github.com/renaldyhidayatt/inventorygoent/interfaces"
)

type ProductKeluarRepository = interfaces.IProductKeluarRepository

type productKeluarRepository struct {
	db      *ent.Client
	context context.Context
}

func NewProductKeluarRepository(db *ent.Client, context context.Context) *productKeluarRepository {
	return &productKeluarRepository{db: db, context: context}
}

func (r *productKeluarRepository) Results() ([]*ent.ProductKeluar, error) {
	res, err := r.db.ProductKeluar.Query().WithCategory().WithProducts().All(r.context)

	if err != nil {
		return nil, fmt.Errorf("failed query result product keluar :%w", err)
	}

	return res, nil
}

func (r *productKeluarRepository) Result(id int) (*ent.ProductKeluar, error) {
	res, err := r.db.ProductKeluar.Query().Where(productkeluar.IDEQ(id)).WithCategory().WithProducts().First(r.context)

	if err != nil {
		return nil, fmt.Errorf("failed query result product keluar: %w", err)
	}

	return res, nil
}

func (r *productKeluarRepository) Create(input request.ProductKeluarRequest) (*ent.ProductKeluar, error) {
	product, err := r.db.Product.Query().Where(product.IDEQ(input.ProductID)).First(r.context)

	if err != nil {
		return nil, fmt.Errorf("failed query result product : %w", err)
	}

	category, err := r.db.Category.Query().Where(category.IDEQ(input.CategoryID)).First(r.context)

	if err != nil {
		return nil, fmt.Errorf("failed query result category : %w", err)
	}

	res, err := r.db.ProductKeluar.Create().SetQty(input.Qty).AddProducts(product).AddCategory(category).Save(r.context)

	if err != nil {
		return nil, fmt.Errorf("failed query create productkeluar : %w", err)
	}

	return res, nil
}

func (r *productKeluarRepository) UpdateById(id int, input request.ProductKeluarRequest) (*ent.ProductKeluar, error) {
	_, err := r.db.ProductKeluar.Query().Where(productkeluar.IDEQ(id)).First(r.context)

	if err != nil {
		return nil, fmt.Errorf("failed query result product keluar: %w", err)
	}

	product, err := r.db.Product.Query().Where(product.IDEQ(input.ProductID)).First(r.context)

	if err != nil {
		return nil, fmt.Errorf("failed query result product : %w", err)
	}

	category, err := r.db.Category.Query().Where(category.IDEQ(input.CategoryID)).First(r.context)

	if err != nil {
		return nil, fmt.Errorf("failed query result category : %w", err)
	}

	productkeluar, err := r.db.ProductKeluar.UpdateOneID(id).SetQty(input.Qty).AddProducts(product).AddCategory(category).Save(r.context)

	if err != nil {
		return nil, fmt.Errorf("failed query update productkeluar: %w", err)
	}

	return productkeluar, nil
}

func (r *productKeluarRepository) DeleteById(id int) (bool, error) {
	_, err := r.db.ProductKeluar.Query().Where(productkeluar.IDEQ(id)).First(r.context)

	if err != nil {
		return false, fmt.Errorf("failed query result product keluar: %w", err)
	}

	err = r.db.ProductKeluar.DeleteOneID(id).Exec(r.context)

	if err != nil {
		return false, fmt.Errorf("failed query delete: %w", err)
	}

	return true, nil
}
