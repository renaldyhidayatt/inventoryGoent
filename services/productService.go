package services

import (
	"github.com/renaldyhidayatt/inventorygoent/dto/request"
	"github.com/renaldyhidayatt/inventorygoent/ent"
	"github.com/renaldyhidayatt/inventorygoent/interfaces"
	"github.com/renaldyhidayatt/inventorygoent/repository"
)

type ProductService = interfaces.IProductService

type productService struct {
	repository repository.ProductRepository
}

func NewProductService(repository repository.ProductRepository) *productService {
	return &productService{repository: repository}
}

func (s *productService) Results() ([]*ent.Product, error) {
	res, err := s.repository.Results()

	return res, err
}

func (s *productService) Result(id int) (*ent.Product, error) {
	res, err := s.repository.Result(id)

	return res, err
}

func (s *productService) Create(input request.ProductRequest) (*ent.Product, error) {
	var productRequest request.ProductRequest

	productRequest.Name = input.Name
	productRequest.Image = input.Image
	productRequest.Qty = input.Qty
	productRequest.CategoryID = input.CategoryID

	res, err := s.repository.Create(input)

	return res, err

}
func (s *productService) UpdateById(id int, input request.ProductRequest) (*ent.Product, error) {
	var productRequest request.ProductRequest

	productRequest.Name = input.Name
	productRequest.Image = input.Image
	productRequest.Qty = input.Qty
	productRequest.CategoryID = input.CategoryID

	res, err := s.repository.UpdateById(id, input)

	return res, err

}

func (s *productService) DeleteById(id int) (bool, error) {
	res, err := s.repository.DeleteById(id)

	return res, err
}
