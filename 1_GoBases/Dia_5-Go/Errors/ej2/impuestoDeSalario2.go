package main

import (
	"errors"
	"fmt"
)

type myerror struct {
}

var salaryError = myerror{}

func (e myerror) Error() string {
	return fmt.Sprint("Error: el salario es menor a 10000")
}

func salaryBigOrSmall(a int) (bool, error) {
	e := salaryError
	if a <= 10000 {
		return false, e
	}
	return true, nil
}

func main() {
	var salary int
	var err1 error

	fmt.Scanln(&salary)

	_, err1 = salaryBigOrSmall(salary)

	if errors.Is(salaryError, err1) {
		fmt.Println(err1)
	}
}
