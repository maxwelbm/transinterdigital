package usecases

import (
	"errors"
	"github.com/jackc/pgx/v5"
	"github.com/maxwelbm/transinterdigital/pkg/cpf"
	"github.com/maxwelbm/transinterdigital/pkg/helper"
	"github.com/maxwelbm/transinterdigital/pkg/token"
	"net/http"
)

type TokenInput struct {
	CPF       string
	Secret    string
	KeySecret string
}

type Token struct {
	Token string
}

func (c *useCase) LoginGetToken(input TokenInput) (Token, *helper.Response) {
	if !cpf.Validate(input.CPF) {
		return Token{}, &helper.Response{Status: http.StatusBadRequest, Err: errors.New("cpf invalid")}
	}

	accountID, err := c.repository.account.GetAccountID(input.CPF, input.Secret)
	if err != nil {
		if errors.Is(pgx.ErrNoRows, err) {
			return Token{}, &helper.Response{Status: http.StatusNotFound, Err: err}
		}
		return Token{}, &helper.Response{Status: http.StatusInternalServerError, Err: errors.New("failed in get item")}
	}

	t, err := token.GenToken(accountID, input.KeySecret)
	if err != nil {
		return Token{}, &helper.Response{Status: http.StatusInternalServerError, Err: err}
	}

	return Token{t}, nil
}
