package usecases

import (
	"errors"
	"github.com/maxwelbm/transinterdigital/pkg/cpf"
	"github.com/maxwelbm/transinterdigital/pkg/token"
)

type TokenInput struct {
	CPF       string
	Secret    string
	KeySecret string
}

type Token struct {
	Token string
}

func (c *useCase) LoginGetToken(input TokenInput) (Token, error) {
	if !cpf.Validate(input.CPF) {
		return Token{}, errors.New("cpf invalid")
	}

	accountID, err := c.repository.account.GetAccountID(input.CPF, input.Secret)
	if err != nil {
		return Token{}, err
	}

	t, err := token.GenToken(accountID, input.KeySecret)
	if err != nil {
		return Token{}, err
	}

	return Token{t}, nil
}
