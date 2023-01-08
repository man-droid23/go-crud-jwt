package middleware

import (
	"fiber-rest/utils"
	"github.com/gofiber/fiber/v2"
)

func AuthMiddleware(ctx *fiber.Ctx) error {
	authHeader := ctx.Get("Authorization")
	if authHeader == "" {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"Message": "Unauthorized",
			"Error":   "Missing",
		})
	}
	tokenValidation, err := utils.ValidateTokenJWT(authHeader)
	if err != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"Message": "Unauthorized",
		})
	}
	if tokenValidation.Valid {
		return ctx.Next()
	}
	return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
		"Message": "Unauthorized",
	})
}

func PermissionMiddleware(ctx *fiber.Ctx) error {
	return ctx.Next()
}
