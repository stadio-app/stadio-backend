package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.26

import (
	"context"
	"fmt"

	"github.com/stadio-app/go-backend/ent"
	"github.com/stadio-app/go-backend/graph/model"
	"github.com/stadio-app/go-backend/utils"
)

// CreateOwner is the resolver for the createOwner field.
func (r *mutationResolver) CreateOwner(ctx context.Context, input model.OwnerInput) (*ent.Owner, error) {
	auth_state := utils.ParseAuthContext(ctx)
	return r.EntityManager.Owner.Create().
		SetFirstName(input.FirstName).
		SetMiddleName(*input.MiddleName).
		SetLastName(input.LastName).
		SetFullName(fmt.Sprintf("%s %s %s", input.FirstName, *input.MiddleName, input.LastName)).
		SetIDURL(input.IDURL).
		SetUser(auth_state.User).
		SetUserID(auth_state.User.ID).
		Save(ctx)
}

// Users is the resolver for the users field.
func (r *queryResolver) Users(ctx context.Context) ([]*ent.User, error) {
	return r.EntityManager.User.Query().All(ctx)
}

// User is the resolver for the user field.
func (r *queryResolver) User(ctx context.Context, id int64) (*ent.User, error) {
	return r.EntityManager.User.Get(ctx, id)
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }
