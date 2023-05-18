package handlers

import (
	"encoding/json"
	"github.com/maxwelbm/transinterdigital/pkg/helper"
	"net/http"
)

func (h Handlers) GetListTransfers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")

	originID := r.Context().Value("origin_id")

	transfers, errResp := h.UseCase.GetListTransfers(originID.(int64))
	if errResp != nil {
		helper.RespError(w, errResp.Status, errResp.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(transfers)
}
