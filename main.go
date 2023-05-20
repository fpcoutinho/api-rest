package main

import (
	"api-rest/database"
	"api-rest/routes"
	"fmt"
)

func main() {
	database.ConnectDB()
	fmt.Println("Servidor online em http://localhost:8000")
	routes.HandleRequest()
}
