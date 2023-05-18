package handlers

import (
	"encoding/json"
	"github.com/maxwelbm/transinterdigital/internal/domain/usecases"
	"github.com/maxwelbm/transinterdigital/pkg/helper"
	"net/http"
)

type AccountInput struct {
	Name    string  `json:"name"`
	CPF     string  `json:"cpf"`
	Secret  string  `json:"secret"`
	Balance float64 `json:"balance"`
}

func (h Handlers) CreateAccount(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")

	var account AccountInput
	err := json.NewDecoder(r.Body).Decode(&account)
	if err != nil {
		helper.RespError(w, http.StatusBadRequest, "failed to serialize input struct body")
		return
	}

	accountUsecase := usecases.AccountInput{
		Name:    account.Name,
		CPF:     account.CPF,
		Secret:  account.Secret,
		Balance: account.Balance,
	}

	errResp := h.UseCase.CreateAccount(accountUsecase)
	if errResp != nil {
		helper.RespError(w, errResp.Status, errResp.Error())
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
