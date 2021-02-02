package model

import (
	"github.com/asaskevich/govalidator"
	uuid "github.com/satori/go.uuid"
	"time"
)

//Working with heritage; `valid:"-"` = validação não necessária
type Bank struct {
	Base     `valid:"required"`
	Code     string     `json:"code" valid:"notnull"`
	Name     string     `json:"name" valid:"notnull"`
	Accounts []*Account `valid:"-"`
}

func (bank *Bank) isValid() error { //creating method
	_, err := govalidator.ValidateStruct(bank)
	if err != nil {
		return err
	}
	return nil
}


// NewBank : creating a new Bank
func NewBank(code string, name string) (*Bank, error) {
	bank := Bank{
		Code: code,
		Name: name,
	}

	bank.ID = uuid.NewV4().String()
	bank.CreatedAt = time.Now()

	err := bank.isValid() //checking for errors

	if err != nil {
		return nil, err
	}
	return &bank, nil
}
