package services

import (
	"fmt"

	"github.com/jip/portfolio-backend/internal/entity"
	"github.com/jmoiron/sqlx"
	_ "github.com/jackc/pgx/v5/stdlib"
)

func NewPostgresClient(cfg *entity.Config) (*sqlx.DB, error) {
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

	return db, nil
}
