package main

import (
	"fmt"
	"testing"
)

func TestAnimal(t *testing.T) {
	//Arrange
	dogAmount := 10
	catAmount := 20
	hamsterAmount := 25
	spiderAmount := 15

	expectedResult := 208.5
	//Act
	animalDog, msg := Animal(dog)
	fmt.Print(msg)
	animalCat, msg := Animal(cat)
	fmt.Print(msg)
	animalHamster, msg := Animal(hamster)
	fmt.Print(msg)
	animalSpider, msg := Animal(spider)
	fmt.Print(msg)

	var amount float64
	amount += animalDog(dogAmount)
	amount += animalCat(catAmount)
	amount += animalHamster(hamsterAmount)
	amount += animalSpider(spiderAmount)

	//Assert
	if int(amount*1000000) != int(expectedResult*1000000) {
		t.Errorf("Obtainded %f, but expected %f", amount, expectedResult)
	}
}
