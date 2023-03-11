package app

import (
	"log"
	"net/http"

	"github.com/m3-app/backend/utils"
)

func (app AppBase) BaseMiddleware() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Add("Content-Type", "application/json")
			next.ServeHTTP(w, r)
		})
	}
}

func (app AppBase) AuthMiddleware() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			auth_header := r.Header.Get("Authorization")
			jwt_token, err := utils.GetBearerToken(auth_header)
			if err != nil {
				utils.ErrorResponse(w, http.StatusUnauthorized, "unauthorized")
				return
			}
			log.Println(jwt_token)

			next.ServeHTTP(w, r)
		})
	}
}
