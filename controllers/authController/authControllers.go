package authController

import (
	"fiber-rest/configs/db"
	"fiber-rest/models/entity"
	"fiber-rest/models/requests"
	"fiber-rest/utils"
	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"log"
	"net/http"
	"time"
)

func LoginController(c *fiber.Ctx) error {
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
	errLogin := db.DB.Where("email = ?", req.Email).First(&user).Error
	if errLogin != nil {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{
			"Message": "No Data Found",
			"Error":   errLogin.Error(),
		})
	}

	pass := utils.CheckPasswordHash(req.Password, user.Password)
	if !pass {
		return c.Status(http.StatusUnauthorized).JSON(fiber.Map{
			"Message": "Wrong Password",
		})
	}

	claims := jwt.MapClaims{}
	claims["email"] = user.Email
	claims["exp"] = time.Now().Add(time.Minute * 2).Unix()
	token, errToken := utils.GenerateTokenJWT(&claims)
	if errToken != nil {
		log.Println(errToken)
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"Message": "Cannot Generate Token",
		})
	}
	return c.Status(http.StatusOK).JSON(fiber.Map{
		"Message": "Login Success",
		"Token":   token,
	})
}

func RegisterController(ctx *fiber.Ctx) error {
	req := new(requests.RegisterRequest)
	if err := ctx.BodyParser(req); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"Message": "Cannot Parse Request Body",
			"Error":   err.Error(),
		})
	}
	errValidate := req.Validate()
	if errValidate != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"Message": "Request Validation Error",
			"Error":   errValidate.Error(),
		})
	}
	newUser := entity.User{
		Email:    req.Email,
		Password: req.Password,
	}
	hash, err := utils.HashPassword(newUser.Password)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"Message": "Cannot Hash Password",
			"Error":   err.Error(),
		})
	}
	newUser.Password = hash
	errCreate := db.DB.Create(&newUser).Error
	if errCreate != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"Message": "Cannot Create User",
			"Error":   errCreate.Error(),
		})
	}
	return ctx.Status(http.StatusCreated).JSON(fiber.Map{
		"Message": "User Created",
		"Email":   newUser.Email,
	})
}

func RefreshTokenController(ctx *fiber.Ctx) error {
	authHeader := ctx.Get("Authorization")
	if authHeader == "" {
		return ctx.Status(http.StatusUnauthorized).JSON(fiber.Map{
			"Message": "Unauthorized",
			"Error":   "Missing",
		})
	}
	tokenValidation, err := utils.ValidateTokenJWT(authHeader)
	if err != nil {
		return ctx.Status(http.StatusUnauthorized).JSON(fiber.Map{
			"Message": "Unauthorized",
		})
	}
	claims := tokenValidation.Claims.(jwt.MapClaims)
	claims["exp"] = time.Now().Add(time.Minute * 2).Unix()
	newToken, err := utils.GenerateTokenJWT(&claims)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"Message": "Cannot Generate Token",
		})
	}
	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"Message": "Token Refreshed",
		"Token":   newToken,
	})
}
