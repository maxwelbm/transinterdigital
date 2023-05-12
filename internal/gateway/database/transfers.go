package database

import (
	"context"
	"errors"
	"github.com/jackc/pgx/v5"
	"github.com/maxwelbm/transinterdigital/internal/domain/entity"
	"github.com/maxwelbm/transinterdigital/pkg/logger"
)

type TransfersRepository struct {
	Db *pgx.Conn
}

func NewTransfersRepository(db *pgx.Conn) *TransfersRepository {
	return &TransfersRepository{Db: db}
}

func (d *TransfersRepository) List(originID int) ([]entity.Transfers, error) {
	query := `
		SELECT id, account_origin_id, account_destination_id, amount, created_at 
		FROM transfers 
		where WHERE id = $1
	`
	rows, err := d.Db.Query(context.Background(), query, originID)

	if err != nil {
		return []entity.Transfers{}, errors.New("failed execute the consult: " + err.Error())
	}
	defer rows.Close()

	var transfers []entity.Transfers

	for rows.Next() {
		var transfer entity.Transfers
		err = rows.Scan(&transfer.ID, &transfer.AccountOriginID, &transfer.AccountDestinationID, &transfer.Amount, &transfer.CreatedAt)
		if err != nil {
			return []entity.Transfers{}, errors.New("failed scan the line: " + err.Error())
		}
		transfers = append(transfers, transfer)
	}

	if err = rows.Err(); err != nil {
		return []entity.Transfers{}, errors.New("error entry iteration the results: " + err.Error())
	}

	return transfers, nil
}

func (d *TransfersRepository) Save(transfer entity.Transfers) error {
	query := `
		INSERT INTO account (account_origin_id, account_destination_id, amount, created_at)
		VALUES ($1, $2, $3, $4)
	`
	_, err := d.Db.Exec(context.Background(), query,
		transfer.AccountOriginID, transfer.AccountDestinationID, transfer.Amount, transfer.CreatedAt)
	if err != nil {
		return errors.New("failed insert transfer from table: " + err.Error())
	}
	logger.Info("create new transfer")
	return nil
}
