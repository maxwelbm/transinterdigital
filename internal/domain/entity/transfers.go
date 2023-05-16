package entity

import "time"

type Transfers struct {
	ID                   int64
	AccountOriginID      int64
	AccountDestinationID int64
	Amount               float64
	CreatedAt            time.Time
}

type TransfersRepository interface {
	List(originID int64) ([]Transfers, error)
	Save(transfers Transfers) error
}
