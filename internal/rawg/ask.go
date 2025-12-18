package rawg

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/nioWhiteHat/gaming-store-backend.git/internal/data"
)

type Communicate struct {
	ApiKey string
}

func (c *Communicate) GetGenres(dbpool *pgxpool.Pool, ctx context.Context) ([]string,error) {
	url:= fmt.Sprintf("https://api.rawg.io/api/genres?key=%s", c.ApiKey)
	
	res, err := http.Get(url)
	if err!=nil{
		return nil,err
	}
	body, errs := io.ReadAll(res.Body)
	if errs!=nil{
		return nil,err

	}
	log.Println("request succeded and body parsed")
	var genresp data.GenreResponse
	err = json.Unmarshal(body, &genresp)
	
	if err!=nil{
		return nil,err
	}
	
	
	type GenNames struct{
		slugs []string
		names []string
	}
	var gn GenNames
	for _,gen := range genresp.Gens{
		gn.names = append(gn.names, gen.Name)
		gn.slugs = append(gn.slugs, gen.Slug)
	}
	gm := data.GenresModel{DB: dbpool}
	log.Println("right before fillGenre call")
	gm.FillGenreTable(ctx, gn.slugs, gn.names)
	var slugs []string
	log.Printf("the first gen is %s, the first game of the gen is %s", genresp.Gens[0].Slug, *genresp.Gens[0].GameSlugs[0].Slug)
	for _,gen := range genresp.Gens{
		log.Println("in the gen for loop")
		for i,game := range gen.GameSlugs{
			log.Println(*game.Slug)
			slugs = append(slugs, *game.Slug)
			
			if i>25{
				break
			}
		}
	}
	return slugs,nil

}

func (c *Communicate) FillAllPipeline(dbpool *pgxpool.Pool, ctx context.Context)error{
	log.Println("DEBUG: 71")
	tf := data.TableFiller{
		DB: dbpool,
	}
	log.Println("DEBUG: 75")
	err := tf.GenerateFakeUsers(50,ctx)
	log.Println("DEBUG: 77")
	if err!=nil{
		return err
	}
	slugs,err := c.GetGenres(dbpool,ctx)
	if err!=nil{
		return err
	}
	err = c.GetStores(ctx,dbpool)
	if err!=nil{
		return err
	}
	err = c.GetPlatforms(ctx,dbpool)
	if err!=nil{
		return err
	}
	log.Println("all good with fill genres")
	
	gamemodel := data.GameModel{DB: dbpool}
	for _,s := range slugs{
		game,scr := c.Ask(s)
		err := gamemodel.InsertGame(ctx, game, scr)
		if err != nil{
			log.Println("DEBUG: line 79 in pipeline")
			return err
		}
	}
	
	err = tf.CreateVendorReviews(ctx)
	if err!=nil{
		log.Printf("Debug: 107 %s",err)
		return err
	}
	err = tf.CreateKeys(ctx)
	if err!=nil{
		log.Printf("Debug: 117 %s",err)
		return err
	}
	err = tf.CreateVendorKey(ctx)
	if err!=nil{
		log.Printf("Debug: 122 %s",err)
		return err
	}
	err = tf.CreateVendorGames(ctx)
	if err!=nil{
		log.Printf("Debug: 112 %s",err)
		return err
	}
	
	
	
	
	
	return nil
}


func (c *Communicate) GetPlatforms(ctx context.Context, dbpool *pgxpool.Pool) error{

	url := fmt.Sprintf("https://api.rawg.io/api/platforms?key=%s", c.ApiKey)


	resp, err := http.Get(url)
	if err != nil {
		log.Printf("failed to call platforms from api: %v", err)
		return err
	}
	
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Printf("API returned non-200 status: %d", resp.StatusCode)
		return err
	}

	
	var response data.PlatformResponse
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		log.Printf("failed to decode json: %v", err)
		return err
	}

	log.Printf("Successfully decoded %d platforms.", len(response.Results))
	log.Printf("Total count available: %d", response.Count)
	

	for _, p := range response.Results {
		log.Printf("Platform: %s (ID: %d)", p.Name, p.ID)
	}
	var tf = data.TableFiller{
		DB: dbpool,
	}
	tf.InsertPlatforms(ctx,response)
	return nil
}
func (c *Communicate) Ask( slug string) (data.Game, data.ScreenshotsResponse) {
	url := fmt.Sprintf("https://api.rawg.io/api/games/%s?key=%s", slug, c.ApiKey)

	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(resp.Status)
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("couldnt parse resp.body")
	}
	var game data.Game

	

	err = json.Unmarshal(body, &game)
	if err != nil {
		id := *game.ExternalID
		log.Println(id)
		log.Fatal("couldnt parse body to game.slug")
	}
	



	url = fmt.Sprintf("https://api.rawg.io/api/games/%s/screenshots?key=%s", strconv.Itoa(*(game.ExternalID)), c.ApiKey)
	resp2, err2 := http.Get(url)
	if err2 != nil {
		log.Fatal(err2)
	}
	defer resp2.Body.Close()
	body, err2 = io.ReadAll(resp2.Body)
	if err2 != nil {
		log.Fatal(err2)
	}
	var screenshotres data.ScreenshotsResponse
	json.Unmarshal(body, &screenshotres)
	max := 10
	if len(screenshotres.Results) > max {
		screenshotres.Results = screenshotres.Results[:max]
	}
	

	return game, screenshotres

}

func (c *Communicate) GetStores(ctx context.Context, pool *pgxpool.Pool)error{
	url := fmt.Sprintf("https://api.rawg.io/api/stores?key=%s",c.ApiKey)
	resp,err := http.Get(url)
	if err!=nil{
		return err
	}
	log.Println("DEBUG: 181")
	body, err:= io.ReadAll(resp.Body)
	if err!=nil{
		log.Println("DEBUG: 184")
		return err
	}
	log.Println("DEBUG: 187")
	log.Println(string(body))
	var response data.APIResponse
	err = json.Unmarshal(body,&response)
	log.Println("DEBUG: 191")
	if err!=nil{
		log.Println("DEBUG: 192")
		return err
	}
	tf := data.TableFiller{
		DB: pool,
	}
	log.Println("before calling insert on stores")
	err = tf.InsertStores(ctx,response) 
	if err!=nil{
		log.Fatal(err)
	}
	return nil
	 
}
