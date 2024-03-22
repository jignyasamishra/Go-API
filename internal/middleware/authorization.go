package middleware

import (
	"errors"
	"net/http"

	"github.com/jignyasamishra/Go-API/api"
	log "github.com/sirupsen/logrus"
)

var UnAuthroizedError = errors.New("Invalid username or token.")

func Authorization(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		var username string = r.URL.Query().Get("username")
		var token = r.Header.Get("Authorization")
		var err error

		if username == "" || token == "" {
			log.Error(UnAuthroizedError)
			api.RequestErrorHandler(w, UnAuthroizedError)
			return
		}
		var loginDetails *tools.loginDetails
		loginDetails = (*database).GetUserLoginDetails(username)
		if loginDetails == nil || (token != (*loginDetails).AuthToken) {
			api.RequestErrorHandler(w, UnAuthroizedError)
			return
		}
		next.ServeHTTP(w, r)
	})

}
