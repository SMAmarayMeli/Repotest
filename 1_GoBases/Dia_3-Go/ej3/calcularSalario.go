package main

import "fmt"

func main() {
	var min float64
	var cat string
	fmt.Println("Ingrese minitos trabajos")
	fmt.Scanln(&min)
	fmt.Println("Ingrese categoria")
	fmt.Scanln(&cat)

	salario := calcularSalario(min, cat)

	fmt.Println("Salario es: $", salario)
}

func calcularSalario(min float64, cat string) (salario float64) {
	horas := min / 60
	switch cat {
	case "A":
		salario = 3000 * horas * 1.5
	case "B":
		salario = 1500 * horas * 1.2
	case "C":
		salario = 1000 * horas
	}
	return
}
