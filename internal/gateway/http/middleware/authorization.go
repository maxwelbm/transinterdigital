package middleware

import (
	"github.com/maxwelbm/transinterdigital/pkg/helper"
	"github.com/maxwelbm/transinterdigital/pkg/token"
	"net/http"
)

func Authenticate(nextFunc http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := token.ValidToken(r); err != nil {
			helper.RespError(w, http.StatusUnauthorized, err.Error())
			return
		}
		nextFunc(w, r)
	}
}
