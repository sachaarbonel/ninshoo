package model

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID        string
	Email     string
	Password  string
	CreatedAt string `db:"created_at"`
	Confirmed bool
	// EmailConfirmedAt string `db:"email_confirmed_at"`
	// EmailSentAt string `db:"email_confirmed_at"`
	//Roles     []*Role
}

func (user *User) HashedPassword() error {
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Println(err)
		return err
	}
	user.Password = string(hash)
	return nil
}

func (user *User) ComparePassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}