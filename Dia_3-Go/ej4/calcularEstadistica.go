package main

import (
	"fmt"
	"os"
)

const (
	min = "minimum"
	ave = "average"
	max = "maximum"
)

func main() {
	fmt.Println("Indique tipo de calculo (minimum, average o maximum): ")
	var oper string
	fmt.Scanln(&oper)

	operacion, err := tipoCalculo(oper)
	if err != "" {
		println(err)
		os.Exit(0)
	}

	notas := make([]float64, 0)
	var num float64

	fmt.Println("Indique notas: 0 para salir")
	fmt.Scanln(&num)

	for num != 0 {
		notas = append(notas, num)
		fmt.Scanln(&num)
	}

	fmt.Println(notas)
	fmt.Println("Resultado de ", oper, ": ", operacion(notas))
}

func tipoCalculo(calculo string) (result func([]float64) float64, err string) {
	switch {
	case calculo == min:
		result = minimum
	case calculo == ave:
		result = average
	case calculo == max:
		result = maximum
	default:
		result = nil
		err = "Error no existe operacion"
	}
	return
}

func minimum(nums []float64) float64 {
	small := nums[0]
	for _, num := range nums {
		if num < small {
			small = num
		}
	}
	return small
}

func maximum(nums []float64) float64 {
	big := nums[0]
	for _, num := range nums {
		if num > big {
			big = num
		}
	}
	return big
}

func average(nums []float64) float64 {
	var sum float64
	println(nums)
	for _, num := range nums {
		println(sum)
		sum += num
	}
	return sum / float64(len(nums))
}
