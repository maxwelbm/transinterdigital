package handlers

import (
	"encoding/json"
	"github.com/maxwelbm/transinterdigital/pkg/helper"
	"net/http"
)

func (h Handlers) GetListAccount(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")

	accounts, err := h.UseCase.GetListAccount()
	if err != nil {
		helper.RespError(w, http.StatusInternalServerError, "failed to get account list ")
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(accounts)
}
