package common

import (
	"payslips/internal/entity"

	"github.com/gofiber/fiber/v2"
)

func SetUserCtx(c *fiber.Ctx, claims *entity.UserProfile) {
	c.Locals("user_ctx", claims)
}

func GetUserCtx(c *fiber.Ctx) *entity.UserProfile {
	return c.Locals("user_ctx").(*entity.UserProfile)
}

func SetTokenCtx(c *fiber.Ctx, token string) {
	c.Locals("token", token)
}

func GetTokenCtx(c *fiber.Ctx) string {
	return c.Locals("token").(string)
}
