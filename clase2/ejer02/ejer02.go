package main

import "fmt"

func main() {
	age := 30
	employment := true
	antiquity := 14
	if age > 22 {
		if employment {
			if antiquity >= 12 {
				fmt.Println("Able to get a mortgage")
			} else {
				fmt.Println("Not enough antiquity")
			}
		} else {
			fmt.Println("Not employed")
		}
	} else {
		fmt.Println("Not old enough")
	}

}
