package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_SalaryCalculationCategoryHours(t *testing.T) {
	t.Run("success - salary from category A", func(t *testing.T) {
		// Arrange
		category := "A"
		hours := 10
		expectedResult := 10000
		// Act
		result := SalaryCalculationCategoryHours(category, hours)
		// Assert
		require.Equal(t, expectedResult, result)
	})
	t.Run("success - salary from category B", func(t *testing.T) {
		// Arrange
		category := "B"
		hours := 10
		expectedResult := 18000
		// Act
		result := SalaryCalculationCategoryHours(category, hours)
		// Assert
		require.Equal(t, expectedResult, result)
	})
	t.Run("success - salary from category C", func(t *testing.T) {
		// Arrange
		category := "C"
		hours := 10
		expectedResult := 45000
		// Act
		result := SalaryCalculationCategoryHours(category, hours)
		// Assert
		require.Equal(t, expectedResult, result)
	})
	t.Run("success - salary from category D", func(t *testing.T) {
		// Arrange
		category := "D"
		hours := 10
		expectedResult := 0
		// Act
		result := SalaryCalculationCategoryHours(category, hours)
		// Assert
		require.Equal(t, expectedResult, result)
	})
}
