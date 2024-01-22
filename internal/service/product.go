package service

import (
	"github.com/renaldyhidayatt/inventorygoent/internal/domain/request"
	"github.com/renaldyhidayatt/inventorygoent/internal/domain/response"
	"github.com/renaldyhidayatt/inventorygoent/internal/mapper"
	"github.com/renaldyhidayatt/inventorygoent/internal/repository"
	"github.com/renaldyhidayatt/inventorygoent/pkg/logger"
	"go.uber.org/zap"
)

type productService struct {
	logger            logger.Logger
	productRepository repository.ProductRepository
	mapper            mapper.ProductMapping
}

func NewProductService(logger logger.Logger, productRepository repository.ProductRepository, mapper mapper.ProductMapping) *productService {
	return &productService{
		logger:            logger,
		productRepository: productRepository,
		mapper:            mapper,
	}
}

func (s *productService) GetAll() (*[]response.Product, error) {
	products, err := s.productRepository.GetAll()
	if err != nil {
		s.logger.Error("failed to get products: ", zap.Error(err))
		return nil, err
	}
	return s.mapper.ToProductResponses(products), nil
}

func (s *productService) GetById(id int) (*response.Product, error) {
	product, err := s.productRepository.GetById(id)
	if err != nil {
		s.logger.Error("failed to get product: ", zap.Error(err))
		return nil, err
	}
	return s.mapper.ToProductResponse(product), nil
}

func (s *productService) Create(request request.CreateProductRequest) (*response.Product, error) {
	create, err := s.productRepository.Create(request)
	if err != nil {
		s.logger.Error("failed to create product: ", zap.Error(err))
		return nil, err
	}
	productEntity := s.mapper.ToProductResponse(create)
	return productEntity, nil
}

func (s *productService) Update(request request.UpdateProductRequest) (*response.Product, error) {
	update, err := s.productRepository.Update(request)
	if err != nil {
		s.logger.Error("failed to update product: ", zap.Error(err))
		return nil, err
	}
	productEntity := s.mapper.ToProductResponse(update)
	return productEntity, nil
}

func (s *productService) Delete(id int) error {
	err := s.productRepository.Delete(id)
	if err != nil {
		s.logger.Error("failed to delete product: ", zap.Error(err))
		return err
	}
	return nil
}
