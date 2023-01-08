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
	Instock  map[string]ProductItem
	OutStock map[string]ProductItem
}

func NewStore() *Store {
	return &Store{
		Instock:  make(map[string]ProductItem),
		OutStock: make(map[string]ProductItem),
	}
}

func (s *Store) AvailableProduct() {
	fmt.Println("All Available Products for sales")
	var allProduct []string
	for _, item := range s.Instock {
		if item.GetQuantity() > 0 {
			jsonResp, err := json.MarshalIndent(item.DisplayProduct(), "", "\t")
			if err != nil {
				log.Fatalln(err)
			}
			allProduct = append(allProduct, string(jsonResp))
		}
	}
	fmt.Println(allProduct)

}

func (s *Store) ListSoldItem() {
	fmt.Println("List of all sold Item")
	var allSoldProduct []any
	total_price := 0.0
	for _, item := range s.OutStock {
		total_price += item.GetPrice() * float64(item.GetQuantity())
		jsonResp, err := json.MarshalIndent(item.DisplayProduct(), "", "\t")
		if err != nil {
			log.Fatalln(err)
		}
		allSoldProduct = append(allSoldProduct, string(jsonResp))
	}
	fmt.Println(allSoldProduct)
	fmt.Println("Total sale", total_price)
}

func (s *Store) AddItem(c ProductItem) {
	id := c.GetID()
	if _, ok := s.Instock[id]; !ok {
		s.Instock[id] = Product{ProductItem: c}
	}
}

func (s *Store) SellItem(c ProductItem, quantity int) {
	id := c.GetID()

	availQuantity := c.GetQuantity()
	if quantity > 0 && availQuantity > 0 {
		if item, ok := s.Instock[id]; ok {
			if availQuantity >= quantity {
				nQ := availQuantity - quantity
				if nQ == 0 {
					item.SetQuantity(0)
					s.Instock[id] = item

				} else {
					item.SetQuantity(nQ)
					s.Instock[id] = item
				}
			} else {
				fmt.Println("Available quantity not up the quantity demanded: ", quantity)
				return
			}
		} else {
			fmt.Println("Product Id not found !!!")
			return
		}
	} else {
		fmt.Println("Not enought quantity for sale")
		return
	}
	for k, v := range s.Instock {
		s.OutStock[k] = v
	}
	s.OutStock[id].SetQuantity(quantity)

}
