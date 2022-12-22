package main

import "fmt"

func main() {
	var employees = map[string]int{"Benjamin": 20, "Nahuel": 26, "Brenda": 19, "Darío": 44, "Pedro": 30}

	fmt.Println("Edad Benjamin: ", employees["Benjamin"])

	var cont int
	for _, v := range employees {
		if v > 21 {
			cont++
		}
	}
	fmt.Println("Empleados mayores a 21: ", cont)
	employees["Federico"] = 21
	fmt.Println("Federico a sido agregado: ", employees)
	delete(employees, "Pedro")
	fmt.Println("Pedro a sido eliminado: ", employees)
}
