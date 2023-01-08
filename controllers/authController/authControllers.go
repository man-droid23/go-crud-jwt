package authController

import (
	"fiber-rest/configs/db"
	"fiber-rest/models/entity"
	"fiber-rest/models/requests"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

func Login(c *fiber.Ctx) error {
	var user entity.User
	req := new(requests.LoginRequest)
	if err := c.BodyParser(req); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"Message": "Cannot Parse Request Body",
			"Error":   err.Error(),
		})
	}
	errValidate := req.Validate()
	if errValidate != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"Message": "Request Validation Error",
			"Error":   errValidate.Error(),
		})
	}
	err := db.DB.Where("email = ?", req.Email).First(&user).Error
	if err != nil {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{
			"Message": "No Data Found",
			"Error":   err.Error(),
		})
	}
	if user.Password != req.Password {
		return c.Status(http.StatusUnauthorized).JSON(fiber.Map{
			"Message": "Invalid Password",
		})
	}
	return c.Status(http.StatusOK).JSON(fiber.Map{
		"Message": "Login Success",
		"Email":   user.Email,
	})
}
