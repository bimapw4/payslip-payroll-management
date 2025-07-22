package routes

import (
	"payslips/internal/handlers"
	"payslips/internal/middleware"

	"github.com/gofiber/fiber/v2"
)

func ReimbursementRouter(app *fiber.App, handler handlers.Handlers, m *middleware.Authentication) {
	app.Post("/reimbursment", m.Authentication, m.AuditLog, handler.Reimbursement.Create)
	app.Put("/reimbursment", m.Authentication, m.AuditLog, handler.Reimbursement.Update)
	app.Get("/reimbursment", m.Authentication, m.AuditLog, handler.Reimbursement.List)
}
