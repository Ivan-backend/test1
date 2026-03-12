package simple_sql

import (
	"context"

	"github.com/jackc/pgx/v5"
)

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
