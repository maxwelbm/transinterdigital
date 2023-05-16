package api

import (
	"encoding/json"
	"github.com/maxwelbm/transinterdigital/pkg/helper"
	"net/http"
	"strconv"
)

func (h Handler) GetListTransfers(w http.ResponseWriter, r *http.Request) {
	originIDString := r.Header.Get("origin_id")
	originID, err := strconv.Atoi(originIDString)
	if err != nil {
		helper.RespError(w, http.StatusBadRequest, "failed in parse origin id")
		return
	}

	transfers, err := h.UseCase.GetListTransfers(originID)
	if err != nil {
		helper.RespError(w, http.StatusInternalServerError, "failed to get transfer list ")
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(transfers)
}
