package routes

import (
	"payslips/internal/handlers"
	"payslips/internal/middleware"

	"github.com/gofiber/fiber/v2"
)

func OvertimeRouter(app *fiber.App, handler handlers.Handlers, m *middleware.Authentication) {
	app.Post("/overtime", m.Authentication, m.AuditLog, handler.Overtime.Create)
	app.Put("/overtime/:id", m.Authentication, m.AuditLog, handler.Overtime.Update)
	app.Get("/overtime", m.Authentication, m.AuditLog, handler.Overtime.List)
}
