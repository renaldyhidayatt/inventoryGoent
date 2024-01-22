package repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/renaldyhidayatt/inventorygoent/internal/domain/request"
	"github.com/renaldyhidayatt/inventorygoent/pkg/database/postgres/ent"
	"github.com/renaldyhidayatt/inventorygoent/pkg/database/postgres/ent/category"
	"github.com/renaldyhidayatt/inventorygoent/pkg/database/postgres/ent/product"
	"github.com/renaldyhidayatt/inventorygoent/pkg/database/postgres/ent/productkeluar"
)

type productKeluarRepository struct {
	db      *ent.Client
	context context.Context
}

func NewProductKeluarRepository(db *ent.Client, ctx context.Context) *productKeluarRepository {
	return &productKeluarRepository{
		db:      db,
		context: ctx,
	}
}

func (r *productKeluarRepository) GetAll() ([]*ent.ProductKeluar, error) {
	res, err := r.db.ProductKeluar.Query().WithCategory().WithProducts().All(r.context)

	if err != nil {
		return nil, errors.New("failed to get product keluars")
	}

	return res, nil
}

func (r *productKeluarRepository) GetById(id int) (*ent.ProductKeluar, error) {
	res, err := r.db.ProductKeluar.Query().Where(productkeluar.IDEQ(id)).WithCategory().WithProducts().First(r.context)

	if err != nil {
		return nil, errors.New("failed to get product keluar")
	}

	return res, nil
}

func (r *productKeluarRepository) Create(request request.CreateProductKeluarRequest) (*ent.ProductKeluar, error) {
	product, err := r.db.Product.Query().Where(product.IDEQ(request.ProductID)).First(r.context)

	if err != nil {
		return nil, fmt.Errorf("failed query get product : %w", err)
	}

	category, err := r.db.Category.Query().Where(category.IDEQ(request.CategoryID)).First(r.context)

	if err != nil {
		return nil, fmt.Errorf("failed query get category : %w", err)
	}

	res, err := r.db.ProductKeluar.Create().SetQty(request.Qty).AddProducts(product).AddCategory(category).Save(r.context)

	if err != nil {
		return nil, fmt.Errorf("failed query create productkeluar : %w", err)
	}

	return res, nil

}

func (r *productKeluarRepository) Update(request request.UpdateProductKeluarRequest) (*ent.ProductKeluar, error) {

	_, err := r.db.ProductKeluar.Query().Where(productkeluar.IDEQ(request.ID)).First(r.context)

	if err != nil {
		return nil, errors.New("failed to get product keluar")
	}

	product, err := r.db.Product.Query().Where(product.IDEQ(request.ProductID)).First(r.context)

	if err != nil {
		return nil, fmt.Errorf("failed query get product : %w", err)
	}

	category, err := r.db.Category.Query().Where(category.IDEQ(request.CategoryID)).First(r.context)

	if err != nil {
		return nil, fmt.Errorf("failed query get category : %w", err)
	}

	res, err := r.db.ProductKeluar.UpdateOneID(request.ID).SetQty(request.Qty).AddProducts(product).AddCategory(category).Save(r.context)

	if err != nil {
		return nil, fmt.Errorf("failed query update productkeluar : %w", err)
	}

	return res, nil
}

func (r *productKeluarRepository) Delete(id int) error {
	_, err := r.db.ProductKeluar.Delete().Where(productkeluar.IDEQ(id)).Exec(r.context)

	if err != nil {
		return errors.New("failed to delete product keluar")
	}

	return nil
}
