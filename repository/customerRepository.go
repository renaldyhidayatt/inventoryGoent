package repository

import (
	"context"
	"fmt"

	"github.com/renaldyhidayatt/inventorygoent/dto/request"
	"github.com/renaldyhidayatt/inventorygoent/ent"
	"github.com/renaldyhidayatt/inventorygoent/ent/customer"
	"github.com/renaldyhidayatt/inventorygoent/interfaces"
)

type CustomerRepository = interfaces.ICustomerRepository

type customerRepository struct {
	db      *ent.Client
	context context.Context
}

func NewCustomerRepository(db *ent.Client, context context.Context) *customerRepository {
	return &customerRepository{db: db, context: context}
}

func (r *customerRepository) Results() ([]*ent.Customer, error) {
	customer, err := r.db.Customer.Query().All(r.context)

	if err != nil {
		return nil, fmt.Errorf("failed results customers: %w", err)
	}

	return customer, nil
}

func (r *customerRepository) Result(id int) (*ent.Customer, error) {
	customer, err := r.db.Customer.Query().Where(customer.IDEQ(id)).First(r.context)

	if err != nil {
		return nil, fmt.Errorf("failed result customers: %w", err)
	}

	return customer, nil
}

func (r *customerRepository) Create(input request.CustomerRequest) (*ent.Customer, error) {
	res, err := r.db.Customer.Create().SetName(input.Name).SetAlamat(input.Alamat).SetTelepon(input.Telepon).Save(r.context)

	if err != nil {
		return nil, fmt.Errorf("failed create customer: %w", err)
	}

	return res, nil
}

func (r *customerRepository) UpdateById(id int, input request.CustomerRequest) (*ent.Customer, error) {
	_, err := r.db.Customer.Query().Where(customer.IDEQ(id)).First(r.context)

	if err != nil {
		return nil, fmt.Errorf("failed result customers: %w", err)
	}
	customer, err := r.db.Customer.UpdateOneID(id).SetName(input.Name).SetAlamat(input.Alamat).SetTelepon(input.Telepon).Save(r.context)

	if err != nil {
		return nil, fmt.Errorf("failed update customer: %w", err)
	}

	return customer, nil
}

func (r *customerRepository) DeleteById(id int) (bool, error) {
	_, err := r.db.Customer.Query().Where(customer.IDEQ(id)).First(r.context)

	if err != nil {
		return false, fmt.Errorf("failed result customers: %w", err)
	}

	err = r.db.Customer.DeleteOneID(id).Exec(r.context)

	if err != nil {
		return false, fmt.Errorf("failed  delete customer: %w", err)
	}
	return true, nil

}
