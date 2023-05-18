package handlers

import (
	"encoding/json"
	"github.com/maxwelbm/transinterdigital/internal/domain/usecases"
	"github.com/maxwelbm/transinterdigital/pkg/helper"
	"net/http"
)

type TransferInput struct {
	AccountOriginID      int64   `json:"origin_id"`
	AccountDestinationID int64   `json:"destination_id"`
	Amount               float64 `json:"amount"`
}

func (h Handlers) TransferAccountToAnother(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")

	var transfer TransferInput
	err := json.NewDecoder(r.Body).Decode(&transfer)
	if err != nil {
		helper.RespError(w, http.StatusBadRequest, "failed to serialize input struct body")
		return
	}

	originID := r.Context().Value("origin_id")

	transferInput := usecases.TransferInput{
		AccountOriginID:      originID.(int64),
		AccountDestinationID: transfer.AccountDestinationID,
		Amount:               transfer.Amount,
	}

	errResp := h.UseCase.TransferAccountToAnother(transferInput)
	if errResp != nil {
		helper.RespError(w, errResp.Status, errResp.Error())
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
