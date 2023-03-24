package app

import (
	"context"
	"fmt"

	"github.com/stadio-app/stadio-backend/ent/user"
	"github.com/stadio-app/stadio-backend/graph/model"
	"github.com/stadio-app/stadio-backend/types"
	"github.com/stadio-app/stadio-backend/utils"
)

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
