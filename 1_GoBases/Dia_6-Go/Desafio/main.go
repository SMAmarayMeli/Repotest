package main

import (
	"errors"
	"fmt"
	"github.com/bootcamp-go/desafio-go-bases/internal/tickets"
	"os"
	"strconv"
	"strings"
)

func sistemEnd() {
	err := recover()

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("ejecución finalizada")
}
func loadData(data []string) (err error) {
	for _, user := range data {
		pass := strings.Split(user, ",")
		if len(pass) == 1 {
			return
		}
		if len(pass) != 6 {
			err = errors.New("el archivo esta corrupto")
			return
		}
		passId := pass[0]
		passName := pass[1]
		passEmail := pass[2]
		passDestination := pass[3]
		passHora := pass[4]
		passPrice := pass[5]

		var auxPassPrice float64
		auxPassPrice, err = strconv.ParseFloat(passPrice, 64)

		newPass := tickets.Pasaje{passId, passName, passEmail, passDestination, passHora, auxPassPrice}
		tickets.Pasajes = append(tickets.Pasajes, newPass)
	}
	return
}

func initSistem() {
	txt, err := os.ReadFile("tickets.csv")
	if err != nil {
		panic("el archivo indicado no fue encontrado o está dañado")
	}
	data := string(txt)
	users := strings.Split(data, "\n")
	errLoad := loadData(users)
	if errLoad != nil {
		panic(errLoad)
	}
}

func main() {
	defer sistemEnd()
	initSistem()

	// Req. 1
	fmt.Println("Ingrese pais para calcular cantidad de pasajes: ")
	var pais string
	fmt.Scanln(&pais)

	total, err := tickets.GetTotalTickets(pais)

	if err != nil {
		panic("Error en la busqueda total para un pais")
	}
	fmt.Printf("Para el pais %v hay un total de pasajes de: %v\n", pais, total)

	// Req. 2
	fmt.Println("Ingrese horario a calcular pasajes: ")
	var horario string
	fmt.Scanln(&horario)

	cantidadHorario, errHorario := tickets.GetHours(horario)

	if errHorario != nil {
		panic("Error en la busqueda de horario ingresado")
	}
	fmt.Printf("Cantidad de pasajes para horario %v es: %v\n", horario, cantidadHorario)

	// Req. 3
	fmt.Println("Ingrese pais a calcular promedio: ")
	var paisPromedio string
	fmt.Scanln(&paisPromedio)

	promedio, errPromedio := tickets.AverageDestination(paisPromedio)

	if errPromedio != nil {
		panic("Error en la busqueda de horario ingresado")
	}
	fmt.Printf("Promedio del total del pais %v es: %v porciento \n", paisPromedio, promedio)
}
