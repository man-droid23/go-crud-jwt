package routes

import (
	"fiber-rest/controllers/authController"
	"github.com/gofiber/fiber/v2"
)

func AuthRoutes(r *fiber.App) {
	r.Post("/login", authController.Login)
	//r.Post("/register", authControllers.Register)
}
