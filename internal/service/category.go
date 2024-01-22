package service

import (
	"github.com/renaldyhidayatt/inventorygoent/internal/domain/request"
	"github.com/renaldyhidayatt/inventorygoent/internal/domain/response"
	"github.com/renaldyhidayatt/inventorygoent/internal/mapper"
	"github.com/renaldyhidayatt/inventorygoent/internal/repository"
	"github.com/renaldyhidayatt/inventorygoent/pkg/logger"
	"go.uber.org/zap"
)

type categoryService struct {
	logger             logger.Logger
	categoryRepository repository.CategoryRepository
	mapper             mapper.CategoryMapping
}

func NewCategoryService(logger logger.Logger, categoryRepository repository.CategoryRepository, mapper mapper.CategoryMapping) *categoryService {
	return &categoryService{
		logger:             logger,
		categoryRepository: categoryRepository,
		mapper:             mapper,
	}
}

func (s *categoryService) GetAll() (*[]response.Category, error) {
	categories, err := s.categoryRepository.GetAll()

	if err != nil {
		s.logger.Error("failed to get categories: ", zap.Error(err))
		return nil, err
	}

	categoryEntities := s.mapper.ToCategoryResponses(categories)

	return categoryEntities, nil
}

func (s *categoryService) GetById(id int) (*response.Category, error) {
	category, err := s.categoryRepository.GetById(id)

	if err != nil {
		s.logger.Error("failed to get category: ", zap.Error(err))
		return nil, err
	}

	categoryEntity := s.mapper.ToCategoryResponse(category)

	return categoryEntity, nil
}

func (s *categoryService) Create(request request.CreateCategoryRequest) (*response.Category, error) {
	create, err := s.categoryRepository.Create(request)

	if err != nil {
		s.logger.Error("failed to create category: ", zap.Error(err))
		return nil, err
	}

	categoryEntity := s.mapper.ToCategoryResponse(create)

	return categoryEntity, nil
}

func (s *categoryService) Update(request request.UpdateCategoryRequest) (*response.Category, error) {
	update, err := s.categoryRepository.Update(request)

	if err != nil {
		s.logger.Error("failed to update category: ", zap.Error(err))
		return nil, err
	}

	categoryEntity := s.mapper.ToCategoryResponse(update)

	return categoryEntity, nil
}

func (s *categoryService) Delete(id int) error {
	err := s.categoryRepository.Delete(id)

	if err != nil {
		s.logger.Error("failed to delete category: ", zap.Error(err))
		return err
	}

	return nil
}
