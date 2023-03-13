package app

import (
	"context"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/m3-app/backend/ent/user"
	"github.com/m3-app/backend/graph/model"
	"github.com/m3-app/backend/types"
	"github.com/m3-app/backend/utils"
	"github.com/markbates/goth/gothic"
)

type FuncHandler func(http.Handler) http.Handler

func (app AppBase) GothMiddleware() FuncHandler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			r = gothic.GetContextWithProvider(r, chi.URLParam(r, "provider"))
			next.ServeHTTP(w, r)
		})
	}
}

func (app AppBase) BaseMiddleware() FuncHandler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Add("Content-Type", "application/json")
			next.ServeHTTP(w, r)
		})
	}
}

func (app AppBase) AuthMiddleware() FuncHandler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			today := time.Now()
			auth_header := r.Header.Get(types.AuthHeader)
			jwt_token, err := utils.GetBearerToken(auth_header)
			if err != nil {
				utils.ErrorResponse(w, http.StatusUnauthorized, "unauthorized")
				return
			}
			
			jwt_claims, err := utils.GetJwtClaims(jwt_token, app.Tokens.JwtKey)
			if err != nil {
				utils.ErrorResponse(w, http.StatusUnauthorized, err.Error())
				return
			}

			expiration, err := jwt_claims.GetExpirationTime()
			if !today.Before(expiration.Time) && err != nil {
				utils.ErrorResponse(w, http.StatusUnauthorized, "token expired")
				return
			}

			email := jwt_claims["email"].(string)
			user, err := app.EntityManager.User.
				Query().
				Where(user.Email(email)).
				First(r.Context())
			if err != nil {
				utils.ErrorResponse(w, http.StatusUnauthorized, "unauthorized")
				return
			}
			r = r.WithContext(context.WithValue(
				r.Context(), 
				types.AuthKey, 
				model.AuthState{
					User: user,
					Token: jwt_token,
				},
			))
			next.ServeHTTP(w, r)
		})
	}
}
