package middleware

import (
	"github.com/maxwelbm/transinterdigital/pkg/helper"
	"github.com/maxwelbm/transinterdigital/pkg/token"
	"net/http"
	"strconv"
)

func Authenticate(next http.HandlerFunc) http.HandlerFunc {
	var originID int64 = 0
	var err error = nil
	return func(w http.ResponseWriter, r *http.Request) {
		if originID, err = token.ValidToken(r); err != nil {
			helper.RespError(w, http.StatusUnauthorized, err.Error())
			return
		}
		r.Header.Set("origin_id", strconv.FormatInt(originID, 10))
		// TODO: set context
		next(w, r)
	}
}
