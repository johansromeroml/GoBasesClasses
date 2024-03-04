package main

import (
	"errors"
	"fmt"
)

func main() {
	var salary int = 100000
	taxes, err := TaxesFromSalary(float64(salary))
	if err == nil {
		fmt.Println(taxes)
	} else {
		fmt.Println(err)
	}
	hours := 100
	value := 2000
	salary2, err := SalaryFromHoursAndValue(hours, value)
	if err == nil {
		fmt.Println(salary2)
	} else {
		fmt.Println(err)
	}
}

/*
type NotTaxableSalaryError struct {
	//error "the salary entered does not reach the taxable minimum"
}

func (e NotTaxableSalaryError) Error() string {
	return "Error: the salary entered does not reach the taxable minimum"
}
*/

/*
type SalaryBelowTenThousandError struct {
	//error "the salary entered does not reach the taxable minimum"
}

func (e SalaryBelowTenThousandError) Error() string {
	return "Error: Salary is less than 10000"
}
*/

func TaxesFromSalary(salary float64) (float64, error) {
	switch {
	case salary <= 10000:
		return 0.0, errors.New("Error: Salary is less than 10000")
	case salary <= 150000:
		return 0.0, fmt.Errorf("Error: the minimum taxable amount is 150,000 and the salary entered is: %0.0f", salary)
	default:
		return salary * 0.1, nil
	}
}

func SalaryFromHoursAndValue(hours, value int) (float64, error) {
	if hours < 80 {
		return 0.0, errors.New("Error: the worker cannot have worked less than 80 hours per month")
	}
	salary := float64(hours * value)
	if salary < 150000 {
		return salary, nil
	}
	return salary * 0.9, nil
}
