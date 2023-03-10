package utils

import (
	"github.com/gofiber/fiber/v2"
	"github.com/m3-app/backend/types"
)

func ResponseWithStatusCode(c *fiber.Ctx, statusCode int, data interface{}) error {
	return c.Status(statusCode).JSON(data)
}

// Generic json response with status code 200
func JsonResponse(c *fiber.Ctx, data interface{}) error {
	return ResponseWithStatusCode(c, fiber.StatusOK, data)
}

// types.Error json response with status code 400
func FailResponse(c *fiber.Ctx, errors ...string) error {
	return ResponseWithStatusCode(c, fiber.StatusBadRequest, types.Errors{
		Errors: errors,
	})
}

// types.Error json response with status code 401
func FailResponseUnauthorized(c *fiber.Ctx, errors ...string) error {
	return ResponseWithStatusCode(c, fiber.StatusUnauthorized, types.Errors{
		Errors: errors,
	})
}

// types.Data json response with status code 200
func DataResponse(c *fiber.Ctx, data interface{}) error {
	return ResponseWithStatusCode(c, fiber.StatusOK, types.Result{
		Data: data,
	})
}

// types.Data json response with status code 201
func DataResponseCreated(c *fiber.Ctx, data interface{}) error {
	return ResponseWithStatusCode(c, fiber.StatusCreated, types.Result{
		Data: data,
	})
}