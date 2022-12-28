package main

import "fmt"

type myerror struct {
}

func (e *myerror) Error() string {
	return fmt.Sprintf("Error: el salario ingresado no alcanza el minimo imponible")
}

func main() {
	var salary int
	var err *myerror

	fmt.Scanln(&salary)

	if salary < 150000 {
		fmt.Println(err)
	}
}
