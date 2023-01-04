package main

import (
	"errors"
	"fmt"
	"os"
)

var (
	err = errors.New("ERROR: El minimo imponible es de 150000")
)

func checkSalario(a int) (bool, error) {
	if a < 150000 {
		return false, fmt.Errorf("%w y el salario ingresado de: %v", err, a)
	}
	return true, nil
}

func main() {
	var salary int
	var err1 error

	fmt.Scanln(&salary)

	_, err1 = checkSalario(salary)

	if errors.Is(err1, err) {
		fmt.Println(err1)
		os.Exit(0)
	}
	fmt.Printf("Salario es $%v\n", salary)
}
