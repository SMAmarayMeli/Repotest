package main

import (
	"fmt"
	"strings"
)

func main() {
	palabra := "academia"

	fmt.Println("La palabra tiene ", len(palabra), " letras")
	p := strings.Split(palabra, "")
	for _, a := range p {
		fmt.Println(a)
	}
}
