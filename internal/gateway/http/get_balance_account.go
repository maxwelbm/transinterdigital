package http

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"github.com/maxwelbm/transinterdigital/pkg/helper"
	"net/http"
	"strconv"
)

func (h Handler) GetBalanceAccount(w http.ResponseWriter, r *http.Request) {
	paramAccountID := chi.URLParam(r, "account_id")
	accountID, err := strconv.Atoi(paramAccountID)
	if err != nil {
		helper.RespError(w, http.StatusInternalServerError, "failed parse type string to integer "+err.Error())
		return
	}

	accounts, err := h.UseCase.GetBalance(accountID)
	if err != nil {
		helper.RespError(w, http.StatusInternalServerError, "failed to get account list "+err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(accounts)
}
