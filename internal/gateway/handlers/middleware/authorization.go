package middleware

import (
	"context"
	"github.com/maxwelbm/transinterdigital/internal/config"
	"github.com/maxwelbm/transinterdigital/pkg/helper"
	"github.com/maxwelbm/transinterdigital/pkg/token"
	"net/http"
)

func Authenticate(next http.HandlerFunc, conf config.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		originID, err := token.ValidToken(r, conf.KeySecret)
		if err != nil {
			helper.RespError(w, http.StatusUnauthorized, err.Error())
			return
		}
		ctx := context.WithValue(r.Context(), "origin_id", originID)
		r = r.WithContext(ctx)
		next(w, r)
	}
}
