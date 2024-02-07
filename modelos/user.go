package modelos

import (
	"time"
)

type User struct {
	Id        int
	Name      string
	CreatedAt time.Time
	Status    bool
}

// funciona como un constructor
func (user *User) UserConstructor(id int, name string, createAt time.Time, status bool) {
	user.Id = id
	user.Name = name
	user.CreatedAt = createAt
	user.Status = status
}
