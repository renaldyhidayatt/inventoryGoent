package service

import (
	"github.com/renaldyhidayatt/inventorygoent/internal/domain/request"
	"github.com/renaldyhidayatt/inventorygoent/internal/domain/response"
	"github.com/renaldyhidayatt/inventorygoent/internal/mapper"
	"github.com/renaldyhidayatt/inventorygoent/internal/repository"
	"github.com/renaldyhidayatt/inventorygoent/pkg/logger"
	"go.uber.org/zap"
)

type productMasukService struct {
	logger                 logger.Logger
	productMasukRepository repository.ProductMasukRepository
	mapper                 mapper.ProductMasukMapping
}

func NewProductMasukService(logger logger.Logger, productMasukRepository repository.ProductMasukRepository, mapper mapper.ProductMasukMapping) *productMasukService {
	return &productMasukService{
		logger:                 logger,
		productMasukRepository: productMasukRepository,
		mapper:                 mapper,
	}
}

func (s *productMasukService) GetAll() (*[]response.ProductMasuk, error) {
	productMasuks, err := s.productMasukRepository.GetAll()
	if err != nil {
		s.logger.Error("failed to get product masuks: ", zap.Error(err))
		return nil, err
	}
	return s.mapper.ToProductMasukResponses(productMasuks), nil
}

func (s *productMasukService) GetById(id int) (*response.ProductMasuk, error) {
	productMasuk, err := s.productMasukRepository.GetById(id)
	if err != nil {
		s.logger.Error("failed to get product masuk: ", zap.Error(err))
		return nil, err
	}
	return s.mapper.ToProductMasukResponse(productMasuk), nil
}

func (s *productMasukService) Create(requests request.CreateProductMasukRequest) (*response.ProductMasuk, error) {

	productMasuk, err := s.productMasukRepository.Create(requests)
	if err != nil {
		s.logger.Error("failed to create product masuk: ", zap.Error(err))
		return nil, err
	}
	return s.mapper.ToProductMasukResponse(productMasuk), nil
}

func (s *productMasukService) Update(requests request.UpdateProductMasukRequest) (*response.ProductMasuk, error) {
	productMasuk, err := s.productMasukRepository.Update(requests)
	if err != nil {
		s.logger.Error("failed to update product masuk: ", zap.Error(err))
		return nil, err
	}
	return s.mapper.ToProductMasukResponse(productMasuk), nil
}

func (s *productMasukService) Delete(id int) error {
	err := s.productMasukRepository.Delete(id)
	if err != nil {
		s.logger.Error("failed to delete product masuk: ", zap.Error(err))
		return err
	}
	return nil
}
