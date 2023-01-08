package main

import (
	"fiber-rest/configs/db"
	"fiber-rest/configs/migration"
	"fiber-rest/configs/static"
	"fiber-rest/routes"
	"github.com/gofiber/fiber/v2"
)

func init() {
	db.DatabaseConnection()
	migration.Migration()
}

func main() {
	defer db.CloseDatabaseConnection()
	r := fiber.New()
	r.Static("/public", static.ProjectRootPath+"/public/assets")
	routes.UserRoutes(r)
	routes.AuthRoutes(r)
	_ = r.Listen(":3000")
}
