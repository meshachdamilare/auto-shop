package shops

type ProductItem interface {
	GetName() string
	GetID() string
	GetQuantity() int
	GetPrice() float64
	SetQuantity(int)
	GetBrand() string
	GetModel() string
	GetColor() string
	GetYearOfManufacture() int
	DisplayProduct() any
	DisplayProductStatus()
}

type Product struct {
	ProductItem
}
