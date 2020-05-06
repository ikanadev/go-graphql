package model

import "golang.org/x/crypto/bcrypt"

// User Represents user in db
type User struct {
	ID       string `json:"id" pg:",pk,type:bigserial"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// HashPassword hashes the password and assign it in
func (u *User) HashPassword(password string) error {
	bytePassword := []byte(password)
	passwordHash, err := bcrypt.GenerateFromPassword(bytePassword, bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(passwordHash)
	return nil
}

// CheckPassword check if password is correct
func (u *User) CheckPassword(password string) error {
	bytePassw := []byte(password)
	byteHashedPasswd := []byte(u.Password)
	return bcrypt.CompareHashAndPassword(byteHashedPasswd, bytePassw)
}
