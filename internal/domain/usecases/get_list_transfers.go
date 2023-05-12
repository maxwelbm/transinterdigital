package usecases

type TransfersOutput struct {
	ID                   int64   `json:"id"`
	AccountOriginID      int64   `json:"account_origin_id"`
	AccountDestinationID int64   `json:"account_destination_id"`
	Amount               float64 `json:"amount"`
	CreatedAt            string  `json:"created_at"`
}

func (c *useCase) GetListTransfers(originID int) ([]TransfersOutput, error) {
	transfers, err := c.repository.transfer.List(originID)
	if err != nil {
		return []TransfersOutput{}, err
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
