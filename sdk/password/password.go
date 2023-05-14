package password

import (
	"errors"
	"os"
	"strconv"

	"golang.org/x/crypto/bcrypt"
)

func GeneratePassword(Password string) (string, error) {
	saltRound, err := strconv.Atoi(os.Getenv("BCRYPT_SALT_ROUND"))
	if err != nil {
		return "", errors.New("FAILED TO PARSE BCRYPT SALT ROUND")
	}

	hashPassword, err := bcrypt.GenerateFromPassword([]byte(Password), saltRound)
	if err != nil {
		return "", errors.New("FAILED TO GENERATE PASSWORD")
	}

	return string(hashPassword), nil
}

func ComparePassword(passwordFound, passwordInput string) error {
	if err := bcrypt.CompareHashAndPassword([]byte(passwordFound), []byte(passwordInput)); err != nil {
		return errors.New("PASSWORD DOESN'T MATCH")
	}

	return nil
}
