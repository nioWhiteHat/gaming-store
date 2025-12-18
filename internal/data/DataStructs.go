package data

import (
	"time"

	
)
type Order struct{
	Id int
	Price int
	Created_at time.Time
	Game_key string
	Vendor_name string
	Game_id int
	Game_title string
	
}
type Vendor_temp struct {
	GameId   int
	Name     string
	Image    string
	Price    float64
	Store_id int
	Store    string
	StoreUrl string
	StoreImg string
	Key_id   int
}

type CartGame struct {
	GameId int
	GameTitle string
	GameImg string
	Price float32
	KeyId int
	VendorId int
	ClientId int

}
type Game_temp struct {
	
	Name              string    `db:"name" json:"name"`
	Description       string    `db:"description" json:"description"`
	Released          time.Time `db:"released" json:"released"` 
	RatingRecommended int       `db:"rating_recommended" json:"rating_recommended"`
	RatingMeh         int       `db:"rating_meh" json:"rating_meh"`
	RatingExceptional int       `db:"rating_exceptional" json:"rating_exceptional"`
	RatingSkip        int       `db:"rating_skip" json:"rating_skip"`
	Image             string    `db:"image" json:"image"`
	MainImage         string    `db:"main_image" json:"main_image"`
	ID                int       `db:"id" json:"id"`

	
	ScreenshotURL string `db:"screenshot_url" json:"screenshot_url"`
	Width         int    `db:"width" json:"width"`
	Height        int    `db:"height" json:"height"`
}
type Vendors_keys struct{
	vk		Vendor_keys
}

type Vendor_keys struct{
	Name string
	Image string
	keys []key

}
type key struct{
	store		string
	price		float64
} 

type User struct{
	Id int					`json:"id"`
    Utype string			`json:"type"`
    Username string			`json:"username"`
    Password string			`json:"-"`
    Email string			`json:"email"`
    Created_at time.Time	`json:"creation"`
	ImgUrl	string			`json:"img"`
	Bio string
}


type Store struct {
	ID              int    `json:"id"`
	Name            string `json:"name"`
	Domain          string `json:"domain"`
	Slug            string `json:"slug"`
	GamesCount      int    `json:"games_count"`
	ImageBackground string `json:"image_background"`
	
}

// APIResponse represents the top-level JSON object.
type APIResponse struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`    
	Previous *string `json:"previous"` 
	Results  []Store `json:"results"`
}










type GenreMetadata struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Slug string `json:"slug"`
}


type Game struct {
	Id				*int	   `json:"realid"`
	ExternalID      *int       `json:"id"`
	Slug            *string    `json:"slug"`
	Name            *string    `json:"name"`
	Description     *string    `json:"description"`
	Released        *string	   `json:"released"`
	Main_image *string    `json:"background_image"`
	Image           *string    `json:"background_image_additional"`
	Ratings         *[]Rating  `json:"ratings"`
	Platforms       *[]PlatformInfo `json:"platforms"`
	Genres 			*[]GenreMetadata 	`json:"genres"`
	Added 			   int			`json:"added"`
	Region			string
	Platform        string			`json:"platform"`
	Screenshots     []Screenshot
	Average_price 	float64		`json:"Average_price"`


}

type Rating struct {
	ID      *int     `json:"id"`
	Title   *string  `json:"title"`
	Count   *int     `json:"count"`
	Percent *float64 `json:"percent"`
}

type PlatformInfo struct {
	Platform *Platform `json:"platform"`
}


type GenreResponse struct{
	Count int `json:"count"`
	Gens []Genre `json:"results"`
}

type Genre struct {
	Id	string `json:"id"`
	Slug string `json:"slug"`
	Name string `json:"name"`
	GameSlugs []Game `json:"games"`
}

type Screenshot struct {
	ID        int    `json:"id"`
	Image     string `json:"image"`
	Width     int    `json:"width"`
	Height    int    `json:"height"`
	IsDeleted bool   `json:"is_deleted"`
}

type ScreenshotsResponse struct {
	Count    int          `json:"count"`
	Next     *string      `json:"next"`    
	Previous *string      `json:"previous"` 
	Results  []Screenshot `json:"results"`
}

type GameDetails struct {
	ID    int    `json:"id"`
	Slug  string `json:"slug"`
	Name  string `json:"name"`
	Added int    `json:"added"`
}

type Platform struct {
	ID              int           `json:"id"`
	Name            string        `json:"name"`
	Slug            string        `json:"slug"`
	GamesCount      int           `json:"games_count"`
	ImageBackground string        `json:"image_background"`
	Image           *string       `json:"image"`      
	YearStart       *int          `json:"year_start"` 
	YearEnd         *int          `json:"year_end"`   
	Games           []GameDetails `json:"games"`
}

type PlatformResponse struct {
	Count    int        `json:"count"`
	Next     *string    `json:"next"`     
	Previous *string    `json:"previous"` 
	Results  []Platform `json:"results"`
}



