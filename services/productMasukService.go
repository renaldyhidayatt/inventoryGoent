package services

import (
	"github.com/renaldyhidayatt/inventorygoent/dto/request"
	"github.com/renaldyhidayatt/inventorygoent/ent"
	"github.com/renaldyhidayatt/inventorygoent/interfaces"
	"github.com/renaldyhidayatt/inventorygoent/repository"
)

type ProductMasukService = interfaces.IProductMasukService

type productMasukService struct {
	repository repository.ProductMasukRepository
}

func NewProductMasukService(repository repository.ProductMasukRepository) *productMasukService {
	return &productMasukService{repository: repository}
}

func (s *productMasukService) Results() ([]*ent.ProductMasuk, error) {
	res, err := s.repository.Results()

	return res, err
}

func (s *productMasukService) Result(id int) (*ent.ProductMasuk, error) {
	res, err := s.repository.Result(id)

	return res, err
}

func (s *productMasukService) Create(input request.ProductMasukRequest) (*ent.ProductMasuk, error) {
	var productMasukRequest request.ProductMasukRequest

	productMasukRequest.Name = input.Name
	productMasukRequest.Qty = input.Qty
	productMasukRequest.ProductID = input.ProductID
	productMasukRequest.SupplierID = input.SupplierID

	res, err := s.repository.Create(productMasukRequest)

	return res, err
}

func (s *productMasukService) UpdateById(id int, input request.ProductMasukRequest) (*ent.ProductMasuk, error) {
	var productMasukRequest request.ProductMasukRequest

	productMasukRequest.Name = input.Name
	productMasukRequest.Qty = input.Qty
	productMasukRequest.ProductID = input.ProductID
	productMasukRequest.SupplierID = input.SupplierID

	res, err := s.repository.UpdateById(id, productMasukRequest)

	return res, err
}

func (s *productMasukService) DeleteById(id int) (bool, error) {
	res, err := s.repository.DeleteById(id)

	return res, err
}
