package main

import (
	"errors"
	"fmt"
	"os"
)

var (
	errNegativo     = errors.New("ERROR: la hora fue negativa")
	errMenorOchenta = errors.New("ERROR: la hora fue menor a 80")
)

func main() {
	var horas, valorHora, salario float64
	var err error
	fmt.Println("Ingrese horas y valor de la hora")
	fmt.Scanln(&horas)
	fmt.Scanln(&valorHora)

	salario, err = calcularSalario(horas, valorHora)

	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
	fmt.Println("Salario: $", salario)
}

func calcularSalario(horas float64, valorHora float64) (float64, error) {

	if horas < 0 {
		return 0, errNegativo
	} else if horas < 80 {
		return 0, errMenorOchenta
	}

	salario := horas * valorHora

	if salario > 150000 {
		return salario * 0.9, nil
	} else {
		return salario, nil
	}
}
