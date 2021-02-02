package model

import (
	"errors"
	"time"

	"github.com/asaskevich/govalidator"
	uuid "github.com/satori/go.uuid"
)

const (
	TransactionPending   string = "pending"
	TransactionCompleted string = "completed"
	TransactionError     string = "error"
	TransactionConfirmed string = "confirmed"
)

// TransactionRepositoryInterface : Mostrando para os getways o que temos na infra.
type TransactionRepositoryInterface interface {
	Register(transaction *Transaction) error
	Save(transaction *Transaction) error
	Find(id string) (*Transaction, error)
}

// Transactions : adiciona transaction a lista.
type Transactions struct {
	Transaction []Transaction
}

// Transaction : estrutura para a transação.
type Transaction struct {
	Base              `valid:"required"`
	AccountFrom       *Account `valid:"-"`
	Amount            float64  `json:"amount" valid:"notnull"`
	PixKeyTo          *PixKey  `valid:"-"`
	PixKeyIDTo 				string   `gorm:"column:pix_key_id;type:uuid" valid:"notnull"`
	Status            string   `json:"status" valid:"notnull"`
	Description       string   `json:"description" valid:"notnull"`
	CancelDescription string   `json:"cancel_description" valid:"-"`
}

func (t *Transaction) isValid() error {
	_, err := govalidator.ValidateStruct(t)

	if t.Amount <= 0 { //bloqueando transferência por valor 0
		return errors.New("the amount must be greater than 0")
	}

	//validação de status
	if t.Status != TransactionPending && t.Status != TransactionCompleted && t.Status != TransactionError { //TODO: better validation
		return errors.New("invalid status for the transaction")
	}

	//regra de negócio
	if t.PixKeyTo.AccountID == t.AccountFrom.ID {
		return errors.New("the source and destination account cannot be the same")
	}

	if err != nil {
		return err
	}
	return nil
}

// NewTransaction : creates a new transaction.
func NewTransaction(accountFrom *Account, amount float64, pixKeyTo *PixKey, description string) (*Transaction, error) {
	transaction := Transaction{
		AccountFrom: accountFrom,
		Amount:      amount,
		PixKeyTo:    pixKeyTo,
		Status:      TransactionPending,
		Description: description,
	}

	transaction.ID = uuid.NewV4().String()
	transaction.CreatedAt = time.Now()

	err := transaction.isValid()
	if err != nil {
		return nil, err
	}

	return &transaction, nil
}

// Complete : emmits the signal for a transaction completed.
func (t *Transaction) Complete() error {
	t.Status = TransactionCompleted
	t.UpdatedAt = time.Now()
	err := t.isValid() //validação
	return err //return only IF have error
}

// Confirm : emmits the signal for a transaction confirmed.
func (t *Transaction) Confirm() error {
	t.Status = TransactionConfirmed
	t.UpdatedAt = time.Now()
	err := t.isValid()
	return err
}

// Cancel : emmits the signal for a transaction canceled.
func (t *Transaction) Cancel(description string) error {
	t.Status = TransactionError
	t.UpdatedAt = time.Now()
	t.Description = description
	err := t.isValid()
	return err
}
