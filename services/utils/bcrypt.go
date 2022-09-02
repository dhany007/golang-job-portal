package utils

import "golang.org/x/crypto/bcrypt"

func HashPassword(pass string) (result string, err error) {
	var (
		salt int
		hash []byte
	)

	salt = 8
	password := []byte(pass)

	hash, err = bcrypt.GenerateFromPassword(password, salt)
	if err != nil {
		return "", err
	}

	return string(hash), nil
}
