package validate

import (
	"errors"
	"fmt"
)

var ErrTokenNotFound = errors.New("El usuario está logeado")

type ValidateToken struct {
	token string
}

func NewValidatorToken(t string) ValidateToken {
	return ValidateToken{token: t}
}

func (vt *ValidateToken) Validate() error {
	if vt.token != "token-valido" {
		return ErrTokenNotFound
	}

	fmt.Println("token válido")
	return nil
}
