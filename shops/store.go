package shops

import (
	"encoding/json"
	"fmt"
	"log"
)

type StoreItem interface {
	AddItem()
	ListAllProduct()
	SellItem()
	ListSoldItem()
}

type Store struct {
	ID       string
	InStock  map[string]Product
	OutStock map[string]Product
}

// NewStore ... Create a new store during use
func NewStore() *Store {
	return &Store{
		ID:       GenerateID(),
		InStock:  make(map[string]Product),
		OutStock: make(map[string]Product),
	}
}

func (s *Store) AvailableProduct() {
	fmt.Println("All Available Products for sales")
	var availableProduct []any
	for _, v := range s.InStock {
		if v.Product.GetQuantity() > 0 {
			jsonResp, err := json.MarshalIndent(v.Product.GetStruct(), "", "\t")
			if err != nil {
				log.Fatalln(err)
			}
			availableProduct = append(availableProduct, string(jsonResp))
		}
	}
	fmt.Println(availableProduct)
}

func (s *Store) ListAllProduct() {
	fmt.Println("All Products in the store")
	var allProduct []any
	for _, v := range s.InStock {
		jsonResp, err := json.MarshalIndent(v.Product.GetStruct(), "", "\t")
		if err != nil {
			log.Fatalln(err)
		}
		allProduct = append(allProduct, string(jsonResp))
	}
	fmt.Println(allProduct)
}

func (s *Store) ListSoldItem() {
	fmt.Println("List of all sold Item")
	var allSoldProduct []any
	total_price := 0.0
	for _, p := range s.OutStock {
		total_price += p.Product.GetPrice() * p.Product.GetQuantity()
		jsonResp, err := json.MarshalIndent(p.Product.GetStruct(), "", "\t")
		if err != nil {
			log.Fatalln(err)
		}
		allSoldProduct = append(allSoldProduct, string(jsonResp))
	}
	fmt.Println(allSoldProduct)
	fmt.Println("Total sale", total_price)
}

func (s *Store) AddItem(c *Car) {
	id := c.GetID()
	s.InStock[id] = Product{
		Product: c,
	}
}

func (s *Store) SellItem(c *Car, quantity float64) {
	id := c.GetID()
	m := make(map[string]Product)
	if item, ok := s.InStock[id]; ok {
		inStockQuantity := item.Product.GetQuantity()
		if inStockQuantity > 0 {
			if quantity > inStockQuantity {
				fmt.Println("Not enough quantity for sale")
				return
			} else if inStockQuantity == 1 {
				s.OutStock[id] = item
				m[id] = item
				fmt.Printf("Congrats you just sold %v %s for %v each\n", quantity, item.Product.GetName(), item.Product.GetPrice())
				delete(s.InStock, id)
			} else {
				item.Product.SetQuantityIn(quantity)
				s.InStock[id] = item
				fmt.Printf("Congrats you just sold %v %s for %v each\n", quantity, item.Product.GetName(), item.Product.GetPrice())

				if outItem, ok := s.OutStock[id]; ok {
					outItem.Product.SetQuantityOut(quantity)
					s.OutStock[id] = outItem

				}

			}

		} else {
			fmt.Println("Not enough quantity up for sale")
		}
	}
}
