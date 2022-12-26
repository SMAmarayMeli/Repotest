package main

import "fmt"

type Producto interface {
	Precio() float64
}

type Grande struct {
	costo float64
}
type Mediano struct {
	costo float64
}
type Peque struct {
	costo float64
}

func (p Peque) Precio() float64 {
	return p.costo
}

func (p Grande) Precio() float64 {
	return p.costo*1.06 + 2500.0
}

func (p Mediano) Precio() float64 {
	return p.costo * 1.03
}

func factory(tipo string, costo float64) float64 {

	var precio Producto
	switch tipo {
	case "grande":
		precio = Grande{costo}
	case "mediano":
		precio = Mediano{costo}
	case "peque":
		precio = Peque{costo}
	}
	return precio.Precio()
}

func main() {
	fmt.Printf("Precio grande: %v\n", factory("grande", 5000))
	fmt.Printf("Precio mediano: %v\n", factory("mediano", 5000))
	fmt.Printf("Precio peque: %v\n", factory("peque", 5000))
}
