package mapper

import (
	"github.com/renaldyhidayatt/inventorygoent/internal/domain/response"
	"github.com/renaldyhidayatt/inventorygoent/pkg/database/postgres/ent"
)

type CategoryMapping interface {
	ToCategoryResponse(category *ent.Category) *response.Category
	ToCategoryResponses(category []*ent.Category) *[]response.Category
}

type CustomerMapping interface {
	ToCustomerResponse(customer *ent.Customer) *response.Customer
	ToCustomerResponses(customers []*ent.Customer) *[]response.Customer
}

type ProductMapping interface {
	ToProductResponse(product *ent.Product) *response.Product
	ToProductResponses(products []*ent.Product) *[]response.Product
}

type ProductKeluarMapping interface {
	ToProductKeluarResponse(productKeluar *ent.ProductKeluar) *response.ProductKeluar
	ToProductKeluarResponses(productKeluar []*ent.ProductKeluar) *[]response.ProductKeluar
}

type ProductMasukMapping interface {
	ToProductMasukResponse(product *ent.ProductMasuk) *response.ProductMasuk
	ToProductMasukResponses(product []*ent.ProductMasuk) *[]response.ProductMasuk
}

type SupplierMapping interface {
	ToSupplierResponse(supplier *ent.Supplier) *response.Supplier
	ToSupplierResponses(suppliers []*ent.Supplier) *[]response.Supplier
}

type UserMapping interface {
	ToUserResponse(user *ent.User) *response.User
	ToUserResponses(users []*ent.User) *[]response.User
}
