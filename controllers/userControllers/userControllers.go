package userControllers

import (
	"fiber-rest/configs/db"
	"fiber-rest/models/entity"
	"fiber-rest/models/requests"
	"fiber-rest/models/response"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

func GetAll(ctx *fiber.Ctx) error {
	var users []entity.User
	result := db.DB.Find(&users)
	if result.Error != nil {
		return ctx.Status(http.StatusNotFound).JSON(fiber.Map{
			"Message": "No Data Found",
			"Error":   result.Error.Error(),
		})
	}
	return ctx.Status(http.StatusFound).JSON(fiber.Map{
		"Message": "Data Has Been Found",
		"Data":    response.NewUserListResponse(users),
	})
}

func GetUser(ctx *fiber.Ctx) error {
	var user entity.User
	id := ctx.Params("id")
	result := db.DB.Find(&user, id)
	if result.Error != nil {
		return ctx.Status(http.StatusNotFound).JSON(fiber.Map{
			"Message": "No Data Found",
			"Error":   result.Error.Error(),
		})
	}
	return ctx.Status(http.StatusFound).JSON(fiber.Map{
		"Message": "Data Has Been Found",
		"Data":    response.NewUserResponse(user),
	})
}

func CreateUser(ctx *fiber.Ctx) error {
	users := new(requests.UserRequest)
	err := ctx.BodyParser(users)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"Message": "Cannot Parse Request Body",
			"Error":   err.Error(),
		})
	}
	errValidate := users.Validate()
	if errValidate != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"Message": "Request Validation Error",
			"Error":   errValidate.Error(),
		})
	}
	newUser := entity.User{
		Email:    users.Email,
		Password: users.Password,
	}
	errCreateUser := db.DB.Create(&newUser).Error
	if errCreateUser != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"Message": "Failed to Create User",
			"Error":   errCreateUser.Error(),
		})
	}
	return ctx.Status(http.StatusCreated).JSON(fiber.Map{
		"Message": "Successfully Created User",
		"Data":    response.NewUserResponse(newUser),
	})
}

func UpdateUser(ctx *fiber.Ctx) error {
	var user entity.User
	id := ctx.Params("id")
	result := db.DB.Find(&user, id)
	if result.Error != nil {
		return ctx.Status(http.StatusNotFound).JSON(fiber.Map{
			"Message": "No Data Found",
			"Error":   result.Error.Error(),
		})
	}
	users := new(requests.UserRequest)
	err := ctx.BodyParser(users)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"Message": "Cannot Parse Request Body",
			"Error":   err.Error(),
		})
	}
	errValidate := users.Validate()
	if errValidate != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"Message": "Request Validation Error",
			"Error":   errValidate.Error(),
		})
	}
	user.Email = users.Email
	user.Password = users.Password
	errUpdateUser := db.DB.Save(&user).Error
	if errUpdateUser != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"Message": "Failed to Update User",
			"Error":   errUpdateUser.Error(),
		})
	}
	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"Message": "Successfully Updated User",
		"Data":    response.NewUserResponse(user),
	})
}

func DeleteUser(ctx *fiber.Ctx) error {
	var user entity.User
	id := ctx.Params("id")
	result := db.DB.Find(&user, id)
	if result.Error != nil {
		return ctx.Status(http.StatusNotFound).JSON(fiber.Map{
			"Message": "No Data Found",
			"Error":   result.Error.Error(),
		})
	}
	errDeleteUser := db.DB.Delete(&user).Error
	if errDeleteUser != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"Message": "Failed to Delete User",
			"Error":   errDeleteUser.Error(),
		})
	}
	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"Message": "Successfully Deleted User",
		"Data":    response.NewUserResponse(user),
	})
}
