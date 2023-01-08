package utils

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/joho/godotenv"
	"os"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
}

var JwtSecret = []byte(os.Getenv("JWT_SECRET"))

func GenerateTokenJWT(claims *jwt.MapClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(JwtSecret)
}

func ValidateTokenJWT(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, isValid := token.Method.(*jwt.SigningMethodHMAC); !isValid {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return JwtSecret, nil
	})
	if err != nil {
		return nil, err
	}
	return token, nil
}

func DecodeTokenJWT(tokenString string) (*jwt.MapClaims, error) {
	tokenValidation, err := ValidateTokenJWT(tokenString)
	if err != nil {
		return nil, err
	}
	claims, ok := tokenValidation.Claims.(jwt.MapClaims)
	if !ok {
		return nil, err
	}
	return &claims, nil
}
