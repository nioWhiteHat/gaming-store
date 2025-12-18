package data

import (
	"context"
	"time"

	"golang.org/x/crypto/bcrypt"
)

func (s *SessionsModel) InsertToSessions(ctx context.Context, sessionID string, status string, userId int64, creationTime time.Time, expiryTime time.Time) error{
	hash,err := bcrypt.GenerateFromPassword([]byte(sessionID),12)
	if err!=nil{
		return ErrInternalServerError
	}
	query := `INSERT INTO SESSIONS(session_id,user_id,status,created,expieres) VALUES ($1,$2,$3,$4)` 
	_,err1:=s.DB.Exec(ctx,query,hash,status,userId,creationTime,expiryTime)
	if err1!=nil{
		return ErrDbConn
	}
	return nil
}



