package product

type ProductCategory struct {
	Category string
}

type Address struct {
	Street string
	Number int64
}

type Product struct {
	ProductCategory *ProductCategory
	Address         *Address
	Name            string
}

func NewProduct() *Product {
	productCategory := &ProductCategory{Category: "clothes"}
	address := &Address{
		Street: "Paulista Avanue",
		Number: 208,
	}
	name := "Red shirt"

	return &Product{
		productCategory,
		address,
		name,
	}
}
