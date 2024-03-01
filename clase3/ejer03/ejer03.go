package main

func SalaryCalculationCategoryHours(cat string, hours int) int {
	switch cat {
	case "A":
		return hours * 1000
	case "B":
		return int(float64(hours*1500) * 1.2)
	case "C":
		return int(float64(hours*3000) * 1.5)
	default:
		return 0
	}
}
