package main

import "fmt"

var Products = []Product{
	{ID: 0, Name: "Escoba", Price: 2.5, Description: "Pa barrer", Category: "Casa"},
	{ID: 1, Name: "Balon futbol", Price: 5.5, Description: "Pa jugar futbol", Category: "Deportes"}}

type Product struct {
	ID          int
	Name        string
	Price       float64
	Description string
	Category    string
}

func Save(p Product) {
	Products = append(Products, p)
}

func GetAll() {
	fmt.Println(Products)
}

func GetById(id int) (p Product) {
	for _, v := range Products {
		if v.ID == id {
			p = v
		}
	}
	return
}

func main() {
	GetAll()
	Save(Product{ID: 3, Name: "Llanta moto", Price: 20.0, Description: "Pa la moto", Category: "motocicletas"})
	GetAll()
	fmt.Println(GetById(1))
}
