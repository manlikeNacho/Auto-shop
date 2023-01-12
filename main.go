package main

import (
	"Auto-shop/products"
	"Auto-shop/store"
)

func main() {
	product1 := products.NewProduct("Benz", 100, 20000.00, "Car", "V8 engine", "Cl3")
	product2 := products.NewProduct("toyota", 100, 5000.00, "motor", "V6 engine", "corolla sports")

	var NachoStore store.Store
	NachoStore.AddProduct(product1)
	NachoStore.AddProduct(product2)
	NachoStore.ListItems()

}
