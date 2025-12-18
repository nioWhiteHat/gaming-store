package genres

import (
	"net/http"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/nioWhiteHat/gaming-store-backend.git/internal/data"
	"github.com/nioWhiteHat/gaming-store-backend.git/internal/public"
)

type Genre struct {
	Db                 *pgxpool.Pool
	GenresModel          data.GenresModel
	SharedModelMethods data.SharedModelMethods
	MetaDataModel      data.MetadataModel
}

func (g *Genre) ModelsInit() {
	g.SharedModelMethods = data.SharedModelMethods{DB: g.Db}
	g.GenresModel = data.GenresModel{DB: g.Db}
	g.MetaDataModel = data.MetadataModel{DB: g.Db}
}

func (g *Genre) GetGenres(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173")
    w.Header().Set("Content-Type", "application/json")
	ctx := r.Context()
	genres,err := g.GenresModel.GetGenres(ctx)
	if err!=nil{
		http.Error(w,err.Error(),http.StatusInternalServerError)
		return
	}
	public.SendJSONResponse(w,genres)


}