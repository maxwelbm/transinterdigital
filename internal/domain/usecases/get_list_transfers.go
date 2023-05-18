package usecases

import (
	"github.com/maxwelbm/transinterdigital/pkg/helper"
	"net/http"
	"time"
)

type TransfersOutput struct {
	ID                   int64
	AccountOriginID      int64
	AccountDestinationID int64
	Amount               float64
	CreatedAt            time.Time
}

func (c *useCase) GetListTransfers(originID int64) ([]TransfersOutput, *helper.Response) {
	transfers, err := c.repository.transfer.List(originID)
	if err != nil {
		return []TransfersOutput{}, &helper.Response{Status: http.StatusInternalServerError, Err: err}
	}

	transfersOutput := []TransfersOutput{}
	for _, v := range transfers {
		var transfer TransfersOutput
		transfer.ID = v.ID
		transfer.AccountOriginID = v.AccountOriginID
		transfer.AccountDestinationID = v.AccountDestinationID
		transfer.Amount = v.Amount
		transfer.CreatedAt = v.CreatedAt
		transfersOutput = append(transfersOutput, transfer)
	}

	return transfersOutput, nil
}
