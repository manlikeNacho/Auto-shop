package products

import (
	"errors"
	"fmt"
)

type Car struct {
	name   string
	engine string
	model  string
}

type ProductInterface interface {
	Display() string
	Status() bool
	Sell(quantity int) error
	Price() int
	Quantity() int
}

type Product struct {
	Car
	Name     string
	quantity int
	price    int
}

func NewProduct(carName string, quantity int, price int, productName string, engineName string, modelName string) *Product {
	return &Product{
		Car: Car{
			name:   carName,
			engine: engineName,
			model:  modelName,
		},
		Name:     productName,
		quantity: quantity,
		price:    price,
	}
}

func (p *Product) Display() string {
	return p.Name
}

func (p *Product) Quantity() int {
	return p.quantity
}

func (p *Product) Status() bool {
	if p.quantity <= 0 {
		fmt.Println("Product Out of Stock")
		return false
	}
	fmt.Printf("Quantity of Prduct in Stock: %v", p.quantity)
	return true
}

func (p *Product) Sell(quantity int) error {
	if p.quantity < quantity {
		return errors.New("not enough products to sell")
	}

	p.quantity = p.quantity - quantity

	return nil
}

func (p *Product) Price() int {
	return p.price
}
