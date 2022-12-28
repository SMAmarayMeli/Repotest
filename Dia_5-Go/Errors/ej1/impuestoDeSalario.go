package main

import "fmt"

type myerror struct {
}

func (e *myerror) Error() (err string) {
	return fmt.Sprintf("Error: el salario ingresado no alcanza el minimo imponible")
}

func main() {
	var salary int

	fmt.Scanln(&salary)

	if salary < 15000 {
		fmt.Print()
	}

}
