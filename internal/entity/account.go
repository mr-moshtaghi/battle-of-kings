package entity

type Account struct {
	ID        int64  `json:"id"`
	FirstName string `json:"first_name"`
	Username  string `json:"username"`

	DisplayName string `json:"display_name"`
}

func (a Account) EntityID() ID {
	return NewID("account", a.ID)
}
