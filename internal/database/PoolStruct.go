package database

import "github.com/jackc/pgx/v5/pgxpool"

type Pooldb struct {
	DbPool *pgxpool.Pool
}