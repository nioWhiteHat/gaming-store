package database

import (
	"context"
	"log"
	"github.com/jackc/pgx/v5/pgxpool"
)

func NewDBPool(Dsn string) (*pgxpool.Pool, error) {


	log.Printf("DEBUG: INSIDE NewDBPool, received DSN: [%s]", Dsn)

	pool, err := pgxpool.New(context.Background(), Dsn)
	if err != nil {
		
		return nil, err
	}

	return pool, nil
}