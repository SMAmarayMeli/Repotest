package main

import "fmt"

func main() {
	notas := make([]int, 0)
	nota := 1
	for nota != 0 {
		fmt.Println("Ingresar notas: ")
		fmt.Scanln(&nota)

		for nota < 0 || nota > 10 {
			fmt.Println("Ingrese nota correctamente")
			fmt.Println("Ingresar notas: ")
			fmt.Scanln(&nota)
		}
		notas = append(notas, nota)
	}

	fmt.Println("Promedio: ", promedio(notas...))
}

func promedio(notas ...int) float64 {
	var num int
	for _, v := range notas {
		num += v
	}
	return float64(num) / float64(len(notas))
}
