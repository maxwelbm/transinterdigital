package entity

import "time"

type Account struct {
	ID        int64
	Name      string
	CPF       string
	Secret    string
	Balance   float64
	CreatedAt time.Time
}

type AccountRepository interface {
	Save(account *Account) error
	Balance(accountID int) (float64, error)
	List() ([]Account, error)
	UpdateBalance(accountID int, balance float64) error
	GetAccountID(cpf, secret string) (int64, error)
}
