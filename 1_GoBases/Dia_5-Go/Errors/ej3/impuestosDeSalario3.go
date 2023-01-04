package main

import (
	"errors"
	"fmt"
	"os"
)

var (
	errSalarySmall = errors.New("ERROR: Salario es menor a 10000")
)

func salaryBigOrSmall(a int) (bool, error) {
	if a <= 10000 {
		return false, errSalarySmall
	}
	return true, nil
}

func main() {
	var salary int
	var err1 error

	fmt.Scanln(&salary)

	_, err1 = salaryBigOrSmall(salary)

	if errors.Is(errSalarySmall, err1) {
		fmt.Println(err1)
		os.Exit(0)
	}
	fmt.Printf("Salario es $%v\n", salary)
}
