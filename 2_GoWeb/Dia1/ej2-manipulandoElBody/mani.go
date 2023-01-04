package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type per struct {
	Nombre   string `json:"nombre"`
	Apellido string `json:"apellido"`
}

var Persona per

func postSaludo(c *gin.Context) {
	var newPersona per
	if err := c.BindJSON(&newPersona); err != nil {
		return
	}
	Persona = newPersona
	c.IndentedJSON(http.StatusCreated, newPersona)
}

func getSaludo(c *gin.Context) {
	a := fmt.Sprintf("Hola %v %v", Persona.Apellido, Persona.Nombre)
	c.String(200, a)
}

func main() {

	router := gin.Default()
	router.POST("/saludo", postSaludo)
	router.GET("/saludo", getSaludo)

	router.Run(":8081")
}
