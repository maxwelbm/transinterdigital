package usecases

type AccountOutput struct {
	ID        int64   `json:"id"`
	Name      string  `json:"name"`
	CPF       string  `json:"cpf"`
	Secret    string  `json:"secret"`
	Balance   float64 `json:"balance"`
	CreatedAt string  `json:"created_at"`
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
