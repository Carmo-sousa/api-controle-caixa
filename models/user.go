package models

import "github.com/google/uuid"

type User struct {
	ID    uuid.UUID `json:"id"`
	Name  string    `json:"name"`
	Email string    `json:"email"`
}

type Users struct {
	Users []User `json:"users"`
}

func (u *Users) Add(user *User) {
	user.ID = uuid.New()
	u.Users = append(u.Users, *user)
}
