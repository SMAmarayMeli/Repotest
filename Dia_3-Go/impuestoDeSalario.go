package main

import "fmt"

func main() {
	fmt.Println("Ingrese sueldo: ")
	var sueldo float64
	fmt.Scanln(&sueldo)

	fmt.Printf("Impuestos a aplicar: $%f \n", impuestoSalario(sueldo))
}

func impuestoSalario(sueldo float64) float64 {
	if sueldo > 50000 {
		return sueldo * 0.17
	}
	if sueldo > 150000 {
		return sueldo * 0.27
	}
	return 0
}
