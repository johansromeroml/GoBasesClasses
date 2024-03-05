package main

import (
	"errors"
	"fmt"
)

type Customer struct {
	File  string
	Name  string
	ID    uint
	Phone string
	Home  string
}

type CustomerList struct {
	Customers []Customer
}

var ourCustomers = []Customer{
	{
		File:  "John.txt",
		Name:  "John",
		ID:    1,
		Phone: "345644",
		Home:  "House",
	},
	{
		File:  "Jane.txt",
		Name:  "Jane",
		ID:    2,
		Phone: "323644",
		Home:  "Flat",
	},
	{
		File:  "Joe.txt",
		Name:  "Joe",
		ID:    3,
		Phone: "347744",
		Home:  "Farm",
	},
}

var errorsInExec uint = 0

func main() {

	defer func() {
		fmt.Println("End of execution")
		if errorsInExec > 0 {
			fmt.Println("Several errors were detected at runtime")
		}
	}()

	fmt.Println(ourCustomers)
	jane := Customer{
		File:  "Jane.txt",
		Name:  "Jane",
		ID:    2,
		Phone: "323644",
		Home:  "Flat",
	}
	juan := Customer{
		File:  "Juan.txt",
		Name:  "Juan",
		ID:    4,
		Phone: "111644",
		Home:  "TinyHome",
	}
	fmt.Println("Customer in list: ", checkCustomerInList(jane))
	zeroCheck, _ := jane.checkCustomerInZeros()
	fmt.Println("Customer in Zeros: ", zeroCheck)
	defer saveCustomer(juan)
	saveCustomer(jane)
}

func checkCustomerInList(c Customer) bool {
	flag := false
	for _, cust := range ourCustomers {
		if c == cust {
			flag = true
		}
	}
	return flag
}

func saveCustomer(c Customer) {
	defer func() {
		r := recover()
		if r != nil {
			fmt.Printf("caugth a panic. Error : %v \n", r)
		}
	}()

	if checkCustomerInList(c) {
		errorsInExec++
		panic("Error: Customer already exists")
	}
	ourCustomers = append(ourCustomers, c)
}

func (c Customer) checkCustomerInZeros() (flag bool, err error) {
	flag = false
	if c.File == "" || c.Name == "" || c.ID == 0 || c.Phone == "" || c.Home == "" {
		flag = true
		err = errors.New("A value in Customer is in zeros (0,\"\",nil)")
	}
	return
}
