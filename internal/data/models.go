package data

import (
	"github.com/jackc/pgx/v5/pgxpool"

)

type Models struct {
	ClientModel        ClientModel
	GameModel          GameModel
	VendorModel        VendorModel
	SharedModelMethods SharedModelMethods
	GenresModel        GenresModel
	SessionsModel	   SessionsModel
	MetadataModel	   MetadataModel
	CartModel		   CartModel
}

func NewModlel(db *pgxpool.Pool) Models {
	return Models{
		GameModel:          GameModel{DB: db},
		ClientModel:        ClientModel{DB: db},
		VendorModel:        VendorModel{DB: db},
		SharedModelMethods: SharedModelMethods{DB: db},
		GenresModel:        GenresModel{DB: db},
		SessionsModel:	    SessionsModel{DB:db},
		MetadataModel: MetadataModel{DB: db},
		CartModel:     CartModel{DB: db},
		
		
	}
}
type CartModel struct{
	DB *pgxpool.Pool
}
type SessionsModel struct{
	DB *pgxpool.Pool
}
type GenresModel struct {
	DB *pgxpool.Pool
}

type SharedModelMethods struct {
	DB *pgxpool.Pool
}
type TableFiller struct{
	DB *pgxpool.Pool
}


type ClientModel struct {
	DB *pgxpool.Pool
}
type GameModel struct {
	DB *pgxpool.Pool
}
type VendorModel struct {
	DB *pgxpool.Pool
}
