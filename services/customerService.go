package services

import (
	"github.com/renaldyhidayatt/inventorygoent/dto/request"
	"github.com/renaldyhidayatt/inventorygoent/ent"
	"github.com/renaldyhidayatt/inventorygoent/interfaces"
	"github.com/renaldyhidayatt/inventorygoent/repository"
)

type CustomerService = interfaces.ICustomerService

type customerService struct {
	repository repository.CustomerRepository
}

func NewCustomerService(repository repository.CustomerRepository) *customerService {
	return &customerService{repository: repository}
}

func (s *customerService) Results() ([]*ent.Customer, error) {
	res, err := s.repository.Results()

	return res, err
}

func (s *customerService) Result(id int) (*ent.Customer, error) {
	res, err := s.repository.Result(id)

	return res, err
}

func (s *customerService) Create(input request.CustomerRequest) (*ent.Customer, error) {
	var customerRequest request.CustomerRequest

	customerRequest.Name = input.Name
	customerRequest.Telepon = input.Telepon
	customerRequest.Alamat = input.Alamat

	res, err := s.repository.Create(customerRequest)

	return res, err
}

func (s *customerService) UpdateById(id int, input request.CustomerRequest) (*ent.Customer, error) {
	var customerRequest request.CustomerRequest

	customerRequest.Name = input.Name
	customerRequest.Telepon = input.Telepon
	customerRequest.Alamat = input.Alamat

	res, err := s.repository.UpdateById(id, customerRequest)

	return res, err
}

func (s *customerService) DeleteById(id int) (bool, error) {
	res, err := s.repository.DeleteById(id)

	return res, err
}
