package helper

import "golang.org/x/crypto/bcrypt"

// for register
func HashPassword(password string) (string, error) {
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(passwordHash), err
}

// for login
func VerifyPassword(HashPassword string, password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(HashPassword), []byte(password))
	return err
}
