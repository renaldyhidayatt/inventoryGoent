package services

import (
	"github.com/renaldyhidayatt/inventorygoent/dto/request"
	"github.com/renaldyhidayatt/inventorygoent/ent"
	"github.com/renaldyhidayatt/inventorygoent/repository"
)

type ProductKeluarService = repository.ProductKeluarRepository

type productKeluarService struct {
	repository repository.ProductKeluarRepository
}

func NewProductKeluarService(repository repository.ProductKeluarRepository) *productKeluarService {
	return &productKeluarService{repository: repository}
}

func (s *productKeluarService) Results() ([]*ent.ProductKeluar, error) {
	res, err := s.repository.Results()

	return res, err
}

func (s *productKeluarService) Result(id int) (*ent.ProductKeluar, error) {
	res, err := s.repository.Result(id)

	return res, err
}

func (s *productKeluarService) Create(input request.ProductKeluarRequest) (*ent.ProductKeluar, error) {
	var productKeluarRequest request.ProductKeluarRequest

	productKeluarRequest.Qty = input.Qty
	productKeluarRequest.CategoryID = input.CategoryID
	productKeluarRequest.ProductID = input.ProductID

	res, err := s.repository.Create(input)

	return res, err
}

func (s *productKeluarService) UpdateById(id int, input request.ProductKeluarRequest) (*ent.ProductKeluar, error) {
	var productKeluarRequest request.ProductKeluarRequest

	productKeluarRequest.Qty = input.Qty
	productKeluarRequest.CategoryID = input.CategoryID
	productKeluarRequest.ProductID = input.ProductID

	res, err := s.repository.UpdateById(id, input)

	return res, err
}

func (s *productKeluarService) DeleteById(id int) (bool, error) {
	res, err := s.repository.DeleteById(id)

	return res, err
}
