package simple_sql

import (
	"context"

	"github.com/jackc/pgx/v5"
)

func CreateTable(ctx context.Context, conn *pgx.Conn) error {
	sqlQuery := `
		CREATE TABLE IF NOT EXISTS subscribes (
			id SERIAL PRIMARY KEY,
			serviceName VARCHAR(150) NOT NULL,
			price INT,
			userId VARCHAR(36),
			dateStart DATE NOT NULL,
			dateEnd DATE NOT NULL
		);
	`

	_, err := conn.Exec(ctx, sqlQuery)

	return err
}
