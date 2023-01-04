package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPromedio(t *testing.T) {
	//Arrange
	expected := 3.00
	//Act
	prom := promedio(1, 2, 3, 4, 5)
	//Assert
	assert.Equal(t, expected, prom)
}
