package usecases

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"os"
	"time"
)

type TokenInput struct {
	CPF    string `json:"cpf"`
	Secret string `json:"secret"`
}

type Token struct {
	Token string
}

func (c *useCase) LoginGetToken(input TokenInput) (Token, error) {
	keySecret := os.Getenv("KEY_SECRET")

	tk := jwt.New(jwt.SigningMethodHS256)
	claims := tk.Claims.(jwt.MapClaims)
	claims["cpf"] = input.CPF
	claims["secret"] = input.Secret
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	token, err := tk.SignedString([]byte(keySecret))
	if err != nil {
		return Token{}, errors.New("error signing token")
	}

	return Token{token}, nil
}
