package usecases

import (
	"errors"
	"github.com/maxwelbm/transinterdigital/internal/domain/entity"
	"time"
)

type TransferInput struct {
	AccountOriginID      int64
	AccountDestinationID int64
	Amount               float64
}

func (c *useCase) TransferAccountToAnother(input TransferInput) error {
	balanceOrigin, err := c.repository.account.Balance(int(input.AccountOriginID))
	if err != nil {
		return err
	}

	if input.AccountDestinationID == input.AccountOriginID {
		return errors.New("there is no way to transfer it to yourself")
	}

	if input.Amount > balanceOrigin {
		return errors.New("insufficient balance to complete the transfer")
	}

	transfer := entity.Transfers{
		AccountOriginID:      input.AccountOriginID,
		AccountDestinationID: input.AccountDestinationID,
		Amount:               input.Amount,
		CreatedAt:            time.Now(),
	}

	err = c.repository.transfer.Save(transfer)
	if err != nil {
		return err
	}

	err = c.repository.account.UpdateBalance(int(input.AccountOriginID), balanceOrigin-input.Amount)
	if err != nil {
		return err
	}

	balanceDestination, err := c.repository.account.Balance(int(input.AccountDestinationID))
	if err != nil {
		return err
	}

	err = c.repository.account.UpdateBalance(int(input.AccountDestinationID), balanceDestination+input.Amount)
	if err != nil {
		return err
	}

	return nil
}
