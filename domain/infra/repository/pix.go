package repository

import (
	"fmt"

	"github.com/diogojorgebasso/imersao-fullstack-fullcycle/domain/model"
	"github.com/jinzhu/gorm"
)

// criando Interaface
/* type PixKeyRepositoryInterface interface {
	RegisterKey(pixKey *PixKey) (*PixKey, error)
	FindKeyByKind(key string, kind string) (*PixKey, error)
	AddBank(bank *Bank) error
	AddAccount(account *Account) error
	FindAccount(id string) (*Account, error)
} */

// PixKeyRepositoryDb estrutura o banco de dados
type PixKeyRepositoryDb struct{
	Db*gorm.DB
}

// AddBank : adding a new bank in Db
func (r PixKeyRepositoryDb) AddBank(bank *model.Bank) error{
	err := r.Db.Create(bank).Error
	if err != nil{
		return err
	}
	return nil
}

// AddAccount : adding a new bank in Db
func (r PixKeyRepositoryDb) AddAccount(account *model.Account) error {
	err := r.Db.Create(account).Error
	if err != nil {
		return err
	}
	return nil
}

// RegisterKey : adding a new bank in Db
func (r PixKeyRepositoryDb) RegisterKey(pixKey *model.PixKey) (*model.PixKey, error) {
	err := r.Db.Create(pixKey).Error
	if err != nil {
		return nil, err
	}
	return pixKey, nil
}

// FindKeyByKind : adding a new bank in Db
func (r PixKeyRepositoryDb) FindKeyByKind(key string, kind string) (*model.PixKey, error) {
	var pixKey model.PixKey
	r.Db.Preload("Account.Bank").First(&pixKey, "kind = ? and key = ?", kind, key)

	if pixKey.ID == "" {
		return nil, fmt.Errorf("no key was found")
	}
	return &pixKey, nil
}

// FindAccount : finding accounting in Db
func (r PixKeyRepositoryDb) FindAccount(id string) (*model.Account, error) {
	var account model.Account
	r.Db.Preload("Bank").First(&account, "id = ?", id)

	if account.ID == "" {
		return nil, fmt.Errorf("no account found")
	}
	return &account, nil
}

// FindBank : finding bank in Db
func (r PixKeyRepositoryDb) FindBank(id string) (*model.Bank, error) {
	var bank model.Bank
	r.Db.First(&bank, "id = ?", id)

	if bank.ID == "" {
		return nil, fmt.Errorf("no bank found")
	}
	return &bank, nil
}