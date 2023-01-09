package shops

// Product interface: any product that has all these methods implement this interface
type ProductItem interface {
	ProductGetter
	ProductSetter
	DisplayProduct() any
	DisplayProductStatus()
}

// Get the value of the Product Field
type ProductGetter interface {
	GetName() string
	GetID() string
	GetQuantity() int
	GetPrice() float64
	GetBrand() string
	GetModel() string
	GetColor() string
	GetYearOfManufacture() int
}

// To set value for the Product fields
type ProductSetter interface {
	SetQuantity(int)
}
