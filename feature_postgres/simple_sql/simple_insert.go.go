package simple_sql

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5"
)

func InsertRow(ctx context.Context, conn *pgx.Conn, serviceName string, price int, userId string, dateStart time.Time, dateEnd time.Time) error {
	sqlQuery := `
		INSERT INTO subscribes(servicename, price, userid, datestart, dateend)
		VALUES($1, $2, $3, $4, $5);
	`

	_, err := conn.Exec(ctx, sqlQuery, serviceName, price, userId, dateStart, dateEnd)

	return err
}
