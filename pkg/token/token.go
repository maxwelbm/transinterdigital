package token

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
	"os"
	"time"
)

func ValidToken(r *http.Request) error {
	tokenString := r.Header.Get("Authorization")

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("invalid method signature")
		}
		return []byte(os.Getenv("KEY_SECRET")), nil
	})

	if err != nil {
		return fmt.Errorf("invalid in valid token: %v", err)
	}

	if token.Valid {
		claims := token.Claims.(jwt.MapClaims)
		exp := claims["exp"].(float64)

		expTime := time.Unix(int64(exp), 0)
		if expTime.Before(time.Now()) {
			return fmt.Errorf("token JWT expiration")
		}
		return nil
	}

	return fmt.Errorf("token JWT invalid")
}
