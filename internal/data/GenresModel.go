package data

import (
	"context"
	"log"
)

func (g *GenresModel) FillGenreTable(ctx context.Context, slugs []string, names []string) {
	BaseQuery := "INSERT INTO genre(slug,name)VALUES($1,$2)"
	for i,_ := range slugs{
		slug:=slugs[i]
		name:=names[i]
		log.Printf("inserting %s, %s", slug, name)
		_,err := g.DB.Exec(ctx, BaseQuery, slug, name)
		
		if err!=nil{
			log.Fatalf("couldnt insert values %s, %s", slug, name)
		}
		log.Println("insert success")

	}

}

func (g *GenresModel) GetGenres(ctx context.Context) ([]Genre, error) {
    query := `select id, name from genre`
    
    // 1. execution
    rows, err := g.DB.Query(ctx, query)
    if err != nil {
        return nil, err 
    }
    defer rows.Close() 

    
    var genres []Genre
	for rows.Next(){
		var genre Genre
		err := rows.Scan(&genre.Id,&genre.Name)
		if err!=nil{
			return nil, err 
		}
		genres = append(genres, genre)
	}

   
    return genres, nil
}
