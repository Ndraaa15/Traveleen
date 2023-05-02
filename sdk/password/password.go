package password

import (
	"os"
	"strconv"

	"golang.org/x/crypto/bcrypt"
)

func GeneratePassword(Password string) (string, error) {
	saltRound, err := strconv.Atoi(os.Getenv("BCRYPT_SALT_ROUND"))
	if err != nil {
		return "", err
	}

	hashPassword, err := bcrypt.GenerateFromPassword([]byte(Password), saltRound)
	if err != nil {
		return "", err
	}

	return string(hashPassword), nil
}

func ComparePassword(passwordFound, passwordInput string) error {
	if err := bcrypt.CompareHashAndPassword([]byte(passwordFound), []byte(passwordInput)); err != nil {
		return err
	}

	return nil
}
