package shops

import (
	"fmt"
)

// A Car is a type of product because it satisfied the Product interface
type Car struct {
	ID                string  `json:"id"`
	Brand             string  `json:"brand"`
	Model             string  `json:"model"`
	Quantity          int     `json:"quantity"`
	Price             float64 `json:"price"`
	Color             string  `json:"color"`
	YearOfManufacture int     `json:"yearOfManufacture"`
}

// Initialize a new Car
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

// Returns car name
func (c *Car) GetName() string {
	return c.Brand + " " + c.Model
}

// Returns car Id
func (c *Car) GetID() string {
	return c.ID
}

// Returns car quantity
func (c *Car) GetQuantity() int {
	return c.Quantity
}

// Modifies car quantity
func (c *Car) SetQuantity(q int) {
	c.Quantity = q
}

// Returns car price
func (c *Car) GetPrice() float64 {
	return c.Price
}

// Returns car brand
func (c *Car) GetBrand() string {
	return c.Brand
}

// Returns car model
func (c *Car) GetModel() string {
	return c.Model
}

// Returns car color
func (c *Car) GetColor() string {
	return c.Color
}

// Returns car year of manufacture
func (c *Car) GetYearOfManufacture() int {
	return c.YearOfManufacture
}

// Return a compact struct of car
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

// Displays the status of car: in stock or out of stock
func (c *Car) DisplayProductStatus() {
	if c.GetQuantity() > 0 {
		fmt.Println("Product In stock")
	}
	fmt.Println("Product out of stock")
}
