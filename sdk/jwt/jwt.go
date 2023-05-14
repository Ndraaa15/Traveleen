package jwt

import (
	"errors"
	"gin/src/entity"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
)

func GenerateJWTToken(user entity.User) (string, error) {

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.Username,
		"exp": time.Now().Add(time.Hour * 1).Unix(),
		"id":  user.ID,
	})

	signedToken, err := token.SignedString([]byte(os.Getenv("JWT_SIGN")))
	if err != nil {
		return "", errors.New("FAILED TO GENERATE JWT TOKEN")
	}

	return signedToken, nil
}

func DecodeJWTToken(token string) (map[string]interface{}, error) {
	decoded, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SIGN")), nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := decoded.Claims.(jwt.MapClaims)

	if !ok {
		return nil, errors.New("FAILED TO DECODE JWT TOKEN")
	}
	if !decoded.Valid {
		return nil, errors.New("invalid JWT Token")
	}

	return claims, nil
}
