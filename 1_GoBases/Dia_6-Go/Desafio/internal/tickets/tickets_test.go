package tickets

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetTotalTickets(t *testing.T) {
	Pasajes = []Pasaje{{"1", "Tait Mc Caughan", "tmc0@scribd.com", "Finland", "17:11", 785},
		{"2", "Padget McKee", "pmckee1@hexun.com", "China", "20:19", 537},
		{"3", "Yalonda Jermyn", "yjermyn2@omniture.com", "China", "18:11", 579},
		{"4", "Diannne Pharrow", "dpharrow3@icio.us", "Mongolia", "23:16", 1238},
		{"5", "Carmella Tapenden", "ctapenden7@wikipedia.org", "China", "0:33", 713}}

	pais := "Pepe Argento"
	expectedPais := 0
	expectedError := ErrPaisNoEncontrado

	result, err := GetTotalTickets(pais)

	assert.Equal(t, expectedPais, result)
	assert.Equal(t, expectedError, err)

	pais2 := "China"
	expectedPais2 := 3
	var expectedError2 error = nil

	result2, err2 := GetTotalTickets(pais2)

	assert.Equal(t, expectedPais2, result2)
	assert.Equal(t, expectedError2, err2)

}

func TestGetHours(t *testing.T) {
	pais := "Pepe Argento"
	expectedHours := 0
	expectedError := ErrHorario

	result, err := GetHours(pais)

	assert.Equal(t, expectedHours, result)
	assert.Equal(t, expectedError, err)

	pais2 := "Noche"
	expectedHours2 := 2
	var expectedError2 error = nil

	result2, err2 := GetHours(pais2)

	assert.Equal(t, expectedHours2, result2)
	assert.Equal(t, expectedError2, err2)

}

func TestAverageDestination(t *testing.T) {
	pais := "Pepe Argento"
	expectedAverage := 0.0
	expectedError := ErrPaisNoEncontrado

	result, err := AverageDestination(pais)

	assert.Equal(t, expectedAverage, result)
	assert.Equal(t, expectedError, err)

	pais2 := "China"
	expectedAverage2 := 60.00
	var expectedError2 error = nil

	result2, err2 := AverageDestination(pais2)

	assert.Equal(t, expectedAverage2, result2)
	assert.Equal(t, expectedError2, err2)

}
