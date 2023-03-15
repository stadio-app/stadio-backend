package services

import (
	"github.com/stadio-app/go-backend/ent"
)

type UserService struct {
	ServiceConfig
}

func (s UserService) FindOrCreate(email string, data interface{}) (*ent.User, error) {
	// user, err := s.EntityManager.User.
	// 	Query().
	// 	Where(user.Email(email)).
	// 	First(context.TODO())
	// if err != nil {
	// 	// create new user
	// 	user_builder := s.EntityManager.User.
	// 		Create().
	// 		SetEmail(email).
	// 		SetName(data.Name).
	// 		SetActive(true)
	// 	if data.AvatarURL != nil {
	// 		user_builder = user_builder.SetAvatar(*data.AvatarURL)
	// 	}
	// 	return user_builder.Save(context.TODO())
	// }
	// return user, nil
	return nil, nil
}
