package usecases

import (
	"errors"
	"github.com/maxwelbm/transinterdigital/internal/domain/entity"
	"github.com/maxwelbm/transinterdigital/pkg/cpf"
	"github.com/maxwelbm/transinterdigital/pkg/helper"
	"net/http"
)

type AccountInput struct {
	Name    string
	CPF     string
	Secret  string
	Balance float64
}

func (c *useCase) CreateAccount(input AccountInput) *helper.Response {

	if !cpf.Validate(input.CPF) {
		return &helper.Response{Status: http.StatusBadRequest, Err: errors.New("cpf invalid")}
	}

	account := entity.Account{
		Name:    input.Name,
		CPF:     input.CPF,
		Secret:  input.Secret,
		Balance: input.Balance,
	}

	if err := c.repository.account.Save(&account); err != nil {
		return &helper.Response{Status: http.StatusInternalServerError, Err: errors.New("failed in save account")}
	}

	return nil
}
