package http

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
)

type SubscribeServiceNameDTO struct {
	ServiceName string
}

type SubscribeDTO struct {
	ServiceName string
	Price       int
	UserId      uuid.UUID
	DateStart   time.Time
}

type ErrorDTO struct {
	Message string
	Time    time.Time
}

func (e ErrorDTO) ToString() string {
	b, err := json.MarshalIndent(e, "", "    ")
	if err != nil {
		panic(err)
	}

	return string(b)
}
