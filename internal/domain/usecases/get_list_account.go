package usecases

import "time"

type AccountOutput struct {
	ID        int64
	Name      string
	CPF       string
	Secret    string
	Balance   float64
	CreatedAt time.Time
}

func (c *useCase) GetListAccount() ([]AccountOutput, error) {
	accounts, err := c.repository.account.List()
	if err != nil {
		return []AccountOutput{}, err
	}

	accountsOutput := []AccountOutput{}
	for _, v := range accounts {
		var account AccountOutput
		account.ID = v.ID
		account.Name = v.Name
		account.CPF = v.CPF
		account.Secret = v.Secret
		account.Balance = v.Balance
		account.CreatedAt = v.CreatedAt
		accountsOutput = append(accountsOutput, account)
	}

	return accountsOutput, nil
}
