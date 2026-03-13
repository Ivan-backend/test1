package main

import (
	"context"
	"fmt"
	"main/feature_postgres/simple_connection"
	"main/feature_postgres/simple_sql"
	"main/http"
	"main/subscribes"
	"main/users"
)

// @title Blueprint Swagger API
// @version 1.0
// @description Swagger API for Golang Project Blueprint.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.email martin7.heinz@gmail.com

// @license.name MIT
// @license.url https://github.com/MartinHeinz/go-project-blueprint/blob/master/LICENSE

// @BasePath /api/v1

func main() {
	ctx := context.Background()
	subscribes := subscribes.NewList()
	users := users.NewList()
	httpHandlers := http.NewHTTPHandlers(users, subscribes)
	httpServer := http.NewHTTPServer(httpHandlers)

	conn, err := simple_connection.CheckConnection(ctx)

	if err != nil {
		panic(err)
	}

	if err := simple_sql.CreateTable(ctx, conn); err != nil {
		panic(err)
	}

	if err := httpServer.StartServer(); err != nil {
		fmt.Println("failed to start server:", err)
	}
}
