package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.26

import (
	"context"
	"strconv"

	"github.com/stadio-app/go-backend/ent"
)

// Users is the resolver for the users field.
func (r *queryResolver) Users(ctx context.Context) ([]*ent.User, error) {
	return r.EntityManager.User.Query().All(ctx)
}

// User is the resolver for the user field.
func (r *queryResolver) User(ctx context.Context, id string) (*ent.User, error) {
	parsed_id, err := strconv.ParseInt(id, 10, 32)
	if err != nil {
		return nil, err
	}
	return r.EntityManager.User.Get(ctx, int(parsed_id))
}
