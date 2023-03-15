package app

import (
	"context"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/markbates/goth/gothic"
	"github.com/stadio-app/go-backend/ent/user"
	"github.com/stadio-app/go-backend/types"
	"github.com/stadio-app/go-backend/utils"
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

			id, err := uuid.Parse(jwt_claims["id"].(string))
			if err != nil {
				utils.ErrorResponse(w, http.StatusUnauthorized, "invalid id claim")
				return
			}
			email := jwt_claims["email"].(string)
			user, err := app.EntityManager.User.
				Query().
				Where(user.And(
					user.ID(id),
					user.Email(email),
				)).
				First(r.Context())
			if err != nil {
				utils.ErrorResponse(w, http.StatusUnauthorized, "unauthorized")
				return
			}
			r = r.WithContext(context.WithValue(
				r.Context(), 
				types.AuthKey, 
				// model.AuthState{
				// 	User: user,
				// 	Token: jwt_token,
				// },
			))
			next.ServeHTTP(w, r)
		})
	}
}
