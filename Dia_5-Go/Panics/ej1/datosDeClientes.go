package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("customer.txt") // For read access.
	if err != nil {
		panic("Archivo no pudo ser abierto")
	}
	var data []byte
	count, err := file.Read(data)
	if err != nil {

	}

	fmt.Printf("read %d bytes: %q\n", count, data[:count])
}
