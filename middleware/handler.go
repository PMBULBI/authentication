package middleware

import (
	"github.com/gofiber/fiber/v2"
)

func (a *AuthMiddleware) IsRegistered(ctx *fiber.Ctx) (err error) {
	dataI := ctx.Locals(a.TokenHeader)
	if dataI == nil {
		return fiber.NewError(fiber.StatusUnauthorized, "Error When Decoding token")
	}

	data, ok := dataI.(string)
	if !ok {
		return fiber.NewError(fiber.StatusUnauthorized, "Error When Decoding token")
	}

	_, err = a.exist.CheckUserExists(ctx.Context(), "", data)
	if err != nil {
		return fiber.NewError(fiber.StatusUnauthorized, "User tidak terdaftar")
	}

	err = ctx.Next()
	return
}
