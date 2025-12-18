package data

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/nioWhiteHat/gaming-store-backend.git/internal/utils"
)

func (sm *SharedModelMethods) PopulateUser(ctx context.Context, username string, password string, email string) (error, User) {
	query := `SELECT * FROM users WHERE email = $1 AND username = $2 AND password_hash = $3`

	var psw utils.Password
	psw.Set(password)

	var user User
	err := sm.DB.QueryRow(ctx, query, email, username, string(psw.Hashed)).Scan(&user.Id, &user.Utype, &user.Username, &user.Password, &user.Email, &user.Created_at, &user.ImgUrl)
	if err != nil {
		return ErrInvalidCredentials, user
	}
	return nil, user

}
type VendorDetails struct {
	ID       int
	Name     string
	Email    string
	Likes    int
	Dislikes int
	Reviews  []string
	Games    []Game
}

type VendorsGame struct {
	Title     string
	MainImage string
}
func (sm *SharedModelMethods) ViewVendor(ctx context.Context,vendorId int)(error, VendorDetails, []VendorsGame){
	var vData VendorDetails
	var vGames []VendorsGame
	batch := &pgx.Batch{}
	query := 
	`
	SELECT u.id,u.name,u.email,v.likes,v.dislikes
	FROM users as u 
	join vendor_likes_dislikes as v on v.id=u.id 
	where u.id = $1
	` 
	batch.Queue(query,vendorId)
	query2 :=
	`
	select vr.reviews from vendor_reviews as vr where vr.id=$1
	`
	batch.Queue(query2,vendorId)
	query3 := 
	`
	select g.title,g.main_image
	from vendor_games as vg  
	join games as g on g.id = vg.game_id
	where vg.id=$1
	`
	batch.Queue(query3,vendorId)

	br := sm.DB.SendBatch(ctx,batch)
	defer br.Close()

	err :=br.QueryRow().Scan(&vData.ID,&vData.Name,&vData.Email,&vData.Likes,&vData.Dislikes)
	if err!=nil{
		return ErrDbConn,vData,vGames	
	}


	rows, err := br.Query()
	if err!=nil{
		return ErrDbConn,vData,vGames
	}
	for rows.Next(){
		var rev string
		err = rows.Scan(&rev)
		if err!=nil{
			return err,vData,vGames
		}
		vData.Reviews = append(vData.Reviews, rev)
	}

	rows,err = br.Query()
	if err!=nil{
		return ErrDbConn,vData,vGames
	}

	for rows.Next(){
		var game VendorsGame
		err = rows.Scan(&game.Title,&game.MainImage)
		if err!=nil{
			return err,vData,vGames
		}
	}
	return nil, vData, vGames

}