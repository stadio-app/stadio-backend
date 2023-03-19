package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.26

import (
	"context"
	"fmt"

	"github.com/stadio-app/stadio-backend/ent"
	"github.com/stadio-app/stadio-backend/graph/model"
	"github.com/stadio-app/stadio-backend/types"
	"github.com/stadio-app/stadio-backend/utils"
)

// CreateOwner is the resolver for the createOwner field.
func (r *mutationResolver) CreateOwner(ctx context.Context, input model.OwnerInput) (*ent.Owner, error) {
	auth_state := utils.ParseAuthContext(ctx)
	user_owner, err := auth_state.User.Owner(ctx)
	if err == nil && user_owner != nil {
		return nil, fmt.Errorf("owner already exists for this account")
	}
	owner := r.EntityManager.Owner.Create().
		SetFirstName(input.FirstName).
		SetLastName(input.LastName)
	if input.MiddleName != nil {
		owner = owner.SetMiddleName(*input.MiddleName)
	}
	return owner.
		SetFullName(utils.CreateFullName(types.FullName{
			FirstName:  input.FirstName,
			MiddleName: input.MiddleName,
			LastName:   input.LastName,
		})).
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
