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

	simple_sql.CreateTable(ctx, conn)

	if err := httpServer.StartServer(); err != nil {
		fmt.Println("failed to start server:", err)
	}
}
