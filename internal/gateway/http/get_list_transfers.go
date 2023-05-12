package http

import (
	"encoding/json"
	"github.com/maxwelbm/transinterdigital/pkg/helper"
	"net/http"
)

func (h Handler) GetListTransfers(w http.ResponseWriter, r *http.Request) {
	transfers, err := h.UseCase.GetListTransfers(0)
	if err != nil {
		helper.RespError(w, http.StatusInternalServerError, "failed to get transfer list "+err.Error())
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(transfers)
}
