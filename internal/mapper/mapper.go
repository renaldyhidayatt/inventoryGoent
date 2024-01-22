package mapper

type Mapper struct {
	CategoryMapper      CategoryMapping
	CustomerMapper      CustomerMapping
	ProductMapper       ProductMapping
	ProductKeluarMapper ProductKeluarMapping
	ProductMasukMapper  ProductMasukMapping
	SupplierMapper      SupplierMapping
	UserMapper          UserMapping
}

func NewMapper() *Mapper {
	return &Mapper{
		CategoryMapper:      NewCategoryMapper(),
		CustomerMapper:      NewCustomerMapper(),
		ProductMapper:       NewProductMapper(),
		ProductKeluarMapper: NewProductKeluarMapper(),
		ProductMasukMapper:  NewProductMasukMapper(),
		SupplierMapper:      NewSupplierMapper(),
		UserMapper:          NewUserMapper(),
	}
}
