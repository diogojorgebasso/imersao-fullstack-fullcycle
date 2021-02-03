package transaction

import (
	"fmt"

	"github.com/diogojorgebasso/imersao-fullstack-fullcycle/domain/model"
	"github.com/jinzhu/gorm"
)

/* type UserRepositoryInterface interface {
	Register(user *User) error
	Save(user *User) error
	Find(id string) (*User, error)
} */

type TransactionRepositoryDb struct {
	Db *gorm.DB
}

func (t *TransactionRepositoryDb) Register(transaction *model.Transaction) error {
	err := t.Db.Create(transaction).Error
	if err != nil {
		return err
	}
	return nil
}

func (t *TransactionRepositoryDb) Save(transaction *model.Transaction) error {
	err := t.Db.Save(transaction).Error
	if err != nil {
		return err
	}
	return nil
}

func (t *TransactionRepositoryDb) Find(id string) (*model.Transaction, error) {
	var transaction model.Transaction
	t.Db.Preload("AccountFrom.Bank").First(&transaction, "id = ?", id) //da onde vem a informação e de qual módulo

	if transaction.ID == "" {
		return nil, fmt.Errorf("No transaction was found")
	}
	return &transaction, nil
}