package services

import (
	"github.com/renaldyhidayatt/inventorygoent/dto/request"
	"github.com/renaldyhidayatt/inventorygoent/ent"
	"github.com/renaldyhidayatt/inventorygoent/interfaces"
	"github.com/renaldyhidayatt/inventorygoent/repository"
)

type SupplierService = interfaces.ISupplierService

type supplierService struct {
	repository repository.SupplierRepository
}

func NewSupplierService(repository repository.SupplierRepository) *supplierService {
	return &supplierService{repository: repository}
}

func (s *supplierService) Results() ([]*ent.Supplier, error) {
	res, err := s.repository.Results()

	return res, err
}

func (s *supplierService) Result(id int) (*ent.Supplier, error) {
	res, err := s.repository.Result(id)

	return res, err
}

func (s *supplierService) Create(input request.SupplierRequest) (*ent.Supplier, error) {
	var supplierRequest request.SupplierRequest

	supplierRequest.Name = input.Name
	supplierRequest.Telepon = input.Telepon
	supplierRequest.Alamat = input.Alamat

	res, err := s.repository.Create(supplierRequest)

	return res, err
}

func (s *supplierService) UpdateById(id int, input request.SupplierRequest) (*ent.Supplier, error) {
	var supplierRequest request.SupplierRequest

	supplierRequest.Name = input.Name
	supplierRequest.Telepon = input.Telepon
	supplierRequest.Alamat = input.Alamat

	res, err := s.repository.UpdateById(id, supplierRequest)

	return res, err
}

func (s *supplierService) DeleteById(id int) (bool, error) {
	res, err := s.repository.DeleteById(id)

	return res, err
}
