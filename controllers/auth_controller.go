package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/m3-app/backend/utils"
	"github.com/shareed2k/goth_fiber"
)

// [GET] /auth/:provider
func (ctx Controller) SignIn(c *fiber.Ctx) error {
	return goth_fiber.BeginAuthHandler(c)
}

// [GET] /auth/:provider/callback
func (ctx Controller) Callback(c *fiber.Ctx) error {
	provider_user, err := goth_fiber.CompleteUserAuth(c)
	if err != nil {
		return utils.FailResponse(c, "could not complete oauth transaction", err.Error())
	}

	return utils.DataResponse(c, provider_user)
}
