package usecases

import (
	"errors"
	"github.com/jackc/pgx/v5"
	"github.com/maxwelbm/transinterdigital/internal/domain/entity"
	"github.com/maxwelbm/transinterdigital/pkg/helper"
	"net/http"
	"time"
)

type TransferInput struct {
	AccountOriginID      int64
	AccountDestinationID int64
	Amount               float64
}

func (c *useCase) TransferAccountToAnother(input TransferInput) *helper.Response {
	balanceOrigin, err := c.repository.account.Balance(int(input.AccountOriginID))
	if err != nil {
		if errors.Is(pgx.ErrNoRows, err) {
			return &helper.Response{Status: http.StatusNotFound, Err: err}
		}
		return &helper.Response{Status: http.StatusInternalServerError, Err: errors.New("failed in get item")}
	}

	if input.AccountDestinationID == input.AccountOriginID {
		return &helper.Response{Status: http.StatusBadRequest, Err: errors.New("there is no way to transfer it to yourself")}
	}

	if input.Amount > balanceOrigin {
		return &helper.Response{Status: http.StatusBadRequest, Err: errors.New("insufficient balance to complete the transfer")}
	}

	transfer := entity.Transfers{
		AccountOriginID:      input.AccountOriginID,
		AccountDestinationID: input.AccountDestinationID,
		Amount:               input.Amount,
		CreatedAt:            time.Now(),
	}

	err = c.repository.transfer.Save(transfer)
	if err != nil {
		return &helper.Response{Status: http.StatusInternalServerError, Err: err}
	}

	err = c.repository.account.UpdateBalance(int(input.AccountOriginID), balanceOrigin-input.Amount)
	if err != nil {
		return &helper.Response{Status: http.StatusInternalServerError, Err: err}
	}

	balanceDestination, err := c.repository.account.Balance(int(input.AccountDestinationID))
	if err != nil {
		return &helper.Response{Status: http.StatusInternalServerError, Err: err}
	}

	err = c.repository.account.UpdateBalance(int(input.AccountDestinationID), balanceDestination+input.Amount)
	if err != nil {
		return &helper.Response{Status: http.StatusInternalServerError, Err: err}
	}

	return nil
}
