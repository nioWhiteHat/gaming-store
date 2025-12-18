package data

import (
	"context"

	"log"
)

func (gm *GameModel) GetGamesSql(ctx context.Context, genre string, offset int, priceMin int, priceMax int, platform string) ([]Game, error) {
	
	log.Printf("min price = %v",priceMin)
	query := `
	SELECT
    games.id,
    games.name,
	games.main_image,
    AVG(keys.price) AS average_price
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
		AND keys.price <= $2
	GROUP BY
		games.id,
		games.name
	ORDER BY
		games.id ASC    -- Required for consistent results
	LIMIT 10            -- How many results to show
	OFFSET $5 ;
	`

	rows, err := gm.DB.Query(ctx,query,genre,priceMax,priceMin,platform,offset)
	var games []Game
	if err != nil {
		log.Println("Debug:23")
		return games, err
	}
	defer rows.Close()

	for rows.Next() {
		
		var g Game
		err1 := rows.Scan(&g.Id, &g.Name, &g.Main_image, &g.Average_price)
		if err1 != nil {
			log.Println("Debug:32")
			return games, err
		}
		games = append(games, g)
		
	}
	
	return games, nil

}

func (gm *GameModel) GetGameData(ctx context.Context, gameId int) (error, []Game_temp, []Vendor_temp) {
	// Unused struct kept to preserve original logic/intent
	type Game_details struct {
	}

	query1 := `
	SELECT 
		g.name,
		g.description,
		g.released,
		g.rating_recommended,
		g.rating_meh,
		g.rating_exceptional,
		g.rating_skip,
		g.image,
		g.main_image,
		g.id,
		s.screenshot_url,
		s.width,
		s.height 
	FROM 
		games as g 
		LEFT JOIN game_screenshots as gs on g.id = gs.game_id 
		LEFT JOIN screenshots as s on s.id = gs.screenshot_id 
	WHERE 
    	g.id = $1;
	`

	// Fixed missing comma between s.img and k.id
	query2 := `
	SELECT DISTINCT ON (v.id, k.store_id) 
		vg.game_id,
		v.name,
		v.image,
		k.price,
		k.store_id,   
		s.name as store_name,
		s.url,
		s.img,
		k.id
	FROM
		vendor_games as vg 
		JOIN users as v ON vg.vendor_id = v.id
		JOIN vendor_keys as vk ON vk.vendor_id = vg.vendor_id 
		JOIN keys as k ON k.id = vk.key_id
		JOIN game_keys as gk ON vk.key_id = gk.key_id 
		JOIN stores as s ON s.id = k.store_id 
	WHERE 
		vg.game_id = $1 
		AND gk.game_id = $1
	ORDER BY 
		v.id, k.store_id, k.price ASC;
	`

	rows, err := gm.DB.Query(ctx, query2, gameId)
	if err != nil {
		return err, nil, nil
	}
	// Defer close is safer, but sticking to your manual close logic
	// defer rows.Close() 
	
	var Vendor_data []Vendor_temp
	for rows.Next() {
		var vdata Vendor_temp
		// Assuming Vendor_temp struct fields align with these types
		err := rows.Scan(
			&vdata.GameId, 
			&vdata.Name, 
			&vdata.Image, 
			&vdata.Price, 
			&vdata.Store_id, 
			&vdata.Store, 
			&vdata.StoreUrl, 
			&vdata.StoreImg, 
			&vdata.Key_id,
		)
		if err != nil {
			rows.Close()
			return err, nil, nil
		}
		Vendor_data = append(Vendor_data, vdata)
	}
	rows.Close() // Close previous query before starting new one

	var Game_data []Game_temp
	
	// Reusing 'rows' variable
	rows, err = gm.DB.Query(ctx, query1, gameId)
	if err != nil {
		return err, nil, nil
	}
	defer rows.Close() // Ensure this closes eventually

	for rows.Next() {
		var game_temp Game_temp
		// Fixed: Added & to RatingSkip
		// Fixed: Added RatingExceptional to match SELECT list order
		err = rows.Scan(
			&game_temp.Name,
			&game_temp.Description,
			&game_temp.Released,
			&game_temp.RatingRecommended,
			&game_temp.RatingMeh,
			&game_temp.RatingExceptional, // Added this to match SQL
			&game_temp.RatingSkip,        // Added '&'
			&game_temp.Image,
			&game_temp.MainImage,
			&game_temp.ID,
			&game_temp.ScreenshotURL,
			&game_temp.Width,
			&game_temp.Height,
		)
		if err != nil {
			return err, nil, nil
		}
		Game_data = append(Game_data, game_temp)
	}

	return nil, Game_data, Vendor_data
}

func (g *GameModel) GetPlatIds(ctx context.Context,plats []PlatformInfo)(error,[]int){
	var platIds []int
	query := `select id from platforms where name = $1`
	
	for i:=0; i<=len(plats)-1; i++{
		var id int
		
		err := g.DB.QueryRow(ctx, query, plats[i].Platform.Name).Scan(&id)
		if err!=nil{
			log.Println(err)
			continue
		}
		platIds = append(platIds, id)

	}
	

	return nil, platIds
}
func (gm *GameModel) InsertGame(ctx context.Context, game Game, screenshots ScreenshotsResponse) error {
	log.Println("in insetGame")
	tx, err := gm.DB.Begin(ctx)
	if err != nil {
		
		return err
	}
	defer tx.Rollback(ctx)

	ratings := *(game.Ratings)
	recommended := *ratings[0].Count
	meh := *ratings[1].Count
	skip := *ratings[2].Count
	added := meh+recommended+skip
	exceptional := *ratings[3].Count
	plats := *(game.Platforms)
	log.Printf("%s, %s, %v, plat in insert game",plats[0].Platform.Name, *game.Name, len(plats))
	
	err,platIds := gm.GetPlatIds(ctx,plats)

	if err!=nil{
		log.Printf("DEBUG: line 191 %s",err)
		return err
	}
	log.Println("DEBUG: in 194")
	genres := *(game.Genres)
	var genre_slugs []string
	for _, g := range genres {
		genre_slugs = append(genre_slugs, g.Slug)
		log.Println("DEBUG: in 186")
	}
	log.Println("DEBUG: in 188")
	BaseQuery := `SELECT id from genre where slug=ANY($1)`
	var genre_ids []int
	rows, err1 := tx.Query(ctx, BaseQuery, genre_slugs)
	if err1 != nil {
		log.Println("DEBUG: in 192")
		return err1
	}

	defer rows.Close()
	for rows.Next() {
		var id int
		err := rows.Scan(&id)
		log.Println("DEBUG: in 200")
		if err != nil {
			return err
		}
		genre_ids = append(genre_ids, id)
	}

	log.Println("DEBUG: in 206")

	query := `
				INSERT INTO games (
				external_id,
				slug,
				name,
				description,
				released,
				main_image,
				image,
				rating_recommended,
				rating_meh,
				rating_exceptional,
				rating_skip,
				added,
				platform_id
				
			) VALUES (
				$1,  -- external_id
				$2,  -- slug
				$3,  -- name
				$4,  -- description
				$5,  -- released
				$6,  -- background_image
				$7,  -- image
				$8,  -- rating_recommended
				$9,  -- rating_meh
				$10, -- rating_exceptional
				$11, -- rating_skip
				$12, -- added
				$13 -- platform
			
			)
			`
	var game_id int
	var screenshotIds []int
	for _, s := range screenshots.Results{
		log.Println("in the for loop for the screenshots")
		var newScreenshotID int

		err := tx.QueryRow(ctx,
			`
			INSERT INTO screenshots (screenshot_url, width, height)
			VALUES ($1, $2, $3)
			RETURNING id
			`,
			s.Image, s.Width, s.Height).Scan(&newScreenshotID)

		if err != nil {
			log.Println("DEBUG in line 273 mabye error in the screenshots insertion ")
			return err
		}
		screenshotIds = append(screenshotIds, newScreenshotID)
	}
	for _, platid := range  platIds{
		log.Println("in the for loop for the platids")
		err = tx.QueryRow(ctx, query+" RETURNING id",
		game.ExternalID,
		game.Slug,
		game.Name,
		game.Description,
		game.Released,
		game.Main_image,
		game.Image,
		recommended,
		meh,
		exceptional,
		skip,
		added,
		platid).Scan(&game_id)
		if err != nil {
			return err
		}
		BaseQuery = `INSERT INTO game_genres(game_id,genre_id)VALUES($1,$2)`
		for _, genre_id := range genre_ids {
			_, err := tx.Exec(ctx, BaseQuery, game_id, genre_id)
			if err != nil {
				log.Println("DEBUG in line 299")
				return err
			}
		}

		for _, s := range screenshotIds{
		
			q := `
			INSERT INTO game_screenshots (screenshot_id, game_id)
			VALUES ($1, $2)`
			_,err = tx.Exec(ctx,q,s,game_id)
			if err != nil {
				return err
			}

		}
	}

	err = tx.Commit(ctx)
	if err!= nil{
		log.Println("comit failed")
		return err
	}
	log.Println("out of the plats loop")
	return err
}



