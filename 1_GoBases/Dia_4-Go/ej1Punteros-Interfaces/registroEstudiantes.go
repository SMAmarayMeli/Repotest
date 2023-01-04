package main

import "fmt"

type Persona struct {
	Name    string
	Surname string
	Dni     int
}

type Alumno struct {
	Info Persona
	Date string
}

func (a Alumno) detalle() {
	fmt.Printf("Nombre %v\nApellido %v\nDni %v\nFecha ingreso %v\n\n",
		a.Info.Name, a.Info.Surname, a.Info.Dni, a.Date)
}

func main() {
	alumnos := []Alumno{
		Alumno{
			Persona{"pepe", "sanchez", 4857348}, "24/11/22"},
		Alumno{
			Persona{"papo", "martinez", 4545456}, "24/12/22"},
		Alumno{
			Persona{"pepu", "fernandez", 7453567}, "24/10/22"}}

	for _, a := range alumnos {
		a.detalle()
	}

}
