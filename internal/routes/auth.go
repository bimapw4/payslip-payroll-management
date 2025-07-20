package routes

import (
	"payslips/internal/handlers"
	"payslips/internal/middleware"

	"github.com/gofiber/fiber/v2"
)

func AuthRouter(app *fiber.App, handler handlers.Handlers, m *middleware.Authentication) {
	app.Post("/login", handler.Auth.Login)
}
