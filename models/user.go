package models

import (
	"errors"

	"github.com/google/uuid"
)

type User struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type Users struct {
	Users []User `json:"users"`
}

func (u *Users) Add(user *User) {
	user.ID = uuid.NewString()
	u.Users = append(u.Users, *user)
}

func (u *Users) GetByName(name string) (user User, err error) {
	if name == "" {
		return user, errors.New("O nome não pode ser vazio!")
	}

	for _, user := range u.Users {
		if user.Name == name {
			return user, nil
		}
	}

	err = errors.New("Usuário não encontrado.")
	return user, err
}

func (u *Users) Delete(email string) (user User, err error) {
	for pos, user := range u.Users {
		if user.Email == email {
			u.Users = append(u.Users[:pos], u.Users[pos+1:]...)
			return user, nil
		}
	}

	return user, errors.New("Usuário não encontrado.")
}
