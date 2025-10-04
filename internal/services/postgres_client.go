package services

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/jip/portfolio-backend/internal/entity"
	"github.com/jmoiron/sqlx"
	_ "github.com/jackc/pgx/v5/stdlib"
)

type txKey struct{}

type DBTX interface {
	sqlx.ExtContext
	BindNamed(query string, arg interface{}) (string, []interface{}, error)
	GetContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error
	SelectContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error
	Rebind(query string) string
	QueryRowContext(ctx context.Context, query string, args ...interface{}) *sql.Row
}

type PostgresClient struct {
	DB *sqlx.DB
}

func NewPostgresClient(cfg *entity.Config) (*PostgresClient, error) {
	host := cfg.Postgres.Host
	port := cfg.Postgres.Port
	user := cfg.Postgres.User
	password := cfg.Postgres.Password
	dbname := cfg.Postgres.DBName
	sslmode := cfg.Postgres.SSLMode

	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s", host, port, user, password, dbname, sslmode)

	db, err := sqlx.Connect("pgx", dsn)
	if err != nil {
		return nil, err
	}

	return &PostgresClient{DB: db}, nil
}

func (pc *PostgresClient) StartProcess(ctx context.Context) (context.Context, error) {
	tx, err := pc.DB.BeginTxx(ctx, nil)
	if err != nil {
		return nil, err
	}
	return context.WithValue(ctx, txKey{}, tx), nil
}

func (pc *PostgresClient) CloseProcess(ctx context.Context, err error) error {
	tx, ok := ctx.Value(txKey{}).(*sqlx.Tx)
	if !ok {
		return fmt.Errorf("transaction not found in context")
	}

	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("tx err: %v, rb err: %v", err, rbErr)
		}
		return err
	}
	return tx.Commit()
}

func (pc *PostgresClient) GetExecutor(ctx context.Context) DBTX {
	if tx, ok := ctx.Value(txKey{}).(*sqlx.Tx); ok {
		return tx
	}
	return pc.DB
}