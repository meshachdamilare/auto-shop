package shops

import (
	"fmt"
)

type Car struct {
	ID                string  `json:"id"`
	Brand             string  `json:"brand"`
	Model             string  `json:"model"`
	Quantity          int     `json:"quantity"`
	Price             float64 `json:"price"`
	Color             string  `json:"color"`
	YearOfManufacture int     `json:"yearOfManufacture"`
}

// NewCar ... creates a new car type during use
func NewCar(brand, model, color string, price float64, quantity, year int) *Car {
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

func (c *Car) GetQuantity() int {
	return c.Quantity
}

func (c *Car) SetQuantity(q int) {
	c.Quantity = q
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

func (c *Car) GetYearOfManufacture() int {
	return c.YearOfManufacture
}

func (c *Car) DisplayProduct() any {
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

func (c *Car) DisplayProductStatus() {
	if c.GetQuantity() > 0 {
		fmt.Println("Product In stock")
	}
	fmt.Println("Product out of stock")
}
