package models

import "golang.org/x/crypto/bcrypt"

type Admin struct {
	Username     string `json:"username"`
	PasswordHash string `json:"password"`
}

func (a *Admin) GeneratePasswordHash(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return err
	}
	a.PasswordHash = string(bytes)
	return nil
}

func (a *Admin) checkPasswordHash(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(a.PasswordHash), []byte(password))
	return err == nil
}
