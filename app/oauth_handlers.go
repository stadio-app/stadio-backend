package app

import (
	"net/http"

	"github.com/markbates/goth/gothic"
	"github.com/stadio-app/go-backend/utils"
)

func (app AppBase) OAuthSignIn(w http.ResponseWriter, r *http.Request) {
	gothic.BeginAuthHandler(w, r)
}

func (app AppBase) OAuthCallback(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	provider_user, err := gothic.CompleteUserAuth(w, r)
	if err != nil {
		utils.ErrorResponse(w, 400, "could not complete oauth transaction")
		return
	}

	// user, err := app.Services.UserService.FindOrCreate(
	// 	provider_user.Email,
	// 	&model.UserInput{
	// 		Email: provider_user.Email,
	// 		Name: provider_user.Name,
	// 		AvatarURL: &provider_user.AvatarURL,
	// 	},
	// )
	// if err != nil {
	// 	utils.ErrorResponse(w, 400, "could not find or create user")
	// 	return
	// }
	// token, err := utils.GenerateJWT(app.Tokens.JwtKey, user)
	// if err != nil {
	// 	utils.ErrorResponse(w, 400, err.Error())
	// 	return
	// }
	// auth_state := model.AuthState{
	// 	User: user,
	// 	Token: token,
	// }
	// utils.DataResponse(w, auth_state)
}
