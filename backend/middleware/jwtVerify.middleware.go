package middleware

import (
	"os"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/kshitij/taskManager/loadEnv"
)

type MiddlewareService interface {
	JWTMiddleWare(ctx *fiber.Ctx) error
}

type middlewareService struct {
	loadEnvService loadEnv.LoadEnvService
}

func NewMiddlewareService(loadEnvService loadEnv.LoadEnvService) MiddlewareService {
	return &middlewareService{loadEnvService}
}

func (m *middlewareService) JWTMiddleWare(ctx *fiber.Ctx) error {

	authToken := ctx.Get("Authorization");
	tokenStr := strings.Split(authToken, " ")[1]
	if tokenStr == "" {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Missing or malformed JWT",
		})
	}

	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fiber.NewError(fiber.StatusUnauthorized, "Invalid token Signing Method")
		}
		m.loadEnvService.LoadEnv();
		jwtSecret := os.Getenv("JWT_SECRET")
		return []byte(jwtSecret), nil;
	})

	if err != nil || !token.Valid {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Invalid or Expired JWt",
		})
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		userIDFloat, ok := claims["user_id"].(float64)
		if !ok {
			return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "Invalid user_id format in token",
			})
		}
		ctx.Locals("user_id", uint64(userIDFloat))
	}

	return ctx.Next();

}