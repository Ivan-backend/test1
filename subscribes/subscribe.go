package subscribes

import (
	"time"

	"github.com/google/uuid"
)

type Subscribe struct {
	ServiceName string
	Price       int
	UserId      uuid.UUID
	DateStart   time.Time
	DateEnd     *time.Time
}

func NewSubscribe(serviceName string, price int, userId uuid.UUID, dateStart time.Time) Subscribe {
	return Subscribe{
		ServiceName: serviceName,
		Price:       price,
		UserId:      userId,
		DateStart:   dateStart,
		DateEnd:     nil,
	}
}
