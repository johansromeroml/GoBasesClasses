package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAverageFromSlice(t *testing.T) {
	// Arrange
	grades := []float64{1, 2, 3, 4, 5, 3.5, 4.2, 2.8, 2.5, 5, 5}
	expectedResult := (1 + 2 + 3 + 4 + 5 + 3.5 + 4.2 + 2.8 + 2.5 + 5 + 5) / 11
	//Act
	result, _ := AverageFromSlice(grades)
	//Assert
	assert.Equal(t, result, expectedResult)
}
