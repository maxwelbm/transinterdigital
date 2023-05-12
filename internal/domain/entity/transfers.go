package entity

type Transfers struct {
	ID                   int64
	AccountOriginID      int64
	AccountDestinationID int64
	Amount               float64
	CreatedAt            string
}

type TransfersRepository interface {
	List(originID int) ([]Transfers, error)
	Save(transfers Transfers) error
}
