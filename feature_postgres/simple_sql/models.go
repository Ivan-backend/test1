package simple_sql

import (
	"time"

	"github.com/google/uuid"
)

type SubscribeModel struct {
	Id          int
	ServiceName string
	Price       int
	UserId      uuid.UUID
	DateStart   time.Time
	DateEnd     time.Time
}
