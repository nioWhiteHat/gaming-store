package data

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type MetadataModel struct {
	DB *pgxpool.Pool

}

func (m *MetadataModel) GetGamesMetadata(ctx context.Context, genre string, priceMin int, priceMax int, platform string)(error,int,int){
	query := `
	SELECT
    COUNT(DISTINCT games.id) as total_games
	FROM
		games
	JOIN
		game_genres ON game_genres.game_id = games.id
	JOIN
		genre ON genre.id = game_genres.genre_id
	JOIN
		keys ON keys.game_id = games.id
	JOIN
		platforms as plat on plat.id = games.platform_id
	WHERE
    plat.name = $4
    AND genre.name = $1
    AND keys.price >= $3
    AND keys.price <= $2`
	var count int 
	err := m.DB.QueryRow(ctx, query, genre, priceMax,priceMin,platform).Scan(&count)
	if err!=nil{
		return err,0,0
	}
	lastPage := count%10
	pageCount := count/10
	if lastPage !=0{
		pageCount = pageCount+1
	}
	
	
	return nil,pageCount,lastPage

}


func (m *MetadataModel) CartMetadata(){

}


