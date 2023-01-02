package tickets

import (
	"errors"
	"strconv"
	"strings"
)

var Pasajes = make([]Pasaje, 0)

type Pasaje struct {
	Id          string
	Name        string
	Email       string
	Destination string
	HoraV       string
	Price       float64
}

var (
	ErrPaisNoEncontrado = errors.New("ERROR: No se encontro ningun pasaje con ese pais")
	ErrHorario          = errors.New("No existe ese horario")
)

func GetTotalTickets(destination string) (total int, err error) {
	for _, pass := range Pasajes {
		if pass.Destination == destination {
			total++
		}
	}
	if total == 0 {
		err = ErrPaisNoEncontrado
	}
	return
}

func GetHours(time string) (cant int, err error) {
	switch time {
	case "Madrugada":
		return getEarlyMornings(), nil
	case "Mañana":
		return getMornings(), nil
	case "Tarde":
		return getMidday(), nil
	case "Noche":
		return getEvenings(), nil
	default:

		return 0, ErrHorario
	}
	return
}

func getEarlyMornings() (cant int) {
	for _, pass := range Pasajes {
		hourStrings := strings.Split(pass.HoraV, ":")
		hourInt, _ := strconv.Atoi(hourStrings[0])
		if hourInt >= 0 && hourInt <= 6 {
			cant++
		}
	}
	return
}

func getMornings() (cant int) {
	for _, pass := range Pasajes {
		hourStrings := strings.Split(pass.HoraV, ":")
		hourInt, _ := strconv.Atoi(hourStrings[0])
		if hourInt >= 7 && hourInt <= 12 {
			cant++
		}
	}
	return
}

func getMidday() (cant int) {
	for _, pass := range Pasajes {
		hourStrings := strings.Split(pass.HoraV, ":")
		hourInt, _ := strconv.Atoi(hourStrings[0])
		if hourInt >= 13 && hourInt <= 19 {
			cant++
		}
	}
	return
}

func getEvenings() (cant int) {
	for _, pass := range Pasajes {
		hourStrings := strings.Split(pass.HoraV, ":")
		hourInt, _ := strconv.Atoi(hourStrings[0])
		if hourInt >= 20 && hourInt <= 23 {
			cant++
		}
	}
	return
}

func AverageDestination(destination string) (float64, error) {
	cant := 0.0
	for _, pass := range Pasajes {
		if pass.Destination == destination {
			cant++
		}
	}
	if cant == 0 {
		return 0, ErrPaisNoEncontrado
	}

	return cant * 100 / float64(len(Pasajes)), nil
}
