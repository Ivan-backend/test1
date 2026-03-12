package simple_sql

import (
	"context"
	"time"

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

func InsertRow(ctx context.Context, conn *pgx.Conn, serviceName string, price int, userId string, dateStart time.Time, dateEnd time.Time) error {
	sqlQuery := `
		INSERT INTO subscribes(servicename, price, userid, datestart, dateend)
		VALUES($1, $2, $3, $4, $5);
	`

	_, err := conn.Exec(ctx, sqlQuery, serviceName, price, userId, dateStart, dateEnd)

	return err
}

func GetSubscribes(ctx context.Context, conn *pgx.Conn) ([]SubscribeModel, error) {
	sqlQuery := `
		SELECT * FROM subscribes
	`

	rows, err := conn.Query(ctx, sqlQuery)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	subscribes := make([]SubscribeModel, 0)

	for rows.Next() {
		var subscribe SubscribeModel

		err := rows.Scan(&subscribe.Id, &subscribe.ServiceName, &subscribe.Price, &subscribe.UserId, &subscribe.DateStart, &subscribe.DateEnd)
		if err != nil {
			return nil, err
		}

		subscribes = append(subscribes, subscribe)
	}

	return subscribes, nil
}

