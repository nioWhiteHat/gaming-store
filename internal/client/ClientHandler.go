package client

import (
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/nioWhiteHat/gaming-store-backend.git/internal/data"
)

type Client struct {
	DB *pgxpool.Pool
	Cart data.CartModel
}

