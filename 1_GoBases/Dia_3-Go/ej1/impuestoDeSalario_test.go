package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestImpuestoSalario(t *testing.T) {
	//Arrange
	var sueldo float64 = 10000.00
	var expected float64 = 0.00
	//Act
	result := impuestoSalario(sueldo)
	//Assert
	assert.Equal(t, expected, result)
}

func TestImpuestoSalarioMayor50000(t *testing.T) {
	//Arrange
	var sueldo float64 = 50001.00
	var expected float64 = 8500.17
	//Act
	result := impuestoSalario(sueldo)
	//Assert
	assert.Equal(t, expected, result)
}

func TestImpuestoSalarioMayor150000(t *testing.T) {
	//Arrange
	var sueldo float64 = 151000.00
	var expected float64 = 40770
	//Act
	result := impuestoSalario(sueldo)
	//Assert
	assert.Equal(t, expected, result)
}
