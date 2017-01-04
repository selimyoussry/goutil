package goutil

import "golang.org/x/crypto/bcrypt"

const (
	BCRYPT_DEFAULT_COST = 7
)

func Encrypt(s string) (string, error) {
	passwordHashBytes, err := bcrypt.GenerateFromPassword([]byte(s), BCRYPT_DEFAULT_COST)
	if err != nil {
		return "", err
	}
	passwordHash := string(passwordHashBytes)
	return passwordHash, nil
}
