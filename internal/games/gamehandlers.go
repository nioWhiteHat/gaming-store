package games

import (
	"log"
	"net/http"
	"strconv"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/nioWhiteHat/gaming-store-backend.git/internal/data"
	"github.com/nioWhiteHat/gaming-store-backend.git/internal/public"
)

type Game struct {
	Db 		  *pgxpool.Pool
	GameModel data.GameModel
	SharedModelMethods data.SharedModelMethods
	MetaDataModel	data.MetadataModel
}

func (g *Game) ModelsInit(){
	g.SharedModelMethods = data.SharedModelMethods{DB: g.Db}
	g.GameModel = data.GameModel{DB: g.Db}
    g.MetaDataModel = data.MetadataModel{DB:g.Db}
}


func (g *Game) GetGames(w http.ResponseWriter, r *http.Request) {
    // 1. CORS HEADERS (Must be first!)
    w.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173")
    w.Header().Set("Content-Type", "application/json")

    ctx := r.Context()

    // 2. USE PathValue, NOT Query()
    genre := r.PathValue("genre")
    pns := r.PathValue("page")
    platform := r.PathValue("platform")
    log.Printf("platform = %v", platform)

    // Parse PriceMin
    priceMin, err := strconv.Atoi(r.PathValue("priceMin"))
    log.Printf("min price = %v", priceMin)
    if err != nil {
        http.Error(w, "Invalid priceMin", http.StatusBadRequest)
        return 
    }

    // Parse PriceMax
    priceMax, err := strconv.Atoi(r.PathValue("priceMax"))
    log.Printf("max price = %v", priceMax)
    if err != nil {
        http.Error(w, "Invalid priceMax", http.StatusBadRequest)
        return // <--- Stop here if error
    }

    // Parse Page Number
    pn, err := strconv.Atoi(pns)
    if err != nil {
        http.Error(w, "Invalid page number", http.StatusBadRequest)
        return
    }
    offset := (pn - 1) * 10
    log.Printf("offset =  = %v", offset)
    log.Println("[DEBUG] before get games ")
    
    games, err := g.GameModel.GetGamesSql(ctx, genre, offset, priceMin, priceMax, platform)
    if err != nil {
    
        log.Println("Database Error:", err) 
        http.Error(w, "Database error", http.StatusInternalServerError)
        return
    }
    log.Printf("[DEBUG] Games retrieved successfully. Count: %d", len(games))

    log.Println("[DEBUG] Calling GetGamesMetadata...")
    
    err, pgCount, LastPage := g.MetaDataModel.GetGamesMetadata(ctx, genre, priceMin,priceMax,platform)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        
        return
    }

    // 5. Response
    metadata := make(map[string]string)
    metadata["pageCount"] = strconv.Itoa(pgCount)
    metadata["lastPage"] = strconv.Itoa(LastPage)

    public.SendJSONResponse(w, games, metadata)
}


