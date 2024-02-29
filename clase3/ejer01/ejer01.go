package main

import "fmt"

func main() {
	employeeSalaries := []float64{40000, 80000, 120000, 180000, 200000, 300000, 400000}
	for _, salary := range employeeSalaries {
		fmt.Println(salary, TaxesFromSalary(salary))
	}
}

func TaxesFromSalary(salary float64) float64 {
	switch {
	case salary <= 50000:
		return 0
	case salary <= 150000:
		return salary * 0.17
	default:
		return salary * 0.27
	}
}
