package mapper

import (
	"github.com/renaldyhidayatt/inventorygoent/internal/domain/response"
	"github.com/renaldyhidayatt/inventorygoent/pkg/database/postgres/ent"
)

type customerMapper struct {
}

func NewCustomerMapper() *customerMapper {
	return &customerMapper{}
}

func (m *customerMapper) ToCustomerResponse(customer *ent.Customer) *response.Customer {
	return &response.Customer{
		Name:      customer.Name,
		Alamat:    customer.Alamat,
		Telepon:   customer.Telepon,
		CreatedAt: customer.CreatedAt,
		UpdatedAt: customer.UpdatedAt,
	}
}

func (m *customerMapper) ToCustomerResponses(customers []*ent.Customer) *[]response.Customer {
	var customerResponses []response.Customer

	for _, customer := range customers {
		customer := m.ToCustomerResponse(customer)

		customerResponses = append(customerResponses, *customer)
	}

	return &customerResponses
}
