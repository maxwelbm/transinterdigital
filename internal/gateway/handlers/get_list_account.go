package handlers

import (
	"encoding/json"
	"github.com/maxwelbm/transinterdigital/pkg/helper"
	"net/http"
)

func (h Handlers) GetListAccount(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")

	accounts, errResp := h.UseCase.GetListAccount()
	if errResp != nil {
		helper.RespError(w, errResp.Status, errResp.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(accounts)
}
