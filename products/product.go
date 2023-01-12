package products

import (
	"errors"
	"fmt"
)

// Car class for car that holds a few car props.
type car struct {
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

type product struct {
	car
	Name     string
	quantity int
	price    int
}

func NewProduct(carName string, quantity int, price int, productName string, engineName string, modelName string) *product {
	return &product{
		car: car{
			name:   carName,
			engine: engineName,
			model:  modelName,
		},
		Name:     productName,
		quantity: quantity,
		price:    price,
	}
}

func (p *product) Display() string {
	return p.Name
}

func (p *product) Quantity() int {
	return p.quantity
}

func (p *product) Status() bool {
	if p.quantity <= 0 {
		fmt.Println("Product Out of Stock")
		return false
	}
	fmt.Printf("Quantity of Prduct in Stock: %v", p.quantity)
	return true
}

func (p *product) Sell(quantity int) error {
	if p.quantity < quantity {
		return errors.New("not enough products to sell")
	}

	p.quantity = p.quantity - quantity

	return nil
}

func (p *product) Price() int {
	return p.price
}
