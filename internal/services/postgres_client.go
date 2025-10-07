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
	NamedExecContext(ctx context.Context, query string, arg interface{}) (sql.Result, error)
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

type txWrapper struct {
	tx      *sqlx.Tx
	nesting int
}

func (pc *PostgresClient) StartProcess(ctx context.Context) (context.Context, error) {
	if wrapper, ok := ctx.Value(txKey{}).(*txWrapper); ok {
		wrapper.nesting++
		return ctx, nil
	}

	tx, err := pc.DB.BeginTxx(ctx, nil)
	if err != nil {
		return nil, err
	}

	return context.WithValue(ctx, txKey{}, &txWrapper{tx: tx, nesting: 1}), nil
}

func (pc *PostgresClient) CloseProcess(ctx context.Context, err error) error {
	wrapper, ok := ctx.Value(txKey{}).(*txWrapper)
	if !ok {
		return fmt.Errorf("transaction not found in context")
	}

	wrapper.nesting--
	if wrapper.nesting > 0 {
		return nil
	}

	tx := wrapper.tx
	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("tx err: %v, rb err: %v", err, rbErr)
		}
		return err
	}
	return tx.Commit()
}

func (pc *PostgresClient) GetExecutor(ctx context.Context) DBTX {
	if wrapper, ok := ctx.Value(txKey{}).(*txWrapper); ok {
		return wrapper.tx
	}
	return pc.DB
}