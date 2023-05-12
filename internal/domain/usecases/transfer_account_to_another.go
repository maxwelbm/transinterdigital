package usecases

import "errors"

type TransferInput struct {
	AccountOriginID      int64   `json:"account_origin_id"`
	AccountDestinationID int64   `json:"account_destination_id"`
	Amount               float64 `json:"amount"`
}

func (c *useCase) TransferAccountToAnother(input TransferInput) error {
	balanceOrigin, err := c.repository.account.Balance(int(input.AccountOriginID))
	if err != nil {
		return err
	}

	if input.Amount > balanceOrigin {
		return errors.New("insufficient balance to complete the transfer")
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
