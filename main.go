package main

import (
	"fmt"
	"github.com/Christomesh/auto-shop/shops"
)

func main() {
	car1 := shops.NewCar("Mercedes", "GLK", "red", 4, 50000, 2020)
	car2 := shops.NewCar("Toyota", "Camry", "blue", 7, 3000, 3018)
	car3 := shops.NewCar("Lexus", "VLC", "blue", 0, 3000, 3018)

	// instantiate a new store
	store := shops.NewStore()

	//  adding item
	store.AddItem(car1)
	store.AddItem(car2)
	store.AddItem(car3)

	// listing all products in the store
	store.ListAllProduct()
	fmt.Println()

	// sell item
	store.SellItem(car1, 1)
	store.SellItem(car2, 5)

	// list all sold item
	store.ListSoldItem()
	fmt.Println()
	store.ListAllProduct()

}
