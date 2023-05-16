package api

import (
	"encoding/json"
	"github.com/maxwelbm/transinterdigital/internal/domain/usecases"
	"github.com/maxwelbm/transinterdigital/pkg/helper"
	"net/http"
)

func (h Handler) CreateAccount(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	var account usecases.AccountInput
	err := json.NewDecoder(r.Body).Decode(&account)
	if err != nil {
		helper.RespError(w, http.StatusBadRequest, "failed to serialize input struct body")
		return
	}
	err = h.UseCase.CreateAccount(account)
	if err != nil {
		helper.RespError(w, http.StatusInternalServerError, "account creation failed ")
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
