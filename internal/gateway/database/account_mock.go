package database

import (
	"context"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

type DbMock struct{}

func (d DbMock) Exec(ctx context.Context, sql string, arguments ...any) (pgconn.CommandTag, error) {
	return pgconn.CommandTag{}, nil
}

func (d DbMock) QueryRow(ctx context.Context, sql string, args ...any) pgx.Row {
	return nil
}

func (d DbMock) Query(ctx context.Context, sql string, args ...any) (pgx.Rows, error) {
	return nil, nil
}
