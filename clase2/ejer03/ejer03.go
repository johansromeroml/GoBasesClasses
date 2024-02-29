package main

import "fmt"

func main() {
	monthNumber := 20
	fmt.Println(monthFromNumber(monthNumber))
}

func monthFromNumber(num int) string {
	switch {
	case num == 1:
		return "January"
	case num == 2:
		return "February"
	case num == 3:
		return "March"
	case num == 4:
		return "April"
	case num == 5:
		return "May"
	case num == 6:
		return "June"
	case num == 7:
		return "July"
	case num == 8:
		return "August"
	case num == 9:
		return "September"
	case num == 10:
		return "October"
	case num == 11:
		return "November"
	case num == 12:
		return "December"
	default:
		return "not a valid month number"
	}
}
