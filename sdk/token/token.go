package token

import (
	"strings"

	"github.com/google/uuid"
)

func GenerateToken() string {
	id := uuid.New()
	token := strings.ToUpper(id.String()[:8])

	return token
}
