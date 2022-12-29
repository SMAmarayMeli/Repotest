package main

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

type Client struct {
	name    string
	dni     string
	phone   string
	legajo  string
	address string
}

var clients = make([]Client, 0)

func loadData(data []string) (err error) {
	for _, user := range data {
		client := strings.Split(user, ",")
		if len(client) == 1 {
			return
		}
		if len(client) != 5 {
			err = errors.New("el archivo esta corrupto")
			return
		}
		clientName := client[0]
		clientDni := client[1]
		clientPhone := client[2]
		clientLegajo := client[3]
		clientAddress := client[4]
		newClient := Client{clientName, clientDni, clientPhone, clientLegajo, clientAddress}
		clients = append(clients, newClient)
	}
	return
}

func initSistem() {
	txt, err := os.ReadFile("customers.txt")
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

func sistemEnd() {
	err := recover()

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("ejecución finalizada")
}

func isExistDni(dni string) (err error) {
	for _, client := range clients {
		if client.dni == dni {
			err = errors.New("el cliente ya existe")
			return
		}
	}
	return
}

func (c Client) save() (err error) {
	if c.dni == "" {
		return errors.New("el atributo dni no puede ser nulo")
	}
	errDni := isExistDni(c.dni)
	if errDni != nil {
		return errDni
	}
	if c.name == "" {
		return errors.New("el atributo name no puede ser nulo")
	}
	if c.phone == "" {
		return errors.New("el atributo phone no puede ser nulo")
	}
	if c.address == "" {
		return errors.New("el atributo address no puede ser nulo")
	}
	if c.legajo == "" {
		return errors.New("el atributo legajo no puede ser nulo")
	}
	clients = append(clients, c)
	updateCustomers(c)
	return
}

func updateCustomers(client Client) {
	text := fmt.Sprintf("%s,%s,%s,%s,%s", client.name, client.dni, client.phone, client.address, client.legajo)

	f, err := os.OpenFile("customers.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		panic(err)
	}

	defer f.Close()
	if len(clients) != 1 {
		text = fmt.Sprintf("\n%s", text)
	}
	if _, err = f.WriteString(text); err != nil {
		panic(err)
	}
}

func main() {
	defer sistemEnd()
	initSistem()

	fmt.Printf("%+v\n", clients)

	for {
		newClient := Client{}
		fmt.Println("inserte el name:")
		fmt.Scanln(&newClient.name)
		fmt.Println("inserte el dni:")
		fmt.Scanln(&newClient.dni)
		fmt.Println("inserte el phone:")
		fmt.Scanln(&newClient.phone)
		fmt.Println("inserte el address:")
		fmt.Scanln(&newClient.address)
		fmt.Println("inserte el legajo:")
		fmt.Scanln(&newClient.legajo)
		err := newClient.save()
		if err != nil {
			panic(err)
		}
		fmt.Println("cliente agregado exitosamente.")
	}

}
