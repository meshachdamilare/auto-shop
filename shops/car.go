package shops

import (
	"fmt"
)

type Car struct {
	ID                string  `json:"id"`
	Brand             string  `json:"brand"`
	Model             string  `json:"model"`
	Quantity          float64 `json:"quantity"`
	Price             float64 `json:"price"`
	Color             string  `json:"color"`
	YearOfManufacture int64   `json:"yearOfManufacture"`
}

type CarOut struct {
	ID                string  `json:"id"`
	Brand             string  `json:"brand"`
	Model             string  `json:"model"`
	Quantity          float64 `json:"quantity"`
	Price             float64 `json:"price"`
	Color             string  `json:"color"`
	YearOfManufacture int64   `json:"yearOfManufacture"`
}

// NewCar ... creates a new car type during use
func NewCar(brand, model, color string, quantity, price float64, year int64) *Car {
	return &Car{
		ID:                GenerateID(),
		Brand:             brand,
		Model:             model,
		Quantity:          quantity,
		Price:             price,
		Color:             color,
		YearOfManufacture: year,
	}
}
func (c *Car) GetName() string {
	return c.Brand + " " + c.Model
}

func (c *Car) GetID() string {
	return c.ID
}

func (c *Car) GetQuantity() float64 {
	return c.Quantity
}

func (c *Car) GetPrice() float64 {
	return c.Price
}

func (c *Car) GetBrand() string {
	return c.Brand
}

func (c *Car) GetModel() string {
	return c.Model
}

func (c *Car) GetColor() string {
	return c.Color
}

func (c *Car) SetQuantityIn(q float64) {
	c.Quantity -= q
}

func (c Car) SetQuantityOut(q float64) {
	c.Quantity += q
}

func (c *Car) GetYearOfManufacture() int64 {
	return c.YearOfManufacture
}

func (c *Car) GetStruct() interface{} {
	v := &Car{
		ID:                c.GetID(),
		Quantity:          c.GetQuantity(),
		Brand:             c.GetBrand(),
		Model:             c.GetModel(),
		Price:             c.GetPrice(),
		Color:             c.GetColor(),
		YearOfManufacture: c.GetYearOfManufacture(),
	}
	return v
}

func (c *Car) DisplayProductStatus() (string, bool) {
	msg := ""
	if c.Quantity > 0 {
		msg = fmt.Sprintf("Product In stock")
		return msg, true
	}
	msg = fmt.Sprintf("Product out of stock")
	return msg, false
}
