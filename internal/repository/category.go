package repository

import (
	"context"
	"errors"
	"time"

	"github.com/renaldyhidayatt/inventorygoent/internal/domain/request"
	"github.com/renaldyhidayatt/inventorygoent/pkg/database/postgres/ent"
	"github.com/renaldyhidayatt/inventorygoent/pkg/database/postgres/ent/category"
)

type categoryRepository struct {
	db      *ent.Client
	context context.Context
}

func NewCategoryRepository(db *ent.Client, ctx context.Context) *categoryRepository {
	return &categoryRepository{
		db:      db,
		context: ctx,
	}
}

func (r *categoryRepository) GetAll() ([]*ent.Category, error) {
	res, err := r.db.Category.Query().All(r.context)

	if err != nil {
		return nil, errors.New("failed to get categories")
	}

	return res, nil
}

func (r *categoryRepository) GetByName(name string) (*ent.Category, error) {
	res, err := r.db.Category.Query().Where(category.NameEQ(name)).First(r.context)

	if err != nil {
		return nil, errors.New("failed to get category")
	}

	return res, nil
}

func (r *categoryRepository) GetById(id int) (*ent.Category, error) {
	res, err := r.db.Category.Query().Where(category.IDEQ(id)).First(r.context)

	if err != nil {
		return nil, errors.New("failed to get category")
	}

	return res, nil
}

func (r *categoryRepository) Create(request request.CreateCategoryRequest) (*ent.Category, error) {
	_, err := r.db.Category.Query().Where(category.NameEQ(request.Name)).First(r.context)
	if err == nil {
		return nil, errors.New("category already exists")
	} else if !ent.IsNotFound(err) {
		return nil, errors.New("failed to check category existence")
	}

	newCategory, err := r.db.Category.Create().SetName(request.Name).SetCreatedAt(time.Now()).Save(r.context)

	if err != nil {
		return nil, errors.New("failed to create category")
	}

	return newCategory, nil
}

func (r *categoryRepository) Update(request request.UpdateCategoryRequest) (*ent.Category, error) {
	_, err := r.db.Category.Query().Where(category.NameEQ(request.Name)).First(r.context)
	if err == nil {
		return nil, errors.New("category already exists")
	} else if !ent.IsNotFound(err) {
		return nil, errors.New("failed to check category existence")
	}

	updatedCategory, err := r.db.Category.UpdateOneID(request.ID).SetName(request.Name).Save(r.context)

	if err != nil {
		return nil, errors.New("failed to update category")
	}

	return updatedCategory, nil
}

func (r *categoryRepository) Delete(id int) error {
	_, err := r.db.Category.Delete().Where(category.IDEQ(id)).Exec(r.context)

	if err != nil {
		return errors.New("failed to delete category")
	}

	return nil
}
