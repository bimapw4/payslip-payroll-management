package routes

import (
	"payslips/internal/handlers"
	"payslips/internal/middleware"

	"github.com/gofiber/fiber/v2"
)

func AuditLogRouter(app *fiber.App, handler handlers.Handlers, m *middleware.Authentication) {
	app.Get("/audit-log", m.Authentication, m.AuditLog, handler.AuditLog.List)
}
