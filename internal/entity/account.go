package entity

import (
	"fmt"
	"time"
)

type Account struct {
	ID        int64     `json:"id"`
	FirstName string    `json:"first_name"`
	Username  string    `json:"username"`
	JoinedAt  time.Time `json:"joined_at"`

	DisplayName string `json:"display_name"`
}

func (a *Account) EntityID() ID {
	return ID(fmt.Sprintf("account:%d", a.ID))
}
