package routes

import (
	"payslips/internal/handlers"
	"payslips/internal/middleware"

	"github.com/gofiber/fiber/v2"
)

func Routes(app *fiber.App, handler handlers.Handlers, m *middleware.Authentication) {
	// register route
	routes := []func(app *fiber.App, handler handlers.Handlers, m *middleware.Authentication){}

	for _, route := range routes {
		route(app, handler, m)
	}
}
