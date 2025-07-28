package utils

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		return "", nil
	}
	return string(bytes), nil
}

func CheckHashedPassword(hash string, password string) bool {
	return nil == bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
}
