package repository

import (
	"context"
	"errors"
	"time"

	"github.com/renaldyhidayatt/inventorygoent/internal/domain/request"
	"github.com/renaldyhidayatt/inventorygoent/pkg/database/postgres/ent"
	"github.com/renaldyhidayatt/inventorygoent/pkg/database/postgres/ent/customer"
)

type customerRepository struct {
	db      *ent.Client
	context context.Context
}

func NewCustomerRepository(db *ent.Client, ctx context.Context) *customerRepository {
	return &customerRepository{
		db:      db,
		context: ctx,
	}
}

func (r *customerRepository) GetAll() ([]*ent.Customer, error) {
	res, err := r.db.Customer.Query().All(r.context)

	if err != nil {
		return nil, errors.New("failed to get customers")
	}

	return res, nil
}

func (r *customerRepository) GetById(id int) (*ent.Customer, error) {
	res, err := r.db.Customer.Query().Where(customer.IDEQ(id)).First(r.context)

	if err != nil {
		return nil, errors.New("failed to get customer")
	}

	return res, nil
}

func (r *customerRepository) GetByName(name string) (*ent.Customer, error) {
	res, err := r.db.Customer.Query().Where(customer.NameEQ(name)).First(r.context)

	if err != nil {
		return nil, errors.New("failed to get customer")
	}

	return res, nil
}

func (r *customerRepository) Create(request request.CreateCustomerRequest) (*ent.Customer, error) {
	_, err := r.db.Customer.Query().Where(customer.NameEQ(request.Name)).First(r.context)
	if err == nil {
		return nil, errors.New("customer already exists")
	} else if !ent.IsNotFound(err) {
		return nil, errors.New("failed to check customer existence")
	}

	newCustomer, err := r.db.Customer.Create().SetName(request.Name).SetAlamat(request.Alamat).SetTelepon(request.Telepon).SetCreatedAt(time.Now()).Save(r.context)

	if err != nil {
		return nil, errors.New("failed to create customer")
	}

	return newCustomer, nil
}

func (r *customerRepository) Update(request request.UpdateCustomerRequest) (*ent.Customer, error) {
	_, err := r.db.Customer.Query().Where(customer.NameEQ(request.Name)).First(r.context)
	if err == nil {
		return nil, errors.New("customer already exists")
	} else if !ent.IsNotFound(err) {
		return nil, errors.New("failed to check customer existence")
	}

	updatedCustomer, err := r.db.Customer.UpdateOneID(request.ID).SetName(request.Name).SetAlamat(request.Alamat).SetTelepon(request.Telepon).Save(r.context)

	if err != nil {
		return nil, errors.New("failed to update customer")
	}

	return updatedCustomer, nil
}

func (r *customerRepository) Delete(id int) error {
	_, err := r.db.Customer.Delete().Where(customer.IDEQ(id)).Exec(r.context)

	if err != nil {
		return errors.New("failed to delete customer")
	}

	return nil
}
