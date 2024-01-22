package service

import (
	"github.com/renaldyhidayatt/inventorygoent/internal/domain/request"
	"github.com/renaldyhidayatt/inventorygoent/internal/domain/response"
	"github.com/renaldyhidayatt/inventorygoent/internal/mapper"
	"github.com/renaldyhidayatt/inventorygoent/internal/repository"
	"github.com/renaldyhidayatt/inventorygoent/pkg/logger"
	"go.uber.org/zap"
)

type productKeluar struct {
	logger                  logger.Logger
	productKeluarRepository repository.ProductKeluarRepository
	mapper                  mapper.ProductKeluarMapping
}

func NewProductKeluarService(logger logger.Logger, productKeluarRepository repository.ProductKeluarRepository, mapper mapper.ProductKeluarMapping) *productKeluar {
	return &productKeluar{
		logger:                  logger,
		productKeluarRepository: productKeluarRepository,
		mapper:                  mapper,
	}
}

func (s *productKeluar) GetAll() (*[]response.ProductKeluar, error) {
	productKeluars, err := s.productKeluarRepository.GetAll()
	if err != nil {
		s.logger.Error("failed to get product keluars: ", zap.Error(err))
		return nil, err
	}
	return s.mapper.ToProductKeluarResponses(productKeluars), nil
}

func (s *productKeluar) GetById(id int) (*response.ProductKeluar, error) {
	productKeluar, err := s.productKeluarRepository.GetById(id)
	if err != nil {
		s.logger.Error("failed to get product keluar: ", zap.Error(err))
		return nil, err
	}
	return s.mapper.ToProductKeluarResponse(productKeluar), nil
}

func (s *productKeluar) Create(requests request.CreateProductKeluarRequest) (*response.ProductKeluar, error) {
	res, err := s.productKeluarRepository.Create(requests)

	if err != nil {
		s.logger.Error("failed to create product keluar: ", zap.Error(err))
		return nil, err
	}

	return s.mapper.ToProductKeluarResponse(res), nil
}

func (s *productKeluar) Update(requests request.UpdateProductKeluarRequest) (*response.ProductKeluar, error) {
	res, err := s.productKeluarRepository.Update(requests)

	if err != nil {
		s.logger.Error("failed to update product keluar: ", zap.Error(err))
		return nil, err
	}

	return s.mapper.ToProductKeluarResponse(res), nil
}

func (s *productKeluar) Delete(id int) error {
	err := s.productKeluarRepository.Delete(id)
	if err != nil {
		s.logger.Error("failed to delete product keluar: ", zap.Error(err))
		return err
	}
	return nil
}
