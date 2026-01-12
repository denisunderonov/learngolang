package main

import "fmt"

type Product struct {
	Name  string
	Price int
}

func (p *Product) Discount(persent int) {
	discountAmount := (p.Price * persent) / 100
	p.Price -= discountAmount
}

func main() {
	var prod Product
	prod.Name = "Laptop"
	prod.Price = 1000

	fmt.Println("Before discount:", prod.Price)
	prod.Discount(10)
	fmt.Println("After discount:", prod.Price)
}