package main

import (
	"errors"
	"fmt"
)

type VendingMachine struct {
	inventory        map[string]Product
	balance          uint
	paymentProcessor PaymentProcessor
	changeDispenser  ChangeDispenser
	cart             []Product
}

func NewVendingMachine(products []Product) *VendingMachine {
	inventory := make(map[string]Product)
	for _, product := range products {
		inventory[product.code] = product
	}

	return &VendingMachine{
		inventory:        inventory,
		balance:          0,
		paymentProcessor: nil,
		changeDispenser:  nil,
	}
}
func (v VendingMachine) DisplayProducts() {
	fmt.Println("Available Products:")
	for _, product := range v.inventory {
		fmt.Printf("Code: %s, Name: %s, Price: %d, Quantity: %d\n", product.code, product.name, product.price, product.quantity)
	}
}

func (v VendingMachine) SelectProduct(productCode string) (Product, error) {
	product, ok := v.inventory[productCode]
	if !ok || product.quantity <= 0 {
		return Product{}, fmt.Errorf("product %s not available", productCode)
	}
	return product, nil
}

func (v VendingMachine) AddProduct(productCode string, p Product) {
	_, ok := v.inventory[productCode]
	if !ok {
		v.inventory[productCode] = p
	} else {
		p.quantity = v.inventory[productCode].quantity + p.quantity
		v.inventory[productCode] = p
	}
}

func (v VendingMachine) AddToCart(code string) error {
	product, ok := v.inventory[code]
	if !ok || product.quantity <= 0 {
		return fmt.Errorf("product %s not available", code)
	}

	v.cart = append(v.cart)
	return nil
}

func (v VendingMachine) CartValue() uint {

	return 0
}

func (v VendingMachine) DispenseProduct() error {
	if len(v.cart) <= 0 {
		return errors.New("no products selected")
	}

	if v.balance < v.CartValue() {
		return errors.New("insufficient balance")
	}

	if v.paymentProcessor.isPaymentDone() {
		for _, p := range v.cart {
			fmt.Printf("Dispensing Product: %v\n", p)
		}
	}

	return nil
}

type PaymentProcessor interface {
	isPaymentDone() bool
}

type ChangeDispenser interface {
}

type Product struct {
	name     string
	code     string
	price    uint
	quantity uint
}

func NewProduct(code, name string, price, quantity uint) Product {
	return Product{
		code:     code,
		name:     name,
		price:    price,
		quantity: quantity,
	}
}

func main() {

}
