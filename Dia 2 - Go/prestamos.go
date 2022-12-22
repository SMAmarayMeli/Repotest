package main

import (
	"fmt"
)

func main() {

	fmt.Println("Ingrese edad")
	var edad int
	fmt.Scanln(&edad)
	fmt.Println("Tiene trabajo? S - N")
	var trabajo string
	fmt.Scanln(&trabajo)
	fmt.Println("Antiguedad en trabajo? Años")
	var ant int
	fmt.Scanln(&ant)
	fmt.Println("Sueldo?")
	var sueldo float32
	fmt.Scanln(&sueldo)

	switch {
	case edad < 22:
		fmt.Println("Eres muy joven para pedir prestamo")
	case trabajo == "N" || trabajo == "n":
		fmt.Println("Necesitas trabajo para prestamo")
	case ant < 1:
		fmt.Println("Necesitas tener mas de un año de antiguedad para prestamo")
	default:
		if sueldo > 100000 {
			fmt.Println("Prestamos aprobado sin interes")
		} else {
			fmt.Println("Prestamos aprobado con interes")
		}
	}

}
