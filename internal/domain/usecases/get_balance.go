package usecases

import (
	"errors"
	"github.com/jackc/pgx/v5"
	"github.com/maxwelbm/transinterdigital/pkg/helper"
	"net/http"
)

type BalanceOutput struct {
	Balance float64
}

func (c *useCase) GetBalance(accountID int) (BalanceOutput, *helper.Response) {
	balance, err := c.repository.account.Balance(accountID)
	if err != nil {
		if errors.Is(pgx.ErrNoRows, err) {
			return BalanceOutput{0}, &helper.Response{Status: http.StatusNotFound, Err: err}
		}
		return BalanceOutput{0}, &helper.Response{Status: http.StatusInternalServerError, Err: errors.New("failed in get balance")}
	}
	return BalanceOutput{balance}, nil
}
