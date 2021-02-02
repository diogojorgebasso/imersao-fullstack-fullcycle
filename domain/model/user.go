//Antes de mais nada, desculpe-me por ter enviado aquela primeira versão apenas como fork e não como resolução propriamente do desafio.
package model

import (
	"github.com/asaskevich/govalidator"
	uuid "github.com/satori/go.uuid"
)

// User : structure
type User struct {
	ID        string     `json:"id" valid:"uuid"`
	Name      string     `json:"name" valid:"notnull"`
	Email     string     `json:"email" valid:"notnull"`
}

//FIXME: Seria mais limpo colocar a função abaixo?
//func init() {
//	govalidator.SetFieldsRequiredByDefault(true)
//}

func (user *User) isValid() error {
	_, err := govalidator.ValidateStruct(user)
	if err != nil {
		return err
	}
	return nil
}

// NewUser : is a function to create Users
func NewUser(name string, email string) (*User, error) {
	user := User{
		Email: email,
		Name: name,
	}
	user.ID = uuid.NewV4().String()
	user.Name = name
	user.Email = email

	err := user.isValid()

	if err != nil {
		return nil, err
	}

	return &user, nil
} 

type UserRepositoryInterface interface {
	Register(user *User) error
	Save(user *User) error
	Find(id string) (*User, error)
}
