package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSalarioA(t *testing.T) {
	//Arrange
	expected := 45000.00
	//Act
	resul := calcularSalario(600, "A")
	//Assert
	assert.Equal(t, expected, resul)
}

func TestSalarioB(t *testing.T) {
	//Arrange
	expected := 18000.00
	//Act
	resul := calcularSalario(600, "B")
	//Assert
	assert.Equal(t, expected, resul)
}

func TestSalarioC(t *testing.T) {
	//Arrange
	expected := 10000.00
	//Act
	resul := calcularSalario(600, "C")
	//Assert
	assert.Equal(t, expected, resul)
}
