package models

import (
	"api/src/security"
	"errors"
	"strings"
	"time"

	"github.com/badoux/checkmail"
)

type User struct {
	Id        uint8     `json:"id,omitempty"`
	Name      string    `json:"name,omitempty"`
	Nick      string    `json:"nick,omitempty"`
	Email     string    `json:"email,omitempty"`
	Password  string    `json:"password,omitempty"`
	CreatedAt time.Time `json:"createdAt,omitempty"`
}

func (u *User) validateUser(step string) error {
	if u.Name == "" {
		return errors.New("o nome é obrigatório")
	}
	if u.Nick == "" {
		return errors.New("o nick é obrigatório")
	}
	if u.Email == "" {
		return errors.New("o email é obrigatório")
	}

	erro := checkmail.ValidateFormat(u.Email)
	if erro != nil {
		return errors.New("o email é inválido")
	}

	if step == "cadastro" && u.Password == "" {
		return errors.New("a senha é obrigatória")
	}
	return nil
}

func (u *User) formatData(step string) error {
	u.Name = strings.TrimSpace(u.Name)
	u.Nick = strings.TrimSpace(u.Nick)
	u.Email = strings.TrimSpace(u.Email)

	if step == "cadastro" {
		hashedPassword, erro := security.HashPassword(u.Password)
		if erro != nil {
			return erro
		}
		u.Password = string(hashedPassword)
	}
	return nil
}

func (u *User) Prepare(step string) error {
	erro := u.validateUser(step)
	if erro != nil {
		return erro
	}
	
	erro = u.formatData(step)
	if erro != nil {
		return erro
	}
	return nil
}
