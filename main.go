package main

import (
	"fmt"

	"github.com/meshachdamilare/auto-shop/shops"
)

func main() {
	car1 := shops.NewCar("Mercedes", "GLK", "red", 50000, 2, 2020)
	car2 := shops.NewCar("Toyota", "Camry", "blue", 3000, 7, 3018)
	car3 := shops.NewCar("Lexus", "VLC", "blue", 3000, 6, 3018)

	// instantiate a new store
	store := shops.NewStore()

	//  adding item
	store.AddProduct(car1)
	store.AddProduct(car2)
	store.AddProduct(car3)

	// listing all available products in the store
	store.ListAvailableProducts()
	fmt.Println()

	// sell item
	store.SellProduct(car1, 1)
	store.SellProduct(car2, 4)
	store.SellProduct(car3, 5)

	// list all sold item
	store.ListSoldProduct()
	fmt.Println()

	// listing all available products in the store
	store.ListAvailableProducts()
	fmt.Println()

}
