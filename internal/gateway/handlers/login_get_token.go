package api

import (
	"encoding/json"
	"github.com/maxwelbm/transinterdigital/internal/domain/usecases"
	"github.com/maxwelbm/transinterdigital/pkg/helper"
	"net/http"
)

func (h Handler) LoginGetToken(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")

	var customer usecases.TokenInput
	err := json.NewDecoder(r.Body).Decode(&customer)
	if err != nil {
		helper.RespError(w, http.StatusBadRequest, "failed to serialize input struct body "+err.Error())
		return
	}

	token, err := h.UseCase.LoginGetToken(customer)
	if err != nil {
		helper.RespError(w, http.StatusInternalServerError, "token generation failed "+err.Error())
		return
	}

	w.WriteHeader(http.StatusNoContent)
	json.NewEncoder(w).Encode(token)
}
