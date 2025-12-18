package server

import (
	"log"
	"net/http"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"

	genres "github.com/nioWhiteHat/gaming-store-backend.git/internal/Genres"
	"github.com/nioWhiteHat/gaming-store-backend.git/internal/games"
	middlewares "github.com/nioWhiteHat/gaming-store-backend.git/internal/middlewares"
	"github.com/nioWhiteHat/gaming-store-backend.git/internal/public"
)

type Config struct {
	Addr string
}

type Application struct {
	Config Config
	Db     *pgxpool.Pool
}

func (app *Application) Mount() *http.ServeMux {
	public := public.Public{Db: app.Db}
	games := games.Game{Db: app.Db}
	genres := genres.Genre{Db: app.Db}
	public.ModelsInit()
	games.ModelsInit()
	genres.ModelsInit()

	ClientAuth := middlewares.NewAuthMiddleware(app.Db, "client") 
	AdminAuth := middlewares.NewAuthMiddleware(app.Db, "admin")
	VendorAuth := middlewares.NewAuthMiddleware(app.Db, "vendor")

	clientStack := middlewares.CreateStack(middlewares.LogMiddleware, ClientAuth)
	vendorStack := middlewares.CreateStack(middlewares.LogMiddleware, VendorAuth)
	adminStack := middlewares.CreateStack(middlewares.LogMiddleware, AdminAuth)
	 _ = clientStack
    _ = vendorStack
    _ = adminStack
	mux := http.NewServeMux()
	//public routes
	mux.HandleFunc("GET /getGenres", genres.GetGenres)
	mux.HandleFunc("GET /viewGames/{genre}/{page}/{priceMin}/{priceMax}/{platform}", games.GetGames)
	/*mux.HandleFunc("GET /gameData/{game_id}", games.GetGameData)
	mux.HandleFunc("GET /viewCritiques/{game_id}", games.ViewCrits)
	mux.HandleFunc("POST /client/sign_in", public.UserSignIn)
	mux.HandleFunc("POST /admin/sign_in", app.AdminSignIn)
	//client routes
	mux.HandleFunc("GET /client/viewCart", clientStack(http.HandlerFunc(app.ClientCart)))
	mux.HandleFunc("GET /client/viewWishList", clientStack(http.HandlerFunc(app.ClientWishList)))
	mux.HandleFunc("GET /client/viewVendor/{v_id}", clientStack(http.HandlerFunc(app.GetUserProfileData)))
	mux.HandleFunc("GET /client/order/{game_id}/{platform}/{v_id}", clientStack(http.HandlerFunc(app.OrderGame)))
	mux.HandleFunc("GET /client/orderHistory", clientStack(http.HandlerFunc(app.OrderHistory)))
	mux.HandleFunc("POST /client/messageVendor/{vendor_id}", clientStack(http.HandlerFunc(app.MessageFromClientToVendor)))
	mux.HandleFunc("POST /client/GetHelp", clientStack(http.HandlerFunc(app.GetHelp)))
	//vendor routes
	mux.HandleFunc("POST /vendor/addNewGame", vendorStack(http.HandlerFunc(app.AddGame)))
	mux.HandleFunc("POST /vendor/addKey", vendorStack(http.HandlerFunc(app.Addkey)))
	mux.HandleFunc("GET /vendor/browseGames", vendorStack(http.HandlerFunc(app.GetGames)))
	mux.HandleFunc("GET /vendor/viewGameDetails", vendorStack(http.HandlerFunc(app.GetStats)))
	mux.HandleFunc("GET /vendor/myGames", vendorStack(http.HandlerFunc(app.MyGames)))
	mux.HandleFunc("GET /vendor/myGameData/{game_id}", vendorStack(http.HandlerFunc(app.MyGameData)))
	mux.HandleFunc("POST /vendor/deleteKey/{key_id}", vendorStack(http.HandlerFunc(app.DelKey)))
	mux.HandleFunc("POST /vendor/deleteGame/{game_id}", vendorStack(http.HandlerFunc(app.)))
	mux.HandleFunc("GET /vendor/allKeys", vendorStack(http.HandlerFunc()))
	mux.HandleFunc("GET /vendor/ordersHistory/{filter}", vendorStack(http.HandlerFunc()))
	mux.HandleFunc("GET /vendor/openChatLog/{client_id}", vendorStack(http.HandlerFunc()))
	mux.HandleFunc("Post /vendor/sendMessage/{client_id}", vendorStack(http.HandlerFunc()))*/

	return mux
}

func (app *Application) Run(mux *http.ServeMux) error {
	srv := &http.Server{
		Addr:         app.Config.Addr,
		Handler:      mux,
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	log.Printf("Starting HTTP server on %s", app.Config.Addr)
	return srv.ListenAndServe()
}
