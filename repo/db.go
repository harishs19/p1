package repository

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

type DB struct {
	*pgxpool.Pool
}

func NewDB(ctx context.Context) (*DB, error) {
	dsn := fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s sslmode=disable pool_max_conns=20",
		os.Getenv("DB_USERNAME"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_DATABASE"),
	)

config, err := pgxpool.ParseConfig(dsn)
if err != nil {
	return nil, err
}
db, err := pgxpool.New(ctx, config.ConnString())
if err != nil {
	return nil, err
}

err = db.Ping(ctx)
if err != nil {
	return nil, err
}

return &DB{
	db,
}, nil

}
// Close closes the database connection
func (db *DB) Close() {
db.Pool.Close()
}

