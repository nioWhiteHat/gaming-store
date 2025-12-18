package public

import (
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/nioWhiteHat/gaming-store-backend.git/internal/data"
)

type Public struct {
	Db *pgxpool.Pool
	Sm data.SharedModelMethods
	Ses data.SessionsModel
}

func (p *Public) ModelsInit(){
	p.Sm = data.SharedModelMethods{DB: p.Db}
	p.Ses = data.SessionsModel{DB: p.Db}
}
