package main

import "fmt"

func main() {
	grades := []float64{1, 2, 3, 4, 5, 3.5, 4.2, 2.8, 2.5, 5, 5}
	avg, err := AverageFromSlice(grades)
	fmt.Println("Grades: ", grades, err, "Average: ", avg)
}

func AverageFromSlice(grades []float64) (average float64, err string) {
	size := len(grades)
	negatives := false
	for _, v := range grades {
		if v < 0 {
			err = "There are negative grades"
			negatives = true
			break
		}
		average += v
	}
	if !negatives {
		average /= float64(size)
	} else {
		average = 0
	}
	return
}
