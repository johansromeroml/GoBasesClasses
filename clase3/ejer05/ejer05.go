package main

import "fmt"

const (
	dog     = "dog"
	cat     = "cat"
	hamster = "hamster"
	spider  = "spider"
)

func main() {
	animalDog, msg := animal(dog)
	fmt.Print(msg)
	animalCat, msg := animal(cat)
	fmt.Print(msg)
	animalHamster, msg := animal(hamster)
	fmt.Print(msg)
	animalSpider, msg := animal(spider)
	fmt.Print(msg)
	var amount float64
	amount += animalDog(10)
	amount += animalCat(10)
	amount += animalHamster(20)
	amount += animalSpider(20)
	fmt.Println()
	fmt.Println("Total food amount: ", amount)
}

func animal(op string) (f func(int) float64, err string) {
	switch op {
	case dog:
		f = dogFood
	case cat:
		f = catFood
	case hamster:
		f = hamsterFood
	case spider:
		f = spiderFood
	default:
		err = op + "animal not defined"
	}
	return
}

func dogFood(n int) float64 {
	return float64(n) * 10.0
}

func catFood(n int) float64 {
	return float64(n) * 5.0
}

func hamsterFood(n int) float64 {
	return float64(n) * 0.25
}

func spiderFood(n int) float64 {
	return float64(n) * 0.15
}
