package feature_postgres

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

func CheckConnection() {
	ctx := context.Background()
	conn, err := pgx.Connect(ctx, "postgres://postgres:pass@localhost:5432/subscribes")
	if err != nil {
		panic(err)
	}

	if err := conn.Ping(ctx); err != nil {
		panic(err)
	}

	fmt.Println("Успешно.")
}

func CreateConnection(ctx context.Context) {

}
