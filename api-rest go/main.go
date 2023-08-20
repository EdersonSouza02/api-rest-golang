package main

import (
	"api-rest-go/database"
	"api-rest-go/routes"
	"fmt"
)

func main() {

	database.ConectaComBancoDeDados()

	fmt.Println("Iniciando o servidor Rest com GO")
	routes.HandleRequest()
}
