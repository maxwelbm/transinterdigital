package midleware

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

func Authorizition() {
	// Defina o token JWT a ser validado
	tokenString := "seu-token-jwt-aqui"

	// Defina a chave secreta usada para assinar o token
	chaveSecreta := "sua-chave-secreta"

	// Parseie o token JWT
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Verifique se o método de assinatura é HMAC-SHA256
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Método de assinatura inválido: %v", token.Header["alg"])
		}

		// Retorne a chave secreta usada para assinar o token
		return []byte(chaveSecreta), nil
	})

	// Verifique se ocorreu algum erro na validação do token
	if err != nil {
		fmt.Println("Erro ao validar o token:", err)
		return
	}

	// Verifique se o token é válido
	if token.Valid {
		fmt.Println("Token JWT válido")

		// Acesse as reivindicações (claims) do token
		claims := token.Claims.(jwt.MapClaims)
		cpf := claims["cpf"].(string)
		senha := claims["senha"].(string)
		exp := claims["exp"].(float64)

		// Verifique a data de expiração do token
		expTime := time.Unix(int64(exp), 0)
		if expTime.Before(time.Now()) {
			fmt.Println("Token JWT expirado")
		} else {
			fmt.Println("CPF:", cpf)
			fmt.Println("Senha:", senha)
		}
	} else {
		fmt.Println("Token JWT inválido")
	}
}
