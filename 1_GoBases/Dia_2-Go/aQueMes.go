package main

import "fmt"

func main() {
	var mes int
	fmt.Scanln(&mes)

	// Uso un array para que no quede el codigo muy largo. Sino se podria usar un switch
	// para que quede sencillo de leer.

	array := [12]string{"Enero", "Febrero", "Marzo", "Abril", "Mayo", "Junio", "Julio",
		"Agosto", "Septiembre", "Octubre", "Noviembre", "Diciembre"}

	if (mes < 1 || mes > 12) {
		fmt.Println("No me rompas")
	} else {
		println(array[mes-1])
	}

}
