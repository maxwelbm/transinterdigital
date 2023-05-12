package database

import (
	"context"
	"errors"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/maxwelbm/transinterdigital/internal/domain/entity"
	"github.com/maxwelbm/transinterdigital/pkg/logger"
)

func NewAccountRepository(db *pgx.Conn) *AccountRepository {
	return &AccountRepository{Db: db}
}

type AccountRepository struct {
	Db Conn
}

type Conn interface {
	Exec(ctx context.Context, sql string, arguments ...any) (pgconn.CommandTag, error)
	QueryRow(ctx context.Context, sql string, args ...any) pgx.Row
	Query(ctx context.Context, sql string, args ...any) (pgx.Rows, error)
}

func (d *AccountRepository) Save(account *entity.Account) error {
	query := `
		INSERT INTO account (name, cpf, secret, balance, created_at)
		VALUES ($1, $2, $3, $4, $5)
	`
	_, err := d.Db.Exec(context.Background(), query,
		account.Name, account.CPF, account.Secret, account.Balance, account.CreatedAt)
	if err != nil {
		return errors.New("failed insert account from table: " + err.Error())
	}
	logger.Info("create new account")
	return nil
}

func (d *AccountRepository) Balance(accountID int) (float64, error) {
	query := `
		SELECT id, name, cpf, secret, balance, created_at
		FROM account
		WHERE id = $1
	`
	row := d.Db.QueryRow(context.Background(), query, accountID)

	var account entity.Account
	err := row.Scan(&account.ID, &account.Name, &account.CPF, &account.Secret, &account.Balance, &account.CreatedAt)
	if err != nil {
		if err == pgx.ErrNoRows {
			return 0, errors.New("none account find with id specified")
		}
		return 0, errors.New("failed execute the consult: " + err.Error())
	}

	return account.Balance, nil
}

func (d *AccountRepository) List() ([]entity.Account, error) {
	query := `SELECT id, name, cpf, secret, balance, created_at FROM account`
	rows, err := d.Db.Query(context.Background(), query)
	if err != nil {
		return []entity.Account{}, errors.New("failed execute the consult: " + err.Error())
	}
	defer rows.Close()

	var accounts []entity.Account

	for rows.Next() {
		var account entity.Account
		err = rows.Scan(&account.ID, &account.Name, &account.CPF, &account.Secret, &account.Balance, &account.CreatedAt)
		if err != nil {
			return []entity.Account{}, errors.New("failed scan the line: " + err.Error())
		}
		accounts = append(accounts, account)
	}

	if err = rows.Err(); err != nil {
		return []entity.Account{}, errors.New("error entry iteration the results: " + err.Error())
	}

	return accounts, nil
}

func (d *AccountRepository) UpdateBalance(accountID int, balance float64) error {
	query := `UPDATE account SET balance = $1, WHERE id = $2`
	_, err := d.Db.Exec(context.Background(),
		query, accountID, balance)
	if err != nil {
		return errors.New("Failed to perform update: " + err.Error())
	}
	return nil
}
