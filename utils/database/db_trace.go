package database

import (
	"context"

	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
)

type DBTrace struct {
	DB Ext
}

var _ Ext = (*DBTrace)(nil)

func (rcv *DBTrace) Query(ctx context.Context, query string, args ...interface{}) (pgx.Rows, error) {
	return rcv.DB.Query(ctx, query, args...)
}

func (rcv *DBTrace) QueryRow(ctx context.Context, query string, args ...interface{}) pgx.Row {
	return rcv.DB.QueryRow(ctx, query, args...)
}

func (rcv *DBTrace) Exec(ctx context.Context, sql string, args ...interface{}) (pgconn.CommandTag, error) {
	return rcv.DB.Exec(ctx, sql, args...)
}

func (rcv *DBTrace) Begin(ctx context.Context) (pgx.Tx, error) {
	tx, err := rcv.DB.Begin(ctx)
	return &TxTrace{
		Tx: tx,
	}, err
}

func (rcv *DBTrace) SendBatch(ctx context.Context, b *pgx.Batch) pgx.BatchResults {
	return rcv.DB.SendBatch(ctx, b)
}
