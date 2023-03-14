package services

import (
	"context"

	"github.com/m3-app/backend/ent"
	"github.com/m3-app/backend/ent/user"
	"github.com/m3-app/backend/graph/model"
)

type UserService struct {
	ServiceConfig
}

func (s UserService) FindOrCreate(email string, data *model.UserInput) (*ent.User, error) {
	user, err := s.EntityManager.User.
		Query().
		Where(user.Email(email)).
		First(context.TODO())
	if err != nil {
		// create new user
		user_builder := s.EntityManager.User.
			Create().
			SetEmail(email).
			SetName(data.Name).
			SetActive(true)
		if data.AvatarURL != nil {
			user_builder = user_builder.SetAvatar(*data.AvatarURL)
		}
		return user_builder.Save(context.TODO())
	}
	return user, nil
}
