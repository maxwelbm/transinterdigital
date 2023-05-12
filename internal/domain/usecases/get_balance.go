package usecases

type BalanceOutput struct {
	Balance float64 `json:"balance"`
}

func (c *useCase) GetBalance(accountID int) (BalanceOutput, error) {
	balance, err := c.repository.account.Balance(accountID)
	if err != nil {
		return BalanceOutput{0}, err
	}
	return BalanceOutput{balance}, nil
}
