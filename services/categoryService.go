package services

import (
	"github.com/renaldyhidayatt/inventorygoent/dto/request"
	"github.com/renaldyhidayatt/inventorygoent/ent"
	"github.com/renaldyhidayatt/inventorygoent/interfaces"
	"github.com/renaldyhidayatt/inventorygoent/repository"
)

type CategoryService = interfaces.ICategoryService

type categoryService struct {
	repository repository.CategoryRepository
}

func NewCategoryService(repository repository.CategoryRepository) *categoryService {
	return &categoryService{repository: repository}
}

func (s *categoryService) Results() ([]*ent.Category, error) {
	res, err := s.repository.Results()

	return res, err
}

func (s *categoryService) Result(id int) (*ent.Category, error) {
	res, err := s.repository.Result(id)

	return res, err
}

func (s *categoryService) Create(input request.CategoryRequest) (*ent.Category, error) {
	var categoryRequest request.CategoryRequest

	categoryRequest.Name = input.Name

	res, err := s.repository.Create(categoryRequest)

	return res, err
}

func (s *categoryService) UpdateById(id int, input request.CategoryRequest) (*ent.Category, error) {
	var categoryRequest request.CategoryRequest

	categoryRequest.Name = input.Name

	res, err := s.repository.UpdateById(id, categoryRequest)

	return res, err
}

func (s *categoryService) DeleteById(id int) (bool, error) {
	res, err := s.repository.DeleteById(id)

	return res, err
}
