package middleware

import "github.com/gofiber/fiber/v2"

func AuthMiddleware(ctx *fiber.Ctx) error {
	auth := ctx.Get("x-token")
	if auth != "secret" {
		return ctx.Status(401).JSON(fiber.Map{
			"Message": "Unauthorized",
		})
	}
	return ctx.Next()
}

func PermissionMiddleware(ctx *fiber.Ctx) error {
	return ctx.Next()
}
