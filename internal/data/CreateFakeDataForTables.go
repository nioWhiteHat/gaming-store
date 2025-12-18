package data

import (
	"context"
	"log"
	"math/rand"
	"strings"
	"time"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/nioWhiteHat/gaming-store-backend.git/internal/utils"
)

func (tf *TableFiller) GenerateFakeUsers(n int, ctx context.Context) (error) {
	var users []User

	for i := 0; i < n; i++ {
		user := User{
			Utype:     getRandomType(),
			Username: gofakeit.Username(),
			Password: "123456",
			Email:    gofakeit.Email(),
			ImgUrl:    gofakeit.ImageURL(200, 200), 
			Bio:      gofakeit.Sentence(10),
		}

		users = append(users, user)
	}
	
	query := `insert into users(type,username,password,email,image,bio)values($1,$2,$3,$4,$5,$6)`
	for _,u := range users{
		
		pas,err := utils.Encrypt(u.Password)
		if err!=nil{
			return err
		}

		_,err =  tf.DB.Exec(ctx,query,&u.Utype,&u.Username,&pas,&u.Email,&u.ImgUrl,&u.Bio)
		if err!=nil{
			log.Printf("%s,%s,%s,%s,%s,%s",u.Utype,u.Username,pas,u.Email,u.ImgUrl,u.Bio)
			log.Panicln(err)
			
			return err
		}	
		
	}
	return nil
	
}


func getRandomType() string {
    n := rand.Intn(100)
    if n < 5 {
        return "admin"
    } else if n < 35 {
        return "vendor"
    }
    return "client"
}

func (tf *TableFiller) CreateVendorReviews(ctx context.Context)error{
	query1 := `select id from users where type = $1`
	query2 := `select id from users where type = $1`
	clients,err := tf.DB.Query(ctx, query1, "client")
	if err!=nil{
		return err
	}
	vendors,err:= tf.DB.Query(ctx, query2, "vendor")
	if err!=nil{
		return err
	}
	var vendorids []int
	var clientids []int
	for clients.Next(){
		var id int
		err := clients.Scan(&id)
		if err!=nil{
			return err
		}
		clientids = append(clientids, id)
	}
	for vendors.Next(){
		var id int
		err := vendors.Scan(&id)
		if err!=nil{
			return err
		}
		vendorids = append(vendorids, id)
	}
	for _,c := range clientids{
		numOfComments := GenerateRandomInts(1,3,20)
		query := `insert into vendor_reviews(client_id, vendor_id, review, rating) VALUES ($1, $2, $3, $4)`
		vendorIdsTheClientCommented := GenerateRandomInts(numOfComments[0],0,len(vendorids)-1)
		for i:=0; i <= len(vendorIdsTheClientCommented)-1; i++{
			review :=  gofakeit.Sentence(10)
			rating := GenerateRandomInts(1, 0, 5)
			_,err := tf.DB.Exec(ctx,query,c,vendorids[vendorIdsTheClientCommented[i]],review,rating[0])
			if err!=nil{
				return err
			}
		}
	}
	return nil


}


func GenerateRandomInts(n, min, max int) []int {
 
    rand.Seed(time.Now().UnixNano())

    nums := make([]int, n)
    for i := 0; i < n; i++ {
        
        nums[i] = rand.Intn(max - min + 1) + min
    }
    return nums
}

func (tf *TableFiller) InsertPlatforms(ctx context.Context,plats PlatformResponse)(error,[]int){
	query := `INSERT INTO PLATFORMS(name,image) values($1,$2)`
	var ids []int
	for _,plat :=range plats.Results{
		var id int
		err := tf.DB.QueryRow(ctx,query+" RETURNING id",plat.Name,plat.Image).Scan(&id)
		if err!=nil{
			return err,nil
		}
		ids = append(ids, id)

	}
	return nil,ids

}

func (tf *TableFiller) InsertStores(ctx context.Context, response APIResponse)error{
	log.Printf("in insertStores, also response has %v stores", len(response.Results))
	for i:=0; i<=len(response.Results)-1; i++{
		log.Println("in for loop")
		store := response.Results[i]
		log.Println("DEBUG: 319")
		query := `insert into stores(name,url,imageurl)values($1,$2,$3)`
		log.Printf("%s %s %s ",store.Name,store.Domain,store.ImageBackground)
		_,err := tf.DB.Exec(ctx,query,store.Name,store.Domain,store.ImageBackground)
		log.Printf("%s %s %s ",store.Name,store.Domain,store.ImageBackground)
		if err!=nil{
			return err
		}
	}
	return nil
}

func (tf *TableFiller) CreateVendorGames(ctx context.Context) error {
	queryUsers := `select id from users where type = $1`
	
	
	queryGames := `select DISTINCT k.game_id from vendor_keys as vk join keys as k on k.id = vk.key_id where vendor_id = $1`

	vendors, err := tf.DB.Query(ctx, queryUsers, "vendor")
	if err != nil {
		return err
	}

	defer vendors.Close()

	var vendorids []int
	for vendors.Next() {
		var id int
		if err := vendors.Scan(&id); err != nil {
			return err
		}
		vendorids = append(vendorids, id)
	}
 
    vendors.Close() 

	for _, vendorID := range vendorids {
		games, err := tf.DB.Query(ctx, queryGames, vendorID)
		if err != nil {
			return err
		}

		for games.Next() {
			var gameID int
			if err := games.Scan(&gameID); err != nil {
                games.Close() 
				return err
			}

		
			_, err = tf.DB.Exec(ctx, 
                "INSERT INTO vendor_games(vendor_id, game_id) VALUES($1, $2) ON CONFLICT DO NOTHING", 
                vendorID, gameID)
			if err != nil {
                games.Close()
				return err
			}
		}
     
		games.Close()
	}

	return nil
}

func (tf *TableFiller) CreateKeys(ctx context.Context)error{
	games,err:=tf.DB.Query(ctx,"select id from games")
	if err!=nil{
		return err
	}
	var gameIds []int
	for games.Next(){
		var id int
		err:= games.Scan(&id)
		if err!=nil{
			return err
		}
		gameIds = append(gameIds, id)
	}
	stores,err:=tf.DB.Query(ctx,"select id from stores")
	if err!=nil{
		return err
	}
	var storeIds []int
	for stores.Next(){
		var id int
		err:= stores.Scan(&id)
		if err!=nil{
			return err
		}
		storeIds = append(storeIds, id)
	}
	query:=`insert into keys(game_id,store_id,key_hash,price)values($1,$2,$3,$4)`
	for i:=0; i<=len(gameIds)-1; i++{
		for k := 0; k<=1; k++{
			for j:=0; j<=len(storeIds)-5; j++{
				key,err := utils.Encrypt(GenerateGameKey())
				if err!=nil{
					return err
				}
				price := 49.99
				_,err = tf.DB.Exec(ctx,query,gameIds[i],storeIds[j],key,price)
				if err!=nil{
					return err
				}
			}
		}		
	}
	return nil
}

  

func GenerateGameKey() string {
	
	const charset = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	

	numSegments := 3    
	segmentLength := 5  
	

	rand.Seed(time.Now().UnixNano())

	var segments []string


	for i := 0; i < numSegments; i++ {
		segment := make([]byte, segmentLength)
		for j := 0; j < segmentLength; j++ {
			
			randomIndex := rand.Intn(len(charset))
			segment[j] = charset[randomIndex]
		}
		segments = append(segments, string(segment))
	}
	return strings.Join(segments, "-")
}



func (tf *TableFiller) CreateVendorKey(ctx context.Context) error {
	// FIX 1: Use single quotes for string literals in SQL
	queryVendors := `SELECT id FROM users WHERE type = 'vendor'`
	queryKeys := `SELECT id FROM keys`

	// 1. Get Vendors
	vendorsRows, err := tf.DB.Query(ctx, queryVendors)
	if err != nil {
		return err
	}
	defer vendorsRows.Close() 

	var vendorIds []int
	for vendorsRows.Next() {
		var id int
		if err := vendorsRows.Scan(&id); err != nil {
			return err
		}
		vendorIds = append(vendorIds, id)
	}
	vendorsRows.Close()

	if len(vendorIds) == 0 {
		return nil // Avoid divide by zero later
	}

	// 2. Get Keys
	keysRows, err := tf.DB.Query(ctx, queryKeys)
	if err != nil {
		return err
	}
	defer keysRows.Close() // FIX 2: Close rows

	var keyIds []int
	for keysRows.Next() {
		var id int
		if err := keysRows.Scan(&id); err != nil {
			return err
		}
		keyIds = append(keyIds, id)
	}
	keysRows.Close()

	if len(keyIds) == 0 {
		return nil
	}

	// 3. Use a Transaction for performance
	tx, err := tf.DB.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	queryInsert := `INSERT INTO vendor_keys(vendor_id, key_id) VALUES ($1, $2)`
	
	// FIX 3: Simplified Logic (Round Robin distribution)
	// We iterate through ALL keys and assign them to vendors sequentially.
	for i, keyID := range keyIds {
		// Use modulo (%) to cycle through vendors: 0, 1, 2, 0, 1, 2...
		vendorIndex := i % len(vendorIds)
		vendorID := vendorIds[vendorIndex]

		_, err := tx.Exec(ctx, queryInsert, vendorID, keyID)
		if err != nil {
			return err
		}
	}

	return tx.Commit(ctx)
}