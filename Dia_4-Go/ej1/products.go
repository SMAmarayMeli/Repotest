package main

import "fmt"

type Product struct {
	Id          int
	Name        string
	Price       float64
	Description string
	Category    string
}

var products = []Product{Product{Id: 1, Name: "Gol", Price: 500, Description: "Gol vehiculo", Category: "Volkswagen"},
	Product{Id: 2, Name: "Bora", Price: 500, Description: "Bora vehiculo", Category: "Volkswagen"},
	Product{Id: 3, Name: "Golf", Price: 500, Description: "Golf vehiculo", Category: "Volkswagen"}}

func (p Product) save() {
	products = append(products, p)
}

func (p Product) getAll() {
	for _, a := range products {
		fmt.Printf("ID: %v\nName: %v\nPrice: %v\nDescripction: %v\nCategory: %v\n", a.Id, a.Name, a.Price, a.Description, a.Category)
	}
}

func main() {
	new := Product{
		Id:          4,
		Name:        "Toro",
		Price:       7200000,
		Description: "Toro vehiculo",
		Category:    "Fiat",
	}

	new.save()

	new.getAll()
}
