package routes

import (
	"fiber-rest/controllers/authController"
	"github.com/gofiber/fiber/v2"
)

func AuthRoutes(r *fiber.App) {
	r.Post("/login", authController.LoginController)
	r.Post("/register", authController.RegisterController)
	r.Get("/refresh-token", authController.RefreshTokenController)
}
