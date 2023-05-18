package entity

type TransfersMock struct{}

func (t TransfersMock) List(originID int64) ([]Transfers, error) {
	if originID == 1 {
		return []Transfers{{ID: 1, AccountDestinationID: 2, AccountOriginID: 1, Amount: 2}}, nil
	}
	return nil, nil
}

func (t TransfersMock) Save(transfers Transfers) error {
	return nil
}
