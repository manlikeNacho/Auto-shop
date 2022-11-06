package store

import (
	"Auto-shop/products"
	"fmt"
)

type Store struct {
	name         string
	productStore []products.Product
	soldProducts []products.Product
}

func (s *Store) AddProduct(p products.Product) {
	s.productStore = append(s.productStore, p)
}

func (s *Store) ListItemsCount() int {
	fmt.Printf("list items count: %+v \n\n", s.productStore)
	var count int
	for _, v := range s.productStore {
		count += v.Quantity()
	}
	return count
}

func (s *Store) ListItems() {
	if len(s.productStore) > 0 {
		for _, v := range s.productStore {
			fmt.Printf("list items: %+v \n\n", v.Display())
		}
	} else {
		fmt.Println("No products in the store atm.")
	}
}

func (s *Store) ItemsTotalPrice() int {
	fmt.Printf("items total price:%+v \n\n", s.productStore)
	var price int
	for _, v := range s.productStore {
		price += v.Quantity() * v.Price()
	}
	return price
}

func (s *Store) SellItems(productName string, quantity int) {
	var product products.Product

	for _, v := range s.productStore {
		if v.Name == productName {
			product = v
			if quantity <= v.Quantity() {
				product.Sell(quantity)
				s.soldProducts = append(s.soldProducts, product)
				fmt.Printf("You just sold %d amount of %s. /n The remaining %s in the store is %d", quantity, productName, productName, v.Quantity())
			} else {
				fmt.Printf("Insufficient Product in stock")
			}

		} else {
			fmt.Printf("%s is not in store", productName)
		}
	}
}

func (s *Store) ShowSoldItemsCount() {
	var sumNumberSold int
	for _, v := range s.soldProducts {
		sumNumberSold += v.Quantity()
	}
	fmt.Printf("total number of items sold: %v /n", sumNumberSold)
}

func (s *Store) ShowSoldItems() {
	fmt.Printf("sold items: %+v \n\n", s.soldProducts)
}

func (s *Store) ShowSoldItemsPrice() {
	var sumAmountSold int
	for _, v := range s.soldProducts {
		sumAmountSold += v.Price() * v.Quantity()
	}
	fmt.Printf("total sold items price: %v \n\n", sumAmountSold)
}
