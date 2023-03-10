package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/m3-app/backend/services"
	"github.com/m3-app/backend/types"
	"github.com/m3-app/backend/utils"
)

type Controller struct {
	types.AppContext
	Service services.Service
}

func New(app_ctx types.AppContext, service services.Service) Controller {
	controller := Controller{
		AppContext: app_ctx,
		Service: service,
	}
	server := app_ctx.Server

	server.Get("/", func(c *fiber.Ctx) error {
		return utils.ResponseWithStatusCode(c, 200, types.Response{
			Message: "Hello world",
		})
	})
	auth := server.Group("/auth")
	auth.Get("/:provider", controller.SignIn)
	auth.Get("/:provider/callback", controller.Callback)
	return controller
}
