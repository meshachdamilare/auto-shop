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
	for _, v := range s.Instock {
		if v.GetQuantity() > 0 {
			jsonResp, err := json.MarshalIndent(v.DisplayProduct(), "", "\t")
			if err != nil {
				log.Fatalln(err)
			}
			allProduct = append(allProduct, string(jsonResp))
		}
	}
	fmt.Println(allProduct)

}

func (s *Store) AddItem(c ProductItem) {
	id := c.GetID()
	if _, ok := s.Instock[id]; !ok {
		s.Instock[id] = Product{ProductItem: c}
	}

	//s.Instock = append(s.Instock, &Product{ProductItem: c})
}

func (s *Store) SellItem(c ProductItem, quantity int) {
	id := c.GetID()
	availQuantity := c.GetQuantity()
	if quantity > 0 && availQuantity > 0 {
		if item, ok := s.Instock[id]; ok {
			item2 := item
			if availQuantity >= quantity {
				nQ := availQuantity - quantity
				item.SetQuantity(nQ)
				s.Instock[id] = item
				item2.SetQuantity(quantity)
				s.OutStock[id] = item2
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

}
