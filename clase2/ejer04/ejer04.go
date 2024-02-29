package main

import (
	"fmt"
)

func main() {
	var employees = map[string]int{"Benjamin": 20, "Nahuel": 26, "Brenda": 19, "Darío": 44, "Pedro": 30}
	/*
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Ingresar nombre de empleado: ")
		selector, _ := reader.ReadString('\n')
		fmt.Println("name: ", selector)
	*/
	ageSelector := 21
	fmt.Println(employees, "Mayor a ", ageSelector, " : ", greatherThanInMap(&employees, ageSelector))
	employees["Federico"] = 25
	delete(employees, "Pedro")
	nameSelector := "Benjamin"
	value, ok := employees[nameSelector]
	if ok {
		fmt.Println(nameSelector, ": ", value)
	}
	fmt.Println(employees, "Mayor a ", ageSelector, " : ", greatherThanInMap(&employees, ageSelector))
}

func greatherThanInMap(m *map[string]int, n int) int {
	counter := 0
	for _, e := range *m {
		if e >= n {
			counter++
		}
	}
	return counter
}
