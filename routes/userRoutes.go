package routes

import (
	"fiber-rest/controllers/userControllers"
	"fiber-rest/middleware"
	"github.com/gofiber/fiber/v2"
)

func UserRoutes(r *fiber.App) {
	r.Get("/users", middleware.AuthMiddleware, userControllers.GetAll)
	r.Get("/users/:id", middleware.AuthMiddleware, userControllers.GetUser)
	r.Post("/users", middleware.AuthMiddleware, userControllers.CreateUser)
	r.Put("/users/:id", middleware.AuthMiddleware, userControllers.UpdateUser)
	r.Delete("/users/:id", middleware.AuthMiddleware, userControllers.DeleteUser)
}
