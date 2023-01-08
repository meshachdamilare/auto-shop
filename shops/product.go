package shops

// Product interface: any product that has all these methods implement this interface
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
