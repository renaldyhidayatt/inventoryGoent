package repository

import (
	"context"
	"fmt"

	"github.com/renaldyhidayatt/inventorygoent/dto/request"
	"github.com/renaldyhidayatt/inventorygoent/ent"
	"github.com/renaldyhidayatt/inventorygoent/ent/supplier"
	"github.com/renaldyhidayatt/inventorygoent/interfaces"
)

type SupplierRepository = interfaces.ISupplierRepository

type supplierRepository struct {
	db      *ent.Client
	context context.Context
}

func NewSupplierRepository(db *ent.Client, context context.Context) *supplierRepository {
	return &supplierRepository{db: db, context: context}
}

func (r *supplierRepository) Results() ([]*ent.Supplier, error) {
	supplier, err := r.db.Supplier.Query().WithProductmasuk().All(r.context)

	if err != nil {
		return nil, fmt.Errorf("failed results supplier: %w", err)
	}

	return supplier, nil
}

func (r *supplierRepository) Result(id int) (*ent.Supplier, error) {
	supplier, err := r.db.Supplier.Query().Where(supplier.IDEQ(id)).WithProductmasuk().First(r.context)

	if err != nil {
		return nil, fmt.Errorf("failed result supplier: %w", err)
	}

	return supplier, nil
}

func (r *supplierRepository) Create(input request.SupplierRequest) (*ent.Supplier, error) {
	res, err := r.db.Supplier.Create().SetName(input.Name).SetAlamat(input.Alamat).SetTelepon(input.Telepon).Save(r.context)

	if err != nil {
		return nil, fmt.Errorf("failed query supplier: %w", err)
	}

	return res, nil
}

func (r *supplierRepository) UpdateById(id int, input request.SupplierRequest) (*ent.Supplier, error) {
	_, err := r.db.Supplier.Query().Where(supplier.IDEQ(id)).First(r.context)

	if err != nil {
		return nil, fmt.Errorf("failed result supplier: %w", err)
	}
	supplier, err := r.db.Supplier.UpdateOneID(id).SetName(input.Name).SetAlamat(input.Alamat).SetTelepon(input.Telepon).Save(r.context)

	if err != nil {
		return nil, fmt.Errorf("failed update supplier: %w", err)
	}

	return supplier, nil
}

func (r *supplierRepository) DeleteById(id int) (bool, error) {
	_, err := r.db.Supplier.Query().Where(supplier.IDEQ(id)).First(r.context)

	if err != nil {
		return false, fmt.Errorf("failed result supplier: %w", err)
	}

	err = r.db.Supplier.DeleteOneID(id).Exec(r.context)

	if err != nil {
		return false, fmt.Errorf("failed  delete supplier: %w", err)
	}
	return true, nil

}
