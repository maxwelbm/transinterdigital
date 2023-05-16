package handlers

import (
	"encoding/json"
	"github.com/maxwelbm/transinterdigital/internal/domain/usecases"
	"github.com/maxwelbm/transinterdigital/pkg/helper"
	"net/http"
)

type TokenInput struct {
	CPF    string `json:"cpf"`
	Secret string `json:"secret"`
}

func (h Handlers) LoginGetToken(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")

	var customer TokenInput
	err := json.NewDecoder(r.Body).Decode(&customer)
	if err != nil {
		helper.RespError(w, http.StatusBadRequest, "failed to serialize input struct body")
		return
	}

	customerUseCase := usecases.TokenInput{
		CPF:       customer.CPF,
		Secret:    customer.Secret,
		KeySecret: h.Config.KeySecret,
	}

	token, err := h.UseCase.LoginGetToken(customerUseCase)
	if err != nil {
		helper.RespError(w, http.StatusInternalServerError, "token generation failed")
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(token)
}
