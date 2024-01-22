package repository

import (
	"context"
	"errors"
	"time"

	"github.com/renaldyhidayatt/inventorygoent/internal/domain/request"
	"github.com/renaldyhidayatt/inventorygoent/pkg/database/postgres/ent"
	"github.com/renaldyhidayatt/inventorygoent/pkg/database/postgres/ent/supplier"
)

type supplierRepository struct {
	db  *ent.Client
	ctx context.Context
}

func NewSupplierRepository(db *ent.Client, ctx context.Context) *supplierRepository {
	return &supplierRepository{
		db:  db,
		ctx: ctx,
	}
}

func (r *supplierRepository) GetAll() ([]*ent.Supplier, error) {
	res, err := r.db.Supplier.Query().All(r.ctx)

	if err != nil {
		return nil, errors.New("failed to get suppliers")
	}

	return res, nil
}

func (r *supplierRepository) GetById(id int) (*ent.Supplier, error) {
	res, err := r.db.Supplier.Query().Where(supplier.IDEQ(id)).First(r.ctx)

	if err != nil {
		return nil, errors.New("failed to get supplier")
	}

	return res, nil
}

func (r *supplierRepository) Create(request request.CreateSupplierRequest) (*ent.Supplier, error) {
	res, err := r.db.Supplier.Create().SetName(request.Name).SetAlamat(request.Alamat).SetTelepon(request.Telepon).SetCreatedAt(time.Now()).Save(r.ctx)

	if err != nil {
		return nil, errors.New("failed to create supplier")
	}

	return res, nil
}

func (r *supplierRepository) Update(request request.UpdateSupplierRequest) (*ent.Supplier, error) {
	_, err := r.db.Supplier.Query().Where(supplier.IDEQ(request.ID)).First(r.ctx)

	if err != nil {
		return nil, errors.New("failed to get supplier")
	}

	res, err := r.db.Supplier.UpdateOneID(request.ID).SetName(request.Name).SetAlamat(request.Alamat).SetTelepon(request.Telepon).SetUpdatedAt(time.Now()).Save(r.ctx)

	if err != nil {
		return nil, errors.New("failed to update supplier")
	}

	return res, nil
}

func (r *supplierRepository) Delete(id int) error {
	_, err := r.db.Supplier.Delete().Where(supplier.IDEQ(id)).Exec(r.ctx)

	if err != nil {
		return errors.New("failed to delete supplier")
	}

	return nil
}
