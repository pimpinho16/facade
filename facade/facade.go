package facade

import (
	"facade/email"
	"facade/storage"
	"facade/validate"
	"log"
)

/***
Facade
estructura que debe conocer todo lo que necesitan los subsistemas
*/
type Facade struct {
	to                  string
	comment             string
	validatorToken      validate.ValidateToken
	validatorPermission validate.ValidatePermission
	store               storage.Storage
	notificador         email.Email
}

func New(to, comment, token, user string) Facade {
	return Facade{
		to:                  to,
		comment:             comment,
		validatorToken:      validate.NewValidatorToken(token),
		validatorPermission: validate.NewValidatePermission(user),
		store:               storage.NewStorage("mysql"),
		notificador:         email.NewEmail(),
	}
}

//Comment
/*
es el método que el cliente llamara para ejecutar todas las acciones. El cliente no sabrá que es lo que se hace
*/
func (f *Facade) Comment() error {
	if err := f.validatorToken.Validate(); err != nil {
		log.Fatal(err)
		return err
	}

	if err := f.validatorPermission.Validate(); err != nil {
		log.Fatal(err)
		return err
	}

	f.store.Save(f.comment)
	f.notificador.Send(f.to, f.comment)
	return nil
}

//Notify
/*
facade conoce todos los procesos que se realizan, asi que es posible dividir en procesos pequeños si asi se
desea
*/
func (f *Facade) Notify() {
	f.notificador.Send(f.to, f.comment)
}
