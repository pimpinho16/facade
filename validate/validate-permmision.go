package validate

import (
	"errors"
	"fmt"
)

var ErrPermissionNotValid = errors.New("El usuario está logeado")

type ValidatePermission struct {
	userId string
}

func NewValidatePermission(ID string) ValidatePermission {
	return ValidatePermission{userId: ID}
}
func (vp *ValidatePermission) Validate() error {
	if vp.userId != "user-blog" {
		return ErrPermissionNotValid
	}

	fmt.Println("usuario autorizado")
	return nil
}
