package database

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/maxwelbm/transinterdigital/internal/domain/entity"
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
	Begin(ctx context.Context) (pgx.Tx, error)
}

func (d *AccountRepository) Save(account *entity.Account) error {
	query := `
		INSERT INTO account (name, cpf, secret, balance, created_at)
		VALUES ($1, $2, $3, $4, $5)
	`
	_, err := d.Db.Exec(context.Background(), query,
		account.Name, account.CPF, account.Secret, account.Balance, account.CreatedAt)
	if err != nil {
		return err
	}
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
			return 0, err
		}
		return 0, err
	}

	return account.Balance, nil
}

func (d *AccountRepository) GetAccountID(cpf, secret string) (int64, error) {
	query := `
		SELECT id, name, cpf, secret, balance, created_at
		FROM account
		WHERE cpf = $1 AND secret = $2
	`
	row := d.Db.QueryRow(context.Background(), query, cpf, secret)

	var account entity.Account
	err := row.Scan(&account.ID, &account.Name, &account.CPF, &account.Secret, &account.Balance, &account.CreatedAt)
	if err != nil {
		if err == pgx.ErrNoRows {
			return 0, err
		}
		return 0, err
	}

	return account.ID, nil
}

func (d *AccountRepository) List() ([]entity.Account, error) {
	query := `select id, name, cpf, secret, balance, created_at from account`
	rows, err := d.Db.Query(context.Background(), query)

	if err != nil {
		return []entity.Account{}, err
	}
	defer rows.Close()

	var accounts []entity.Account

	for rows.Next() {
		var account entity.Account
		err = rows.Scan(&account.ID, &account.Name, &account.CPF, &account.Secret, &account.Balance, &account.CreatedAt)
		if err != nil {
			fmt.Println("list: ", err)
			return []entity.Account{}, err
		}
		accounts = append(accounts, account)
	}

	if err = rows.Err(); err != nil {
		return []entity.Account{}, err
	}

	return accounts, nil
}

func (d *AccountRepository) UpdateBalance(accountID int, balance float64) error {
	query := `UPDATE public.account SET balance = $1 WHERE id = $2`
	_, err := d.Db.Exec(context.Background(), query, balance, accountID)
	if err != nil {
		return err
	}
	return nil
}

func (d *AccountRepository) TransferAccountToAnother(originID, destinationID int, balance float64) error {
	tx, err := d.Db.Begin(context.Background())
	if err != nil {
		return err
	}

	_, err = tx.Exec(context.Background(), "UPDATE public.account SET balance = $1 WHERE id = $2", balance, originID)
	if err != nil {
		if err = tx.Rollback(context.Background()); err != nil {
			return err
		}
		return err
	}

	return nil
}
