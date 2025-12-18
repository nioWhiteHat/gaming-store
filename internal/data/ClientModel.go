package data

import (
	"context"
)

func (c *ClientModel) ViewClientHistory(ctx context.Context, usrId int) (error,[]Order) {
	query := `
		SELECT o.id, o.price_at_sale, o.created_at, o.game_key, u.name, o.game_id, g.title
		FROM orders as o join users as u on u.id = o.vendor_id join games as g on u.gid = g.id 
		WHERE client_id =$1 
	`
	rows, err := c.DB.Query(ctx,query,usrId)
	if err!=nil{
		return ErrDbConn,nil
	}
	var orders []Order
	for rows.Next(){
		var order Order
		err = rows.Scan(&order.Id,&order.Price,&order.Created_at,&order.Game_key,&order.Vendor_name,&order.Game_id,&order.Game_title)
		if err!=nil{
			return err,nil;
		}
		orders = append(orders, order)

	}
	return nil,orders

}

func (c *ClientModel) ViewCart(ctx context.Context, client_id string) (error,[]CartGame){
	var query = `
	select g.id,g.title,g.main_image,k.price,c.key_id,v.vendor_id
	FROM cart as c join games as g on g.id=c.game_id join game_keys as k on k.id=c.key_id join vendor_keys as v on v.key_id=k.id
	where c.user_id = $1
	`
	
	rows, err := c.DB.Query(ctx,query,client_id)
	if err!=nil{
		return ErrDbConn,nil
	}
	var games []CartGame
	for rows.Next(){
		var game CartGame
		err = rows.Scan(&game.GameId,&game.GameTitle,&game.GameImg,&game.Price,&game.KeyId,&game.VendorId)
		if err!=nil{
			return err,nil
		}
		games = append(games, game)
	}
	if len(games)!=0{
		return nil,games
	}else{
		return nil,nil
	}
}

func (c *ClientModel) AddToCart(ctx context.Context, client_id int, keyId int, gameId int)error{
	query := `insert into cart(user_id,game_id,key_id)values($1,$2,$3)`
	_,err := c.DB.Exec(ctx,query,client_id,gameId,keyId)
	if err!=nil{
		return err
	}
	return nil


}

func (c *ClientModel) RemoveFromCart(ctx context.Context, client_id int, keyId int, gameId int)error{
	query := `delete from cart where user_id = $1 and client_id = $2 and key_id = $3`
	_,err := c.DB.Exec(ctx,query,client_id,gameId,keyId)
	if err!=nil{
		return err
	}
	return nil


}


	