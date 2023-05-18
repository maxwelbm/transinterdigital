package entity

import (
	"errors"
	"github.com/jackc/pgx/v5"
)

type AccontMock struct{}

func (u AccontMock) Balance(accountID int) (float64, error) {
	if accountID == 1 {
		return 1000000.12, nil
	}
	if accountID == 2 {
		return 0, errors.New("failed in get balance")
	}
	if accountID == 3 {
		return 0, pgx.ErrNoRows
	}
	return 0, nil
}

func (u AccontMock) List() ([]Account, error) {
	return []Account{
		{ID: 1, Name: "Max", CPF: "12312312312", Secret: "maxsecret", Balance: 1000000.12},
		{ID: 2, Name: "Salty", CPF: "32132132132", Secret: "saltysecret", Balance: 1000000.13},
	}, nil
}

func (u AccontMock) UpdateBalance(accountID int, balance float64) error {
	return nil
}

func (u AccontMock) GetAccountID(cpf, secret string) (int64, error) {
	return 0, nil
}

func (u AccontMock) Save(account *Account) error {
	if len(account.Name) == 0 {
		return errors.New("failed in create account")
	}
	return nil
}
