package main

import (
	"testing"
)

func TestTaxesFromSalary_under50000(t *testing.T) {
	//Arrange
	salary := 40000.0
	expectedResult := 0.0
	//Act
	result := TaxesFromSalary(salary)
	//Assert
	if result != expectedResult {
		t.Errorf("Obtainded %f, but expected %f", result, expectedResult)
	}
}
func TestTaxesFromSalary_above50000(t *testing.T) {
	//Arrange
	salary := 80000.0
	expectedResult := 13600.0
	//Act
	result := TaxesFromSalary(salary)
	//Assert
	if int(result*1000000) != int(expectedResult*1000000) {
		t.Errorf("Obtainded %f, but expected %f", result, expectedResult)
	}
}
func TestTaxesFromSalary_under150000(t *testing.T) {
	//Arrange
	salary := 200000.0
	expectedResult := 54000.0
	//Act
	result := TaxesFromSalary(salary)
	//Assert
	if result != expectedResult {
		t.Errorf("Obtainded %f, but expected %f", result, expectedResult)
	}
}
