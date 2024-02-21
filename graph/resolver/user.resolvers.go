package gresolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.44

import (
	"context"
	"fmt"

	"github.com/stadio-app/stadio-backend/graph"
	"github.com/stadio-app/stadio-backend/graph/gmodel"
)

// CreateAccount is the resolver for the createAccount field.
func (r *mutationResolver) CreateAccount(ctx context.Context, input gmodel.CreateAccountInput) (*gmodel.User, error) {
	new_user, err := r.Service.CreateInternalUser(ctx, input)
	if err != nil {
		return nil, fmt.Errorf("could not create user. %s", err.Error())
	}
	return &new_user, nil
}

// Login is the resolver for the login field.
func (r *queryResolver) Login(ctx context.Context, email string, password string) (*gmodel.Auth, error) {
	auth, err := r.Service.LoginInternal(ctx, email, password)
	return &auth, err
}

// Mutation returns graph.MutationResolver implementation.
func (r *Resolver) Mutation() graph.MutationResolver { return &mutationResolver{r} }

// Query returns graph.QueryResolver implementation.
func (r *Resolver) Query() graph.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
