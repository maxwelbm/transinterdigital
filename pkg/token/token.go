package token

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
	"time"
)

func ValidToken(r *http.Request, keySecret string) (int64, error) {
	tokenString := r.Header.Get("Authorization")
	return CheckToken(tokenString, keySecret)
}

func CheckToken(tokenString, keySecret string) (int64, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("invalid method signature")
		}
		return []byte(keySecret), nil
	})

	if err != nil {
		return 0, fmt.Errorf("invalid in valid token: %v", err)
	}

	var originID int64 = 0
	if token.Valid {
		claims := token.Claims.(jwt.MapClaims)
		exp := claims["exp"].(float64)
		originID = int64(claims["origin_id"].(float64))

		expTime := time.Unix(int64(exp), 0)
		if expTime.Before(time.Now()) {
			return 0, fmt.Errorf("token JWT expiration")
		}
		return originID, nil
	}

	return originID, fmt.Errorf("token JWT invalid")
}

func GenToken(accountID int64, keySecret string) (string, error) {
	tk := jwt.New(jwt.SigningMethodHS256)
	claims := tk.Claims.(jwt.MapClaims)
	claims["origin_id"] = accountID
	claims["exp"] = time.Now().Add((time.Hour * 24) * 30).Unix()
	token, err := tk.SignedString([]byte(keySecret))
	if err != nil {
		return "", err
	}
	return token, nil
}
