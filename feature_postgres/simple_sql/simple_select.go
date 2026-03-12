package simple_sql

import (
	"context"

	"github.com/jackc/pgx/v5"
)

func GetByServiceNameSubscribe(ctx context.Context, conn *pgx.Conn, serviceName string) ([]SubscribeModel, int, error) {
	sqlQuery := `
		SELECT * FROM subscribes
		WHERE servicename = $1
	`

	rows, err := conn.Query(ctx, sqlQuery, serviceName)
	if err != nil {
		return nil, 0, err
	}

	defer rows.Close()

	subscribes := make([]SubscribeModel, 0)
	price := 0

	for rows.Next() {
		var subscribe SubscribeModel

		err := rows.Scan(&subscribe.Id, &subscribe.ServiceName, &subscribe.Price, &subscribe.UserId, &subscribe.DateStart, &subscribe.DateEnd)
		if err != nil {
			return nil, 0, err
		}

		price += subscribe.Price
		subscribes = append(subscribes, subscribe)
	}

	return subscribes, price, nil
}

func GetByServiceNameAndDateSubscribe(ctx context.Context, conn *pgx.Conn, serviceName string, dateSubscribeStart string, dateSubscribeEnd string) ([]SubscribeModel, int, error) {
	sqlQuery := `
		SELECT * FROM subscribes 
		WHERE servicename = $1
		AND datestart >= $2
		AND datestart <= $3
	`

	rows, err := conn.Query(ctx, sqlQuery, serviceName, dateSubscribeStart, dateSubscribeEnd)
	if err != nil {
		return nil, 0, err
	}

	defer rows.Close()

	subscribes := make([]SubscribeModel, 0)
	price := 0

	for rows.Next() {
		var subscribe SubscribeModel

		err := rows.Scan(&subscribe.Id, &subscribe.ServiceName, &subscribe.Price, &subscribe.UserId, &subscribe.DateStart, &subscribe.DateEnd)
		if err != nil {
			return nil, 0, err
		}

		price += subscribe.Price
		subscribes = append(subscribes, subscribe)
	}

	return subscribes, price, nil
}

func GetByUserIdSubscribe(ctx context.Context, conn *pgx.Conn, userId string) ([]SubscribeModel, int, error) {
	sqlQuery := `
		SELECT * FROM subscribes 
		WHERE userid = $1
	`

	rows, err := conn.Query(ctx, sqlQuery, userId)
	if err != nil {
		return nil, 0, err
	}

	defer rows.Close()

	subscribes := make([]SubscribeModel, 0)
	price := 0

	for rows.Next() {
		var subscribe SubscribeModel

		err := rows.Scan(&subscribe.Id, &subscribe.ServiceName, &subscribe.Price, &subscribe.UserId, &subscribe.DateStart, &subscribe.DateEnd)
		if err != nil {
			return nil, 0, err
		}

		price += subscribe.Price
		subscribes = append(subscribes, subscribe)
	}

	return subscribes, price, nil
}

func GetByUserIdAndDateSubscribe(ctx context.Context, conn *pgx.Conn, userId string, dateSubscribeStart string, dateSubscribeEnd string) ([]SubscribeModel, int, error) {
	sqlQuery := `
		SELECT * FROM subscribes 
		WHERE userid = $1
		AND datestart >= $2
		AND datestart <= $3
	`

	rows, err := conn.Query(ctx, sqlQuery, userId, dateSubscribeStart, dateSubscribeEnd)
	if err != nil {
		return nil, 0, err
	}

	defer rows.Close()

	subscribes := make([]SubscribeModel, 0)
	price := 0

	for rows.Next() {
		var subscribe SubscribeModel

		err := rows.Scan(&subscribe.Id, &subscribe.ServiceName, &subscribe.Price, &subscribe.UserId, &subscribe.DateStart, &subscribe.DateEnd)
		if err != nil {
			return nil, 0, err
		}

		price += subscribe.Price
		subscribes = append(subscribes, subscribe)
	}

	return subscribes, price, nil
}
