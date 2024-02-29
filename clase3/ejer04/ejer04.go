package main

import "fmt"

const (
	minimum = "minimum"
	average = "average"
	maximum = "maximum"
)

func main() {
	minFunc, err := operationSelector(minimum)
	fmt.Println(err)
	averageFunc, err := operationSelector(average)
	fmt.Println(err)
	maxFunc, err := operationSelector(maximum)
	fmt.Println(err)
	minValue := minFunc(2, 3, 3, 4, 10, 2, 4, 5)
	fmt.Println(minValue)
	averageValue := averageFunc(2, 3, 3, 4, 1, 2, 4, 5)
	fmt.Println(averageValue)
	maxValue := maxFunc(2, 3, 3, 4, 1, 2, 4, 5)
	fmt.Println(maxValue)
}

func operationSelector(op string) (f func(...int) float64, err string) {
	switch op {
	case minimum:
		f = min
	case maximum:
		f = max
	case average:
		f = avg
	default:
		err = "operation not defined"
	}
	return
}

func min(vals ...int) float64 {
	if len(vals) > 0 {
		min := vals[0]
		for _, v := range vals {
			if v < min {
				min = v
			}
		}
		return float64(min)
	}
	return -1
}

func max(vals ...int) float64 {
	if len(vals) > 0 {
		max := 0
		for _, v := range vals {
			if v > max {
				max = v
			}
		}
		return float64(max)
	}
	return -1
}

func avg(vals ...int) float64 {
	if len(vals) > 0 {
		sum := 0
		for _, v := range vals {
			sum += v
		}
		return float64(sum / len(vals))
	}
	return -1
}
