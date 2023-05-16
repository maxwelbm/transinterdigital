package usecases

import (
	"errors"
	"github.com/maxwelbm/transinterdigital/internal/domain/entity"
	"github.com/maxwelbm/transinterdigital/pkg/cpf"
)

type AccountInput struct {
	Name    string
	CPF     string
	Secret  string
	Balance float64
}

func (c *useCase) CreateAccount(input AccountInput) error {

	if !cpf.Validate(input.CPF) {
		return errors.New("cpf invalid")
	}

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
