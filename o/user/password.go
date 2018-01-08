package user

import (
	"golang.org/x/crypto/bcrypt"
)

type password string

const LENGTH = 10

func (p password) gererateHashedPassword() (password, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(p), LENGTH)
	return password(hashed), err
}

func (p password) comparePassword(pwd string) error {
	return bcrypt.CompareHashAndPassword([]byte(p), []byte(pwd))
}
