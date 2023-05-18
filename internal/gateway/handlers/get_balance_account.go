package handlers

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"github.com/maxwelbm/transinterdigital/pkg/helper"
	"net/http"
	"strconv"
)

func (h Handlers) GetBalanceAccount(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")

	paramAccountID := chi.URLParam(r, "account_id")
	accountID, err := strconv.Atoi(paramAccountID)
	if err != nil {
		helper.RespError(w, http.StatusBadRequest, "failed parse type string to integer "+err.Error())
		return
	}

	accounts, errResp := h.UseCase.GetBalance(accountID)
	if errResp != nil {
		helper.RespError(w, errResp.Status, errResp.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(accounts)
}
