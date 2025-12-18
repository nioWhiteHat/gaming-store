package public

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"time"

	"net/http"

	"github.com/nioWhiteHat/gaming-store-backend.git/internal/data"
	errors "github.com/nioWhiteHat/gaming-store-backend.git/internal/errors"
	"github.com/nioWhiteHat/gaming-store-backend.git/internal/utils"
)

func (p *Public) SignIn(w http.ResponseWriter, r *http.Request) {
	
	type UserData struct {
		Username string `json:"username"`
		Password string `json:"password"`
		Email    string `json:"email"`
	}
	ctx := r.Context()
	var userdata UserData
	err := json.NewDecoder(r.Body).Decode(&userdata)

	if err != nil {
		errors.ErrorHandler(w, data.ErrInternalServerError)
		return
	}


	err,usr := p.Sm.PopulateUser(ctx, userdata.Username, userdata.Password, userdata.Email)
	if err!=nil{
		errors.ErrorHandler(w,err)
		return
	}
	
	
	sid,expiryTime,error := p.CreateSessionID(ctx,usr.Utype,int64(usr.Id))

	if error!=nil{
		errors.ErrorHandler(w,err)
		return
	}

	cookie:= &http.Cookie{
		Name: "session_id",
		Value: sid,
		Expires: expiryTime,
		HttpOnly: true,
		Secure: true,
		SameSite: http.SameSiteLaxMode,
	}

	http.SetCookie(w,cookie)

	err = p.SendValRes(w,&usr)
	if err!=nil{
		errors.ErrorHandler(w, err)
		return
	}
	

}

func (p *Public) CreateSessionID(ctx context.Context, status string, userId int64)(string,time.Time,error) {
	randomBytes := make([]byte, 32)

	creationTime := time.Now().UTC()

	expiryTime := creationTime.Add(2 * time.Hour)
	
	_, err := rand.Read(randomBytes)
	if err != nil {
		
		return "", expiryTime, data.ErrInternalServerError
	}
	
	randomString := base64.URLEncoding.EncodeToString(randomBytes)
	randomString,err = utils.Hash(randomString)
	if err!=nil{
		return "", expiryTime, data.ErrInternalServerError
	}
	sessionID := fmt.Sprintf("%s_%s", status, randomString)
	
	err = p.Ses.InsertToSessions(ctx,sessionID,status,userId,creationTime,expiryTime)
	if err!=nil{
		return "", expiryTime, err
	}
	return sessionID,expiryTime,nil
	

}


func (p *Public) SendValRes(w http.ResponseWriter, usr *data.User) error{
	
	
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err:=json.NewEncoder(w).Encode(*usr)
	if err!=nil{
		return data.ErrInternalServerError
	}
	return nil

}

func SendJSONResponse[T any](w http.ResponseWriter, data []T, metadata ...interface{}) {
	response := map[string]interface{}{
		
		"data":    data,
	}

	if len(metadata) > 0 {
		response["metadata"] = metadata[0]
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}