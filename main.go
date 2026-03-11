package main

import (
	"fmt"
	"main/feature_postgres"
	"main/http"
	"main/subscribes"
	"main/users"
)

func main() {
	subscribes := subscribes.NewList()
	users := users.NewList()
	httpHandlers := http.NewHTTPHandlers(users, subscribes)
	httpServer := http.NewHTTPServer(httpHandlers)

	feature_postgres.CheckConnection()

	if err := httpServer.StartServer(); err != nil {
		fmt.Println("failed to start server:", err)
	}
}
