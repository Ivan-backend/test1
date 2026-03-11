package users

import (
	"github.com/google/uuid"
)

type List struct {
	users map[uuid.UUID]User
}

func NewList() *List {
	return &List{
		users: make(map[uuid.UUID]User),
	}
}

func (l *List) CreateUser(user User) error {
	l.users[user.Id] = user

	return nil
}

func (l *List) GetUsers() map[uuid.UUID]User {
	tmp := make(map[uuid.UUID]User, len(l.users))
	for k, v := range l.users {
		tmp[k] = v
	}

	return tmp
}
