package errors

import (
	"errors"
	"net/http"

	"github.com/nioWhiteHat/gaming-store-backend.git/internal/data"

)

func  ErrorHandler(w http.ResponseWriter,  err error) {

	switch {
	case errors.Is(err, data.ErrMismatchedHashAndPassword):
		
		writeError(w, http.StatusUnauthorized, "invalid authentication credentials")

	case errors.Is(err, data.ErrInvalidCredentials):
		
		writeError(w, http.StatusUnauthorized, "invalid authentication credentials")


	case errors.Is(err, data.ErrDbConn):
		writeError(w, http.StatusUnauthorized, "dbcon is facing a problem")


	case errors.Is(err,data.ErrInternalServerError):
		writeError(w, http.StatusInternalServerError, "Server is facing a problem")


	default:
		
		writeError(w, http.StatusInternalServerError, "the server encountered a problem and could not process your request")
	}
}

func  writeError(w http.ResponseWriter, status int, message string){
	w.WriteHeader(status)
	w.Write([]byte(message))
}