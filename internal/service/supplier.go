package service

import (
	"github.com/renaldyhidayatt/inventorygoent/internal/domain/request"
	"github.com/renaldyhidayatt/inventorygoent/internal/domain/response"
	"github.com/renaldyhidayatt/inventorygoent/internal/mapper"
	"github.com/renaldyhidayatt/inventorygoent/internal/repository"
	"github.com/renaldyhidayatt/inventorygoent/pkg/logger"
	"go.uber.org/zap"
)

type supplierService struct {
	logger             logger.Logger
	supplierRepository repository.SupplierRepository
	mapper             mapper.SupplierMapping
}

func NewSupplierService(logger logger.Logger, supplierRepository repository.SupplierRepository, mapper mapper.SupplierMapping) *supplierService {
	return &supplierService{
		logger:             logger,
		supplierRepository: supplierRepository,
		mapper:             mapper,
	}
}

func (s *supplierService) GetAll() (*[]response.Supplier, error) {
	suppliers, err := s.supplierRepository.GetAll()
	if err != nil {
		s.logger.Error("failed to get suppliers: ", zap.Error(err))
		return nil, err
	}
	return s.mapper.ToSupplierResponses(suppliers), nil
}

func (s *supplierService) GetByID(id int) (*response.Supplier, error) {
	supplier, err := s.supplierRepository.GetById(id)
	if err != nil {
		s.logger.Error("failed to get supplier: ", zap.Error(err))
		return nil, err
	}
	return s.mapper.ToSupplierResponse(supplier), nil
}

func (s *supplierService) Create(requests request.CreateSupplierRequest) (*response.Supplier, error) {
	supplier, err := s.supplierRepository.Create(requests)
	if err != nil {
		s.logger.Error("failed to create supplier: ", zap.Error(err))
		return nil, err
	}
	return s.mapper.ToSupplierResponse(supplier), nil
}

func (s *supplierService) Update(requests request.UpdateSupplierRequest) (*response.Supplier, error) {
	supplier, err := s.supplierRepository.Update(requests)
	if err != nil {
		s.logger.Error("failed to update supplier: ", zap.Error(err))
		return nil, err
	}
	return s.mapper.ToSupplierResponse(supplier), nil
}

func (s *supplierService) Delete(id int) error {
	err := s.supplierRepository.Delete(id)
	if err != nil {
		s.logger.Error("failed to delete supplier: ", zap.Error(err))
		return err
	}
	return nil
}
