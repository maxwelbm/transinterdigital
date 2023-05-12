package http

import (
	"encoding/json"
	"github.com/maxwelbm/transinterdigital/internal/domain/usecases"
	"github.com/maxwelbm/transinterdigital/pkg/helper"
	"net/http"
)

func (h Handler) TransferAccountToAnother(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")

	var transfer usecases.TransferInput
	err := json.NewDecoder(r.Body).Decode(&transfer)
	if err != nil {
		helper.RespError(w, http.StatusBadRequest, "failed to serialize input struct body "+err.Error())
		return
	}

	err = h.UseCase.TransferAccountToAnother(transfer)
	if err != nil {
		helper.RespError(w, http.StatusInternalServerError, "account creation failed "+err.Error())
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
