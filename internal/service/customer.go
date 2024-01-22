package service

import (
	"github.com/renaldyhidayatt/inventorygoent/internal/domain/request"
	"github.com/renaldyhidayatt/inventorygoent/internal/domain/response"
	"github.com/renaldyhidayatt/inventorygoent/internal/mapper"
	"github.com/renaldyhidayatt/inventorygoent/internal/repository"
	"github.com/renaldyhidayatt/inventorygoent/pkg/logger"
	"go.uber.org/zap"
)

type customerService struct {
	logger             logger.Logger
	customerRepository repository.CustomerRepository
	mapper             mapper.CustomerMapping
}

func NewCustomerService(logger logger.Logger, customerRepository repository.CustomerRepository, mapper mapper.CustomerMapping) *customerService {
	return &customerService{
		logger:             logger,
		customerRepository: customerRepository,
		mapper:             mapper,
	}
}

func (s *customerService) GetAll() (*[]response.Customer, error) {
	customers, err := s.customerRepository.GetAll()
	if err != nil {
		s.logger.Error("failed to get customers: ", zap.Error(err))
		return nil, err
	}
	return s.mapper.ToCustomerResponses(customers), nil
}

func (s *customerService) GetById(id int) (*response.Customer, error) {
	customer, err := s.customerRepository.GetById(id)
	if err != nil {
		s.logger.Error("failed to get customer: ", zap.Error(err))
		return nil, err
	}
	return s.mapper.ToCustomerResponse(customer), nil
}

func (s *customerService) FindByName(name string) (*response.Customer, error) {
	customer, err := s.customerRepository.GetByName(name)
	if err != nil {
		s.logger.Error("failed to get customer: ", zap.Error(err))
		return nil, err
	}
	return s.mapper.ToCustomerResponse(customer), nil
}

func (s *customerService) Create(request request.CreateCustomerRequest) (*response.Customer, error) {
	customer, err := s.customerRepository.Create(request)
	if err != nil {
		s.logger.Error("failed to create customer: ", zap.Error(err))
		return nil, err
	}
	return s.mapper.ToCustomerResponse(customer), nil
}

func (s *customerService) Update(request request.UpdateCustomerRequest) (*response.Customer, error) {
	customer, err := s.customerRepository.Update(request)
	if err != nil {
		s.logger.Error("failed to update customer: ", zap.Error(err))
		return nil, err
	}
	return s.mapper.ToCustomerResponse(customer), nil
}

func (s *customerService) Delete(id int) error {
	err := s.customerRepository.Delete(id)
	if err != nil {
		s.logger.Error("failed to delete customer: ", zap.Error(err))
		return err
	}
	return nil
}
