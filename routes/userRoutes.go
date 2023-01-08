package routes

import (
	"fiber-rest/controllers/userControllers"
	"fiber-rest/middleware"
	"github.com/gofiber/fiber/v2"
)

func UserRoutes(r *fiber.App) {
	r.Get("/users", middleware.AuthMiddleware, userControllers.GetAll)
	r.Get("/users/:id", userControllers.GetUser)
	r.Post("/users", userControllers.CreateUser)
	r.Put("/users/:id", userControllers.UpdateUser)
	r.Delete("/users/:id", userControllers.DeleteUser)
}
