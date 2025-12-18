package middlewares

import (
	"log"
	"net/http"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)
type Middleware func(http.Handler) http.Handler

func CreateStack(mw ...Middleware)Middleware{
	return func(next http.Handler)http.Handler{
		for i := len(mw) - 1; i >= 0; i-- {
			next = mw[i](next)
		}
		return next
	}
}

func NewAuthMiddleware(db *pgxpool.Pool, userType string) Middleware{
	return func(next http.Handler) http.Handler{
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
			session_cookie, err := r.Cookie("session_id")
			if err!=nil{
				http.Error(w,"no cookie", http.StatusNotAcceptable)
				return 
			}
			if time.Now().After(session_cookie.Expires){
				http.Error(w,"expiered",http.StatusUnauthorized)
				return 
			}
			var status string
			var userid string
			err = db.QueryRow(r.Context(), "SELECT status,user_id FROM sessions WHERE session_id = $1",session_cookie.Value).Scan(&status,&userid)
			if err!=nil{
				if  err!= pgx.ErrNoRows{
					http.Error(w,"internal server error", http.StatusInternalServerError)
				}else {
					http.Error(w,"the session id is invalid", http.StatusBadRequest)
				}
				return 
			}
			if status != "public" && status != userType{
				http.Error(w, "status unothorized", http.StatusUnauthorized)
				return 
			}
			next.ServeHTTP(w,r)
			

		})

	}
}

func LogMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		executionStart := time.Now()

		next.ServeHTTP(w, r)

		log.Printf("%s %s %v", r.Method, r.RequestURI, time.Since(executionStart))
	})
}
