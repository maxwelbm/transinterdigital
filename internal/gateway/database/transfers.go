package database

import (
	"context"
	"github.com/jackc/pgx/v5"
	"github.com/maxwelbm/transinterdigital/internal/domain/entity"
)

type TransfersRepository struct {
	Db *pgx.Conn
}

func NewTransfersRepository(db *pgx.Conn) *TransfersRepository {
	return &TransfersRepository{Db: db}
}

func (d *TransfersRepository) List(originID int64) ([]entity.Transfers, error) {
	query := `
		select id, account_origin_id, account_destination_id, amount, created_at 
		from transfers 
		where account_origin_id = $1
	`
	rows, err := d.Db.Query(context.Background(), query, originID)

	if err != nil {
		return []entity.Transfers{}, err
	}
	defer rows.Close()

	var transfers []entity.Transfers

	for rows.Next() {
		var transfer entity.Transfers
		err = rows.Scan(&transfer.ID, &transfer.AccountOriginID, &transfer.AccountDestinationID, &transfer.Amount, &transfer.CreatedAt)
		if err != nil {
			return []entity.Transfers{}, err
		}
		transfers = append(transfers, transfer)
	}

	if err = rows.Err(); err != nil {
		return []entity.Transfers{}, err
	}

	return transfers, nil
}

func (d *TransfersRepository) Save(transfer entity.Transfers) error {
	query := `
		INSERT INTO transfers (account_origin_id, account_destination_id, amount, created_at)
		VALUES ($1, $2, $3, $4)
	`
	_, err := d.Db.Exec(context.Background(), query,
		transfer.AccountOriginID, transfer.AccountDestinationID, transfer.Amount, transfer.CreatedAt)
	if err != nil {
		return err
	}

	return nil
}
