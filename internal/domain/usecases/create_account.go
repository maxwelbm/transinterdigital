package usecases

import "github.com/maxwelbm/transinterdigital/internal/domain/entity"

type AccountInput struct {
	Name    string  `json:"name"`
	CPF     string  `json:"cpf"`
	Secret  string  `json:"secret"`
	Balance float64 `json:"balance"`
}

func (c *useCase) CreateAccount(input AccountInput) error {
	account := entity.Account{
		Name:    input.Name,
		CPF:     input.CPF,
		Secret:  input.Secret,
		Balance: input.Balance,
	}

	if err := c.repository.account.Save(&account); err != nil {
		return err
	}

	return nil
}
