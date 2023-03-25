package app

import (
	"context"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/markbates/goth/gothic"
	"github.com/stadio-app/stadio-backend/ent/user"
	"github.com/stadio-app/stadio-backend/graph/model"
	"github.com/stadio-app/stadio-backend/types"
	"github.com/stadio-app/stadio-backend/utils"
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
			auth_header := r.Header.Get(types.Authorization)
			r = r.WithContext(context.WithValue(
				r.Context(),
				types.AuthHeader,
				auth_header,
			))
			w.Header().Add("Content-Type", "application/json")
			next.ServeHTTP(w, r)
		})
	}
}

func (app AppBase) AuthMiddleware() FuncHandler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			auth_header := r.Context().Value(types.AuthHeader).(string)
			ctx, err := app.BearerAuthentication(r.Context(), auth_header)
			if err != nil {
				utils.ErrorResponse(w, http.StatusUnauthorized, "unauthorized")
				return
			}
			r = r.WithContext(ctx)
			next.ServeHTTP(w, r)
		})
	}
}

func (app *AppBase) BearerAuthentication(ctx context.Context, auth_header string) (context.Context, error) {
	jwt_token, err := utils.GetBearerToken(auth_header)
	if err != nil {
		return nil, err
	}
	jwt_claims, err := utils.GetJwtClaims(jwt_token, app.Tokens.JwtKey)
	if err != nil {
		return nil, err
	}

	raw_id, ok := jwt_claims["id"].(float64)
	if !ok {
		return nil, fmt.Errorf("could not parse id")
	}
	id := int64(raw_id)
	email := jwt_claims["email"].(string)
	user, err := app.EntityManager.User.
		Query().
		Where(user.And(
			user.ID(id),
			user.Email(email),
		)).
		First(ctx)
	if err != nil {
		return nil, err
	}
	return context.WithValue(
		ctx,
		types.AuthKey, 
		model.AuthState{
			User: user,
			Token: jwt_token,
		},
	), nil
}
