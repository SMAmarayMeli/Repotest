package main

import (
	"github.com/stretchr/testify/assert"
	"reflect"
	"runtime"
	"testing"
)

func TestEstadiscticaSinOperacion(t *testing.T) {
	//Arrange
	expectedErr := "Error no existe operacion"
	//Act
	resultFunc, resultErr := tipoCalculo("suma")
	funcname := runtime.FuncForPC(reflect.ValueOf(resultFunc).Pointer()).Name()
	//Assert
	assert.Equal(t, expectedErr, resultErr)
	assert.Equal(t, "", funcname)
}

func TestMinimum(t *testing.T) {
	//Arrange
	notas := []float64{10, 8, 9, 4}
	expected := 4.00
	//Act
	result := minimum(notas)
	//Assert
	assert.Equal(t, expected, result)
}

func TestMaximum(t *testing.T) {
	//Arrange
	notas := []float64{10, 8, 9, 4}
	expected := 10.00
	//Act
	result := maximum(notas)
	//Assert
	assert.Equal(t, expected, result)
}
